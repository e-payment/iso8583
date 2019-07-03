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
func (iso *Message) ToString() (string, error) {
	var str string
	// get done with the mti and the bitmap
	bitmapString := strings.ToUpper(hex.EncodeToString(iso.Bitmap()))

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
func (iso *Message) AddMTI(data string) error {
	if err := ValidateMti(data); err != nil {
		return err
	}

	iso.mti = data
	//fmt.Printf("MTI: %s\n", iso.Mti)
	return nil
}
func (iso *Message) AddTransactionAmount(currency, amount, decimals int64) (err error) {
	var amountField string

	if len(fmt.Sprint(currency)) > 3 || len(fmt.Sprint(amount)) > 12 || len(fmt.Sprint(decimals)) > 1 {
		return errors.New("Amounts are too long for field size")
	}

	switch iso.spec.version {
	case "1987", "1993":
		amountField = fmt.Sprintf("%012d", amount)
	case "2003":
		amountField = fmt.Sprintf("%03d%d%012d", currency, decimals, amount)
	}

	return iso.AddField(BIT87_TRANSACTION_AMOUNT, amountField)
}

func (iso *Message) UpdateTransactionAmount(amount int64) (err error) {
	if len(fmt.Sprint(amount)) > 12 {
		return errors.New("Amount is too long for field size")
	}

	amountField, err := iso.GetField(BIT87_TRANSACTION_AMOUNT)
	if err != nil {
		return
	}

	amountStr := amountField.Value

	switch len(amountStr) {
	case 12:
		amountStr = fmt.Sprint(amount)
	case 16:
		amountStr = amountStr[:4] + fmt.Sprint(amount)
	default:
		return errors.New("Invalid amount field length")
	}

	return iso.AddField(BIT87_TRANSACTION_AMOUNT, amountStr)

}

func (iso *Message) AddOriginalTransactionAmount(currency, amount, decimals int64) (err error) {
	var amountField string

	if len(fmt.Sprint(currency)) > 3 || len(fmt.Sprint(amount)) > 12 || len(fmt.Sprint(decimals)) > 1 {
		return errors.New("Amounts are too long for field size")
	}

	switch iso.spec.version {
	case "1987", "1993":
		amountField = fmt.Sprintf("%012d", amount)
	case "2003":
		amountField = fmt.Sprintf("%03d%d%012d", currency, decimals, amount)
	}

	return iso.AddField(BIT87_ORIGINAL_AMOUNT, amountField)
}

func (iso *Message) UpdateOriginalTransactionAmount(amount int64) (err error) {
	if len(fmt.Sprint(amount)) > 12 {
		return errors.New("Amount is too long for field size")
	}

	amountField, err := iso.GetField(BIT87_ORIGINAL_AMOUNT)
	if err != nil {
		return
	}

	amountStr := amountField.Value

	switch len(amountStr) {
	case 12:
		amountStr = fmt.Sprint(amount)
	case 16:
		amountStr = amountStr[:4] + fmt.Sprint(amount)
	default:
		return errors.New("Invalid amount field length")
	}

	return iso.AddField(BIT87_ORIGINAL_AMOUNT, amountStr)

}

func (iso *Message) ResponseStr(messageClass, responseCode string) (response string, err error) {
	if err = ValidateMti(iso.mti); err != nil {
		return
	}

	new := Message{spec: iso.spec, bitmap: iso.bitmap, data: iso.data}

	if err = new.AddMTI(string(iso.mti[0]) + messageClass); err != nil {
		return
	}

	if responseCode != "" {
		if err = new.AddField(BIT87_ACTION_CODE, responseCode); err != nil {
			return
		}
	}

	if messageClass == AUTH_ADVICE_RESP || messageClass == AUTH_REQ_RESP || messageClass == AUTH_NOTIFY_RESP &&
		!(responseCode == RC_APPROVED || responseCode == RC87_APPROVED) {
		new.UpdateTransactionAmount(0)
		new.UpdateOriginalTransactionAmount(0)

	}

	return new.ToString()
}

func (iso *Message) Response(messageClass, responseCode string) (response Message, err error) {
	if err = ValidateMti(iso.mti); err != nil {
		return
	}

	response = Message{spec: iso.spec, bitmap: iso.bitmap, data: iso.data}

	if err = response.AddMTI(string(iso.mti[0]) + messageClass); err != nil {
		return
	}

	if responseCode != "" {
		if err = response.AddField(BIT87_ACTION_CODE, responseCode); err != nil {
			return
		}
	}

	return
}

// AddField adds the provided iso8583 field into the current struct
// also updates the bitmap in the process
func (iso *Message) AddField(field int, data string) error {
	if field < 2 || field > 128 {
		return fmt.Errorf("expected field to be between %d and %d found %d instead", 2, 128, field)
	}

	if field > 64 {
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
func Parse(i string) (iso Message, err error) {
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

	iso = Message{spec: spec, mti: mti, bitmap: bitMap, data: data}
	return iso, nil
}

func (iso *Message) packElements() (string, error) {
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
	mapLen := 16 //becasue each hex character is 4 bits, not a whole byte, so we get twice the characters
	if len(rest) < mapLen {
		err = fmt.Errorf("Bitmap length of %d is greater than the data length of %d", mapLen, len(rest))
		return
	}

	//only 64 bit primary bitmap
	if bytes, err = hex.DecodeString(rest[:mapLen]); err != nil {
		err = fmt.Errorf("Error Decoding string: %v", err)
		return
	}

	if GetBit(bytes, 0) {
		mapLen = 32
		if len(rest) < mapLen {
			err = fmt.Errorf("Bitmap length of %d is greater than the data length of %d", mapLen, len(rest))
			return
		}

		// secondary bitmap exists, so total 128 bits
		if bytes, err = hex.DecodeString(rest[:mapLen]); err != nil {
			err = fmt.Errorf("Error Decoding string: %v", err)
			return
		}
	}
	fmt.Printf("Bitmap length: %d, Bitmap: %08b, Data: %s\n", BitMapLen(bytes), bytes, hex.EncodeToString(bytes))

	elementsString = rest[mapLen:]
	fmt.Printf("Bitmap length: %d\n", BitMapLen(bytes))
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
		err = fmt.Errorf("Fixed length Field %d of %d is longer than data string %s", field, fieldDescription.MaxLen, str)
		return
	}

	result.ID = fieldID
	fmt.Printf("Index: %d, Description: %v, Data: %s\n", field, fieldDescription, str)
	if fieldDescription.LenType == "fixed" {
		result.Value = str[:fieldLength]
		rest = strings.TrimPrefix(str, result.Value)
		fmt.Printf("Index: %d, Description: %v, Extracted: %s, Data: %s\n", field, fieldDescription, result.Value, str)
	} else {
		// varianle length fields have their lengths embedded into the string
		digitCount, err2 := getVariableLengthFromString(fieldDescription.LenType)
		if err2 != nil {
			err = fmt.Errorf("spec error: field %d: %s", field, err2.Error())
			return
		}

		fieldLengthStr := str[0:digitCount]
		fmt.Println("Digits: " + fieldLengthStr)

		fieldLength, err = strconv.Atoi(fieldLengthStr)
		if err != nil {
			err = fmt.Errorf("spec error: field %d: %s", field, "invalid length digits")
			return
		}

		rest = strings.TrimPrefix(str, fieldLengthStr)

		if fieldLength > len(rest) {
			err = fmt.Errorf("Field %d's length of %d is longer than data string %s", field, fieldLength, rest)
			return
		}

		result.Value = rest[:fieldLength]
		rest = strings.TrimPrefix(rest, result.Value)
		fmt.Printf("Index: %d, Description: %v, Extracted: %s, Data: %s\n", field, fieldDescription, result.Value, rest)
	}

	if err = result.Validate(); err != nil {
		return
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

func NewMessage(specName string) (iso Message, err error) {
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
	iso = Message{spec: spec, mti: mti, bitmap: bitMap, data: elements}
	return
}
