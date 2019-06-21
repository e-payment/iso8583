package iso8583

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func SetConfig(specFiles []string) (err error) {
	specs = make(map[string]Spec)
	for _, specfile := range specFiles {
		if _, err = os.Stat(specfile); os.IsNotExist(err) {
			fmt.Printf("%s does not exist: %v", specfile, err)
		}

		var spec Spec
		spec, err = SpecFromFile(specfile)
		if err != nil {
			return
		}

		specs[spec.Version()] = spec
	}
	return
}

func leftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return padding + s
}

func rightPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return s + padding
}

// MtiValidator validates and iso8583 mti
func MtiValidator(mti MtiType) (bool, error) {
	mtiString := mti.mti
	len := len(mtiString)
	if len != 4 {
		err := errors.New("MTI must be length (4)")
		return false, err
	}

	_, err := strconv.ParseInt(mtiString, 10, 64)
	if err != nil {
		err := errors.New("MTI can only contain integers")
		return false, err
	}

	return true, nil
}

// FixedLengthIntegerValidator checks that an integer that is supposed to be
// of fixed width is of that length
func FixedLengthIntegerValidator(field int, length int, data string) (bool, error) {
	var verify bool
	if length != len(data) {
		return verify, fmt.Errorf("field %d: expected length %d found %d instead", field, length, len(data))
	}
	return true, nil
}

// VariableLengthIntegerValidator checks that a variable length integer field
// is within the min and max lengths specified by the spec
func VariableLengthIntegerValidator(field int, min int, max int, data string) (bool, error) {
	var verify bool
	dataLen := len(data)
	verify = (dataLen >= min) && (dataLen <= max)
	if verify == true {
		return verify, nil
	}
	return verify, fmt.Errorf("field %d: expected max length %d and min length %d found %d", field, max, min, dataLen)
}

// VariableLengthAlphaNumericValidator checks variable length alphanum Fields
// for the correct length
func VariableLengthAlphaNumericValidator(field int, min int, max int, data string) (bool, error) {
	var verify bool
	dataLen := len(data)
	verify = (dataLen >= max) && (dataLen <= max)
	if verify == true {
		return verify, nil
	}
	return verify, fmt.Errorf("field %d: expected max length %d and min length %d", field, max, min)
}

func ValidateField(index int, fieldData string) (err error) {

	return
}

func ValidateSubField(index, subIndex int, fieldData string) (err error) {
	return
}

func (s Spec) Version() string {
	return s.version
}

// readFromFile reads a yaml specfile and loads
// and iso8583 spec from it
func (s *Spec) readFromFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	type TempSpec struct {
		Version string                     `yaml:"Version"`
		Fields  map[int64]FieldDescription `yaml:"Fields"`
	}

	var temp TempSpec
	yaml.Unmarshal(content, &temp) // expecting content to be valid yaml
	if !strings.Contains(strings.Join([]string{SPEC1987, SPEC1993, SPEC2003}, ","), temp.Version) {
		return fmt.Errorf("Invalid spec version %v.", temp.Version)
	}

	if temp.Fields == nil || len(temp.Fields) == 0 {
		return fmt.Errorf("Invalid spec file %s", filename)
	}
	s.version, s.fields = temp.Version, temp.Fields
	return nil
}

// SpecFromFile returns a brand new empty spec
func SpecFromFile(filename string) (Spec, error) {
	s := Spec{}
	err := s.readFromFile(filename)
	if err != nil {
		return s, err
	}
	return s, nil
}

// Get returns the value of bit i from map m
// i = 0 gets the left most bit or most significant bit
func GetBit(m []byte, i int) bool {
	return m[i/8]&tA[i%8] != 0
}

// Set sets bit i of map m to value v.
// It doesn't check the bounds of the slice.
func SetBit(m []byte, i int) {
	index := i / 8
	bit := i % 8
	m[index] = m[index] | tA[bit]
}

func ClearBit(m []byte, i int) {
	index := i / 8
	bit := i % 8
	m[index] = m[index] & tB[bit]
}

// Len returns the length (in bits) of the provided byteslice.
// It will always be a multipile of 8 bits.
func BitMapLen(m []byte) int {
	return len(m) * 8
}
