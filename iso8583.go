package iso8583

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/albenik/bcd"
)

func fromBCD(m []byte) (str string, err error) {
	count := len(m)
	if count != 8 && count != 16 {
		return "", errors.New("Invalid bitmap length")
	}

	//fmt.Println(GetBit(m, 0), m[0])
	if !GetBit(m, 0) && count == 16 {
		count /= 2
	}

	for index := 0; index < count; index++ {
		str += fmt.Sprint(bcd.ToUint8(m[index]))
	}
	return
}

// ToString packs the mti, bitmap and elements into a string
func (iso *IsoStruct) ToString() (string, error) {
	var str string
	// get done with the mti and the bitmap
	bitmapString := hex.EncodeToString(iso.Bitmap())

	// bcdInts, err := fromBCD(iso.Bitmap)
	// if err != nil {
	// 	return "", err
	// }

	// fmt.Printf("Bitmap: %v\n", bcdInts)

	elementsStr, err := iso.packElements()
	if err != nil {
		return str, err
	}
	str = iso.mti + bitmapString + elementsStr
	return str, nil
}

// AddMTI adds the provided iso8583 MTI into the current struct
// also updates the bitmap in the process
func (iso *IsoStruct) AddMTI(data string) error {
	if err := ValidateMti(data); err != nil {
		return err
	}

	iso.mti = data
	//fmt.Printf("MTI: %s\n", iso.Mti)
	return nil
}

// AddField adds the provided iso8583 field into the current struct
// also updates the bitmap in the process
func (iso *IsoStruct) AddField(field int, data string) error {
	if field < 2 || field > BitMapLen(iso.bitmap) {
		return fmt.Errorf("expected field to be between %d and %d found %d instead", 2, len(iso.bitmap), field)
	}

	if field >= 65 {
		SetBit(iso.bitmap, 0)
		if len(iso.bitmap) == 8 {
			iso.bitmap = append(iso.bitmap, make([]byte, 8)...) //increase bitmap
		}
	}

	SetBit(iso.bitmap, int(field-1))
	iso.data[field] = Field{Value: data}
	//fmt.Printf("Field %d: %s\n", field, iso.Elements.elements[field])
	return nil
}

// Parse parses an iso8583 string
func Parse(i string) (iso IsoStruct, err error) {
	spec, mti, rest, err := extractMTI(i)
	if err != nil {
		return
	}

	bitMap, elementString, err := extractBitmap(rest)
	if err != nil {
		return
	}

	//fmt.Printf("BitMap: %08b, Elements: %v", bitMap, elementString)

	// validat the mti
	if err = ValidateMti(mti); err != nil {
		return
	}

	data, err := unpackElements(bitMap, elementString, spec)
	if err != nil {
		return
	}

	//fmt.Printf("BitMap: %08b, Elements: %v, Unpacked: %v\n", bitMap, elementString, elements)

	iso = IsoStruct{spec: spec, mti: mti, bitmap: bitMap, data: data}
	return iso, nil
}

func (iso *IsoStruct) packElements() (string, error) {
	var str string
	bitmapLength := BitMapLen(iso.bitmap)
	elementsMap := iso.data
	elementsSpec := iso.Spec()

	for index := 1; index < bitmapLength; index++ { // index 0 of bitmap isn't need here
		if GetBit(iso.bitmap, index) { // if the field is present
			field := index + 1
			fieldDescription := elementsSpec.fields[field]
			if fieldDescription.LenType == "fixed" {
				str = str + elementsMap[field].Value
			} else {
				lengthType, err := getVariableLengthFromString(fieldDescription.LenType)
				if err != nil {
					return str, err
				}
				actualLength := len(elementsMap[field].Value)
				paddedLength := leftPad(strconv.Itoa(actualLength), int(lengthType), "0")
				str = str + (paddedLength + elementsMap[field].Value)
			}
		}
	}
	return str, nil
}

// extractMTI extracts the mti from an iso8583 string
func extractMTI(str string) (spec Spec, mti, rest string, err error) {
	mti, rest = str[0:4], str[4:len(str)]

	specName := ""
	switch string(mti[0]) {
	case "0":
		specName = "1987"
	case "1":
		specName = "1993"
	case "2":
		specName = "2003"
	default:
		return Spec{}, "", "", fmt.Errorf("Invalid mti version %v", string(mti[0]))
	}

	spec, ok := specs[specName]
	if !ok {
		return Spec{}, "", "", fmt.Errorf("iso8583:%s is not supported", specName)
	}

	return
}

func extractBitmap(rest string) (bytes []byte, elementsString string, err error) {
	mapLen := 16

	if len(rest) < mapLen {
		err = errors.New("Invalid bitmap length")
		return
	}
	tempMap := rest[:mapLen]

	// only 64 bit primary bitmap
	if bytes, err = hex.DecodeString(tempMap); err != nil {
		err = fmt.Errorf("Error Decoding string: %v", err)
		return
	}

	if GetBit(bytes, 0) {
		mapLen = 32
		// secondary bitmap exists, so total 128 bits
		if bytes, err = hex.DecodeString(tempMap[:mapLen]); err != nil {
			err = fmt.Errorf("Error Decoding string: %v", err)
			return
		}
	}

	elementsString = rest[mapLen:]

	return bytes, elementsString, nil
}

func getVariableLengthFromString(str string) (int, error) {
	if str != "llvar" && str != "lllvar" && str != "llllvar" {
		return 0, fmt.Errorf("%s is an invalid LenType", str)
	}
	return strings.Count(str, "l"), nil
}

func extractFieldFromElements(fieldID, str string) (result Field, rest string, err error) {
	fieldDescription, indices, level, err := DescribeField(fieldID)
	if err != nil {
		return
	}

	field := indices[level]

	fieldLength := fieldDescription.MaxLen
	if fieldDescription.LenType == "fixed" && fieldDescription.MaxLen > len(str) {
		err = fmt.Errorf("Fixed length Field %d is longer than data string", field)
		return
	}

	result.ID = fieldID
	//fmt.Printf("Index: %d, Description: %v, Data: %s\n", field, fieldDescription, str)
	if fieldDescription.LenType == "fixed" {
		result.Value = str[:fieldLength]
		rest = strings.TrimPrefix(str, result.Value)
		//fmt.Printf("Index: %d, Description: %v, Extracted: %s, Data: %s\n", field, fieldDescription, extractedField, str)
	} else {
		// varianle length fields have their lengths embedded into the string
		digitCount, err2 := getVariableLengthFromString(fieldDescription.LenType)
		if err2 != nil {
			err = fmt.Errorf("spec error: field %d: %s", field, err2.Error())
			return
		}

		fieldLength, err = strconv.Atoi(str[0:digitCount])
		if err != nil {
			err = fmt.Errorf("spec error: field %d: %s", field, "invalid length digits")
			return
		}

		if fieldLength > len(str) {
			err = fmt.Errorf("Field %d's length is longer than data string", field)
			return
		}

		result.Value = str[:fieldLength]
		rest = strings.TrimPrefix(str, result.Value)
		//fmt.Printf("Index: %d, Description: %v, Extracted: %s, Data: %s\n", field, fieldDescription, extractedField, str)
	}

	if fieldLength > fieldDescription.MaxLen {
		err = fmt.Errorf("Field %d's length of %d (%s) is greater than the maximum allowed of %d", field, fieldLength, result.Value, fieldDescription.MaxLen)
		return
	}

	currentString := result.Value
	for subIndex := range fieldDescription.Subfields {
		subResult, subRest, err2 := extractFieldFromElements(fmt.Sprintf("%s.%d", fieldID, subIndex), currentString)
		if err = err2; err != nil {
			return
		}
		if subRest == "" {
			break
		}

		currentString = subRest
		if result.Subfields == nil {
			result.Subfields = make(map[int]Field, len(fieldDescription.Subfields))
		}
		result.Subfields[subIndex] = subResult
	}

	return result, rest, nil
}

func unpackElements(bitMap []byte, elements string, spec Spec) (d map[int]Field, err error) {
	bitmapLength := BitMapLen(bitMap)
	d = make(map[int]Field)
	currentString := elements

	// The first (index 0) bit of the bitmap shows the presense(1)/absense(0) of the secondary
	// we therefore start with the second bit (index 1) which is field (2)
	//fmt.Printf("Bitmap: %b\n", bitMap)

	for index := 1; index < bitmapLength; index++ {
		//Index starts at 1 because index 0 is the second bitmap flag

		field := index + 1 // Index 1 represents field 2

		if !GetBit(bitMap, index) {
			continue //Field it not set, we can skip
		}

		fieldID := fmt.Sprintf("%s.%d", spec.Version(), field)
		extractedField, rest, err2 := extractFieldFromElements(fieldID, currentString)
		if err = err2; err != nil {
			return
		}

		currentString = rest
		//fmt.Println(currentString)

		d[field] = extractedField
		//fmt.Printf("Bitmap index: %d, Field %d: %v\n", index, field, m[field])

		if rest == "" {
			break
		}
	}

	return d, nil
}

func NewMessage(specName string) (iso IsoStruct, err error) {
	spec, ok := specs[specName]
	if !ok {
		err = errors.New("Invalid spec")
		return
	}

	var bitMap []byte
	mti := ""

	//We start with a single bitmap, and if we add any fields from position 65 onwards, we double the map
	bitMap = make([]byte, 8)

	elements := make(map[int]Field)
	iso = IsoStruct{spec: spec, mti: mti, bitmap: bitMap, data: elements}
	return
}
