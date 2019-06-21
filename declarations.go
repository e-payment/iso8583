package iso8583

const (
	SPEC1987 = "1987"
	SPEC1993 = "1993"
	SPEC2003 = "2003"
)

var (
	specs = map[string]Spec{}

	//The maps below are used to perform bitwise operations on bitmaps, to detect if a certain bit is set or not
	//10000000, 01000000, 00100000, 00010000, 00001000, 00000100, 00000010, 00000001
	tA = [8]byte{128, 64, 32, 16, 8, 4, 2, 1}

	//01111111, 10111111, 11011111, 11101111, 11110111, 11111011, 11111101, 11111110
	tB = [8]byte{127, 191, 223, 239, 247, 251, 253, 254}
)

type (
	// IsoStruct is an iso8583 container
	Message struct {
		spec   Spec
		mti    string
		bitmap []byte
		data   map[int]Field
	}

	Field struct {
		ID        string
		Value     string //the decoded value that can be used as is
		Subfields map[int]Field
	}

	// FieldDescription contains fields that describes an iso8583 Field
	FieldDescription struct {
		ContentType string                   `yaml:"ContentType"`
		MaxLen      int                      `yaml:"MaxLen"`
		MinLen      int                      `yaml:"MinLen"`
		LenType     string                   `yaml:"LenType"`
		Label       string                   `yaml:"Label"`
		Format      string                   `yaml:"Format"`
		Subfields   map[int]FieldDescription `yaml:"Subfields"`
	}

	// Spec contains a strutured description of an iso8583 spec
	// properly defined by a spec file
	Spec struct {
		version string
		fields  map[int]FieldDescription
	}

	// ValidationError happens when validation fails
	ValidationError struct {
		message string
	}
)
