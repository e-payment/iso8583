package iso8583

import (
	"testing"
)

func TestLeftPad(t *testing.T) {
	ts := []struct {
		s, expected string
	}{
		{leftPad("foo", 5, "0"), "00foo"},
		{leftPad("foobar", 6, " "), "foobar"},
		{leftPad("1", 2, "0"), "01"},
	}
	for _, v := range ts {
		if v.expected != v.s {
			t.Errorf("Expected %s but got %s", v.expected, v.s)
		}
	}
}

func TestRightPad(t *testing.T) {
	ts := []struct {
		s, expected string
	}{
		{rightPad("foo", 5, "0"), "foo00"},
		{rightPad("foobar", 6, " "), "foobar"},
		{rightPad("1", 2, "0"), "10"},
	}
	for _, v := range ts {
		if v.expected != v.s {
			t.Errorf("Expected %s but got %s", v.expected, v.s)
		}
	}
}

func Test_ReadFile(t *testing.T) {
	if _, err := SpecFromFile("specs/spec1987.yml"); err != nil {
		t.Errorf("failed to parse valid spec file %s:  %s", "spec1987.yml", err.Error())
	}

	if _, err := SpecFromFile("specs/spec1993.yml"); err != nil {
		t.Errorf("failed to parse valid spec file %s:  %s", "spec1993.yml", err.Error())
	}

	if _, err := SpecFromFile("specs/spec2003.yml"); err != nil {
		t.Errorf("failed to parse valid spec file %s:  %s", "spec2003.yml", err.Error())
	}

	if _, err := SpecFromFile("specs/spec2003full.yml"); err != nil {
		t.Errorf("failed to parse valid spec file %s:  %s", "spec2003full.yml", err.Error())
	}
}

func TestMtiValidator(t *testing.T) {
	if err := ValidateMti("0200"); err != nil {
		t.Errorf("failed to verify a valid mti")
	}

	if err := ValidateMti("020"); err == nil {
		t.Errorf("did not detect an invalid length mti")
	}

	if err := ValidateMti("a0200"); err == nil {
		t.Errorf("did not detect an invalid character in mti")
	}
}
