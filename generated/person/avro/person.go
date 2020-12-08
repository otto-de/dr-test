package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

// A person
type Person struct {
	// The name of the person
	Name string `json:"name"`
	// The person's age in years
	Age int32 `json:"age"`
	// The person's height in meters (attention it's a float)
	Height float32 `json:"height"`
	// The person's account balance in cent (positive or negative)
	AccountBalance int64 `json:"accountBalance"`

	IsFemale bool `json:"isFemale"`
	// The (optional) bytes of a person's avatar
	ImageData *UnionNullBytes `json:"imageData"`

	Location *Location `json:"location"`
	// The person's hobbies
	Hobbies []string `json:"hobbies"`
}

const PersonAvroCRC64Fingerprint = "\xd7+K2\xac\f\xce\xf0"

func NewPerson() *Person {
	return &Person{}
}

func DeserializePerson(r io.Reader) (*Person, error) {
	t := NewPerson()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializePersonFromSchema(r io.Reader, schema string) (*Person, error) {
	t := NewPerson()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writePerson(r *Person, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Age, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Height, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.AccountBalance, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.IsFemale, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBytes(r.ImageData, w)
	if err != nil {
		return err
	}
	err = writeLocation(r.Location, w)
	if err != nil {
		return err
	}
	err = writeArrayString(r.Hobbies, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Person) Serialize(w io.Writer) error {
	return writePerson(r, w)
}

func (r *Person) Schema() string {
	return "{\"doc\":\"A person\",\"fields\":[{\"doc\":\"The name of the person\",\"name\":\"name\",\"type\":\"string\"},{\"doc\":\"The person's age in years\",\"name\":\"age\",\"type\":\"int\"},{\"doc\":\"The person's height in meters (attention it's a float)\",\"name\":\"height\",\"type\":\"float\"},{\"doc\":\"The person's account balance in cent (positive or negative)\",\"name\":\"accountBalance\",\"type\":\"long\"},{\"name\":\"isFemale\",\"type\":\"boolean\"},{\"default\":null,\"doc\":\"The (optional) bytes of a person's avatar\",\"name\":\"imageData\",\"type\":[\"null\",\"bytes\"]},{\"name\":\"location\",\"type\":{\"fields\":[{\"doc\":\"The latitude of the location\",\"name\":\"lat\",\"type\":\"double\"},{\"doc\":\"The longitude of the location\",\"name\":\"lon\",\"type\":\"double\"}],\"name\":\"location\",\"type\":\"record\"}},{\"doc\":\"The person's hobbies\",\"name\":\"hobbies\",\"type\":{\"items\":\"string\",\"type\":\"array\"}}],\"name\":\"de.otto.dr.test.Person\",\"type\":\"record\"}"
}

func (r *Person) SchemaName() string {
	return "de.otto.dr.test.Person"
}

func (_ *Person) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Person) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Person) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Person) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Person) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Person) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Person) SetString(v string)   { panic("Unsupported operation") }
func (_ *Person) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Person) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Name}
	case 1:
		return &types.Int{Target: &r.Age}
	case 2:
		return &types.Float{Target: &r.Height}
	case 3:
		return &types.Long{Target: &r.AccountBalance}
	case 4:
		return &types.Boolean{Target: &r.IsFemale}
	case 5:
		r.ImageData = NewUnionNullBytes()

		return r.ImageData
	case 6:
		r.Location = NewLocation()

		return r.Location
	case 7:
		r.Hobbies = make([]string, 0)

		return &ArrayStringWrapper{Target: &r.Hobbies}
	}
	panic("Unknown field index")
}

func (r *Person) SetDefault(i int) {
	switch i {
	case 5:
		r.ImageData = nil
		return
	}
	panic("Unknown field index")
}

func (r *Person) NullField(i int) {
	switch i {
	case 5:
		r.ImageData = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ *Person) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Person) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Person) Finalize()                        {}

func (_ *Person) AvroCRC64Fingerprint() []byte {
	return []byte(PersonAvroCRC64Fingerprint)
}
