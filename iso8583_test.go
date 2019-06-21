package iso8583

import (
	"fmt"
	"testing"
)

func TestISOParse(t *testing.T) {
	// MTI = 0200
	// Bitmap = 3220000000808000 = 0011001000100000000000000000000000000000100000001000000000000000
	// Field (3) = 000010
	// Field (4) = 000000001500
	// Field (7) = 1206041200
	// Field (11) = 000001
	// Field (41) = 12340001
	// Field (49) = 840

	err := SetConfig([]string{"specs/spec1987.yml", "specs/spec1993.yml", "specs/spec2003full.yml"})
	if err != nil {
		fmt.Println(err)
		t.Errorf("SetConfig failed: %v", err)
	}

	//0200 3220000000808000 000010 000000001500 1206041200 000001 12340001 840
	isomsg := "02003220000000808000000010000000001500120604120000000112340001840"

	parsed, err := Parse(isomsg)
	if err != nil {
		fmt.Println(err)
		t.Errorf("parse iso message failed: %v", err)
	}

	isomsgUnpacked, err := parsed.ToString()
	if err != nil {
		fmt.Println(err)
		t.Errorf("failed to unpack valid isomsg")
	}
	if isomsgUnpacked != isomsg {
		t.Errorf("%s should be %s", isomsgUnpacked, isomsg)
	}
	// fmt.Printf("%#v, %#v\n%#v", parsed.Mti, parsed.Bitmap, parsed.Elements)
}

func TestEmpty(t *testing.T) {
	err := SetConfig([]string{"specs/spec1987.yml", "specs/spec1993.yml", "specs/spec2003full.yml"})
	if err != nil {
		fmt.Println(err)
		t.Errorf("SetConfig failed: %v", err)
	}

	one, err := NewMessage("2003")
	if err != nil {
		fmt.Println(err)
		t.Errorf("failed to unpack valid isomsg")
	}

	if one.Mti.String() != "" {
		t.Errorf("Empty generates invalid MTI")
	}
	if err = one.AddMTI("0200"); err != nil {
		t.Errorf("Failed adding MTI")
	}

	if err = one.AddField(3, "000010"); err != nil {
		t.Errorf("Failed adding field: %s", err.Error())
	}
	if err = one.AddField(4, "000000001500"); err != nil {
		t.Errorf("Failed adding field: %s", err.Error())
	}
	if err = one.AddField(7, "1206041200"); err != nil {
		t.Errorf("Failed adding field: %s", err.Error())
	}
	if err = one.AddField(11, "000001"); err != nil {
		t.Errorf("Failed adding field: %s", err.Error())
	}
	if err = one.AddField(41, "12340001"); err != nil {
		t.Errorf("Failed adding field: %s", err.Error())
	}
	if err = one.AddField(49, "840"); err != nil {
		t.Errorf("Failed adding field: %s", err.Error())
	}

	expected := "02003220000000808000000010000000001500120604120000000112340001840"
	unpacked, _ := one.ToString()
	if unpacked != expected {
		t.Errorf("Manually constructed isostruct produced %s not %s", unpacked, expected)
	}
}
