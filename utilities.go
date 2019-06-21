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

func (iso IsoStruct) Spec() Spec {
	return iso.spec
}
func (iso IsoStruct) Mti() string {
	return iso.mti
}
func (iso IsoStruct) Bitmap() []byte {
	return iso.bitmap
}

func (iso IsoStruct) Data() map[int]Field {
	return iso.data
}

//returns a field description for a field
//field id must be in the format spec.index.subindex.subindex...
//Example: 2003.30.1.2
func DescribeField(fieldID string) (desc FieldDescription, indices []int, level int, err error) {
	specName, indices, level, err := DecodeFieldID(fieldID)
	if err != nil {
		return
	}

	//Get spec
	spec, ok := specs[specName]
	if !ok {
		err = fmt.Errorf("Spect %s is not supported", specName)
		return
	}

	//Get top field
	desc, ok = spec.fields[indices[0]]
	if !ok {
		err = fmt.Errorf("Spec %s does not contain field %d", spec.Version(), indices[0])
		return
	}

	if len(indices) == 1 {
		return
	}

	//Get Subfield
	desc, ok = desc.Subfields[indices[1]]
	if !ok {
		err = fmt.Errorf("Spec %s does not contain field %s", spec.Version(), fieldID)
		return
	}

	if len(indices) == 2 {
		return
	}

	//Get sub subfield
	desc, ok = desc.Subfields[indices[2]]
	if !ok {
		err = fmt.Errorf("Spec %s does not contain field %s", spec.Version(), fieldID)
		return
	}

	return
}

func GetFieldSpec(fieldID string) (spec Spec, err error) {
	specName, _, _, err := DecodeFieldID(fieldID)
	if err != nil {
		return
	}

	spec, ok := specs[specName]
	if !ok {
		err = fmt.Errorf("Spect %s is not supported", specName)
		return
	}

	return
}

func DecodeFieldID(fieldID string) (spec string, indices []int, level int, err error) {
	data := strings.Split(fieldID, ".")
	if len(data) < 2 || len(data) > 4 { //specs only have 4 levels
		err = errors.New("Invalid field ID")
		return
	}

	if data[0] != SPEC1987 && data[0] != SPEC1993 && data[0] != SPEC2003 {
		err = errors.New("Invalid field SPEC. Must be 1987, 1997 or 2003")
		return
	}
	spec = data[0]

	level = len(data) - 2
	for _, index := range data[1:] {
		val, err2 := strconv.Atoi(index)
		if err2 != nil {
			return
		}
		indices = append(indices, val)
	}

	if level >= len(indices) || level < 0 {
		err = errors.New("Invlaid field level.")
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
func ValidateMti(mti string) error {
	len := len(mti)
	if len != 4 {
		return errors.New("MTI must be length (4)")
	}

	_, err := strconv.ParseInt(mti, 10, 64)
	if err != nil {
		return errors.New("MTI can only contain integers")
	}

	return nil
}

func validateField(data string, desc FieldDescription) (err error) {
	return
}

func validateSubField(data string, subIndex int, desc FieldDescription) (err error) {
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
		Version string                   `yaml:"Version"`
		Fields  map[int]FieldDescription `yaml:"Fields"`
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
