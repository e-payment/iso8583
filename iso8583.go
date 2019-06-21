package iso8583

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/albenik/bcd"
)

// String returns the mti as a string
func (m *MtiType) String() string {
	return m.mti
}

// ElementsType stores iso8583 elements in a map
type ElementsType struct {
	elements map[int64]string
}

// GetElements returns the available elemts as a map
func (e *ElementsType) GetElements() map[int64]string {
	return e.elements
}

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
	bitmapString := hex.EncodeToString(iso.Bitmap)

	// bcdInts, err := fromBCD(iso.Bitmap)
	// if err != nil {
	// 	return "", err
	// }

	// fmt.Printf("Bitmap: %v\n", bcdInts)

	elementsStr, err := iso.packElements()
	if err != nil {
		return str, err
	}
	str = iso.Mti.String() + bitmapString + elementsStr
	return str, nil
}

// AddMTI adds the provided iso8583 MTI into the current struct
// also updates the bitmap in the process
func (iso *IsoStruct) AddMTI(data string) error {
	mti := MtiType{mti: data}
	_, err := MtiValidator(mti)
	if err != nil {
		return err
	}
	iso.Mti = mti
	//fmt.Printf("MTI: %s\n", iso.Mti)
	return nil
}

// AddField adds the provided iso8583 field into the current struct
// also updates the bitmap in the process
func (iso *IsoStruct) AddField(field int64, data string) error {
	if field < 2 || field > int64(BitMapLen(iso.Bitmap)) {
		return fmt.Errorf("expected field to be between %d and %d found %d instead", 2, len(iso.Bitmap), field)
	}

	if field >= 65 {
		SetBit(iso.Bitmap, 0)
		if len(iso.Bitmap) == 8 {
			iso.Bitmap = append(iso.Bitmap, make([]byte, 8)...) //increase bitmap
		}
	}

	SetBit(iso.Bitmap, int(field-1))
	iso.Elements.elements[field] = data
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
	if _, err = MtiValidator(mti); err != nil {
		return
	}

	elements, err := unpackElements(bitMap, elementString, spec)
	if err != nil {
		return
	}

	//fmt.Printf("BitMap: %08b, Elements: %v, Unpacked: %v\n", bitMap, elementString, elements)

	iso = IsoStruct{Spec: spec, Mti: mti, Bitmap: bitMap, Elements: elements}
	return iso, nil
}

func (iso *IsoStruct) packElements() (string, error) {
	var str string
	bitmapLength := BitMapLen(iso.Bitmap)
	elementsMap := iso.Elements.GetElements()
	elementsSpec := iso.Spec

	for index := 1; index < bitmapLength; index++ { // index 0 of bitmap isn't need here
		if GetBit(iso.Bitmap, index) { // if the field is present
			field := int64(index + 1)
			fieldDescription := elementsSpec.fields[field]
			if fieldDescription.LenType == "fixed" {
				str = str + elementsMap[field]
			} else {
				lengthType, err := getVariableLengthFromString(fieldDescription.LenType)
				if err != nil {
					return str, err
				}
				actualLength := len(elementsMap[field])
				paddedLength := leftPad(strconv.Itoa(actualLength), int(lengthType), "0")
				str = str + (paddedLength + elementsMap[field])
			}
		}
	}
	return str, nil
}

// extractMTI extracts the mti from an iso8583 string
func extractMTI(str string) (spec Spec, mti MtiType, rest string, err error) {
	mti.mti, rest = str[0:4], str[4:len(str)]

	specName := ""
	switch string(mti.mti[0]) {
	case "0":
		specName = "1987"
	case "1":
		specName = "1993"
	case "2":
		specName = "2003"
	default:
		return Spec{}, MtiType{}, "", fmt.Errorf("Invalid mti version %v", string(mti.mti[0]))
	}

	spec, ok := specs[specName]
	if !ok {
		return Spec{}, MtiType{}, "", fmt.Errorf("iso8583:%s is not supported", specName)
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

func getVariableLengthFromString(str string) (int64, error) {
	if str != "llvar" && str != "lllvar" && str != "llllvar" {
		return 0, fmt.Errorf("%s is an invalid LenType", str)
	}
	return int64(strings.Count(str, "l")), nil
}

func extractFieldFromElements(spec Spec, field int64, str string) (extractedField, rest string, err error) {
	fieldDescription := spec.fields[field]
	length := int64(fieldDescription.MaxLen)
	if fieldDescription.MaxLen > len(str) {
		length = int64(len(str))
		return str, "", nil
	}
	//fmt.Printf("Index: %d, Description: %v, Data: %s\n", field, fieldDescription, str)
	if fieldDescription.LenType == "fixed" {
		extractedField = str[:length]
		rest = strings.TrimPrefix(str, extractedField)
		//fmt.Printf("Index: %d, Description: %v, Extracted: %s, Data: %s\n", field, fieldDescription, extractedField, str)

	} else {
		// varianle length fields have their lengths embedded into the string
		digitCount, err2 := getVariableLengthFromString(fieldDescription.LenType)
		if err2 != nil {
			return extractedField, "", fmt.Errorf("spec error: field %d: %s", field, err2.Error())
		}

		tempLength, err2 := strconv.Atoi(str[0:digitCount])
		if err != nil {
			err = fmt.Errorf("spec error: field %d: %s", field, "invalid length digits")
			return
		}

		extractedField = str[:int64(tempLength)]
		rest = strings.TrimPrefix(str, extractedField)
		//fmt.Printf("Index: %d, Description: %v, Extracted: %s, Data: %s\n", field, fieldDescription, extractedField, str)
	}

	return extractedField, rest, nil
}

func unpackElements(bitMap []byte, elements string, spec Spec) (elem ElementsType, err error) {
	bitmapLength := BitMapLen(bitMap)
	var m = make(map[int64]string)
	currentString := elements

	// The first (index 0) bit of the bitmap shows the presense(1)/absense(0) of the secondary
	// we therefore start with the second bit (index 1) which is field (2)
	//fmt.Printf("Bitmap: %b\n", bitMap)

	for index := 1; index < bitmapLength; index++ {
		//Index starts at 1 because index 0 is the second bitmap flag

		field := int64(index + 1) // Index 1 represents field 2

		if !GetBit(bitMap, index) {
			continue //Field it not set, we can skip
		}

		extractedField, rest, err2 := extractFieldFromElements(spec, field, currentString)
		if err = err2; err != nil {
			return
		}

		currentString = rest
		//fmt.Println(currentString)

		m[field] = extractedField
		//fmt.Printf("Bitmap index: %d, Field %d: %v\n", index, field, m[field])

		if rest == "" {
			break
		}
	}

	elem = ElementsType{elements: m}
	return elem, nil
}

func NewMessage(specName string) (iso IsoStruct, err error) {
	spec, ok := specs[specName]
	if !ok {
		err = errors.New("Invalid spec")
		return
	}

	var bitMap []byte
	mti := MtiType{mti: ""}

	//We start with a single bitmap, and if we add any fields from position 65 onwards, we double the map
	bitMap = make([]byte, 8)

	emap := make(map[int64]string)
	elements := ElementsType{elements: emap}
	iso = IsoStruct{Spec: spec, Mti: mti, Bitmap: bitMap, Elements: elements}
	return
}
