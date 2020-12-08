package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type Location struct {
	// The latitude of the location
	Lat float64 `json:"lat"`
	// The longitude of the location
	Lon float64 `json:"lon"`
}

const LocationAvroCRC64Fingerprint = "y\xb1B\xc6A6\x1c\xf2"

func NewLocation() *Location {
	return &Location{}
}

func DeserializeLocation(r io.Reader) (*Location, error) {
	t := NewLocation()
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

func DeserializeLocationFromSchema(r io.Reader, schema string) (*Location, error) {
	t := NewLocation()

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

func writeLocation(r *Location, w io.Writer) error {
	var err error
	err = vm.WriteDouble(r.Lat, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Lon, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Location) Serialize(w io.Writer) error {
	return writeLocation(r, w)
}

func (r *Location) Schema() string {
	return "{\"fields\":[{\"doc\":\"The latitude of the location\",\"name\":\"lat\",\"type\":\"double\"},{\"doc\":\"The longitude of the location\",\"name\":\"lon\",\"type\":\"double\"}],\"name\":\"de.otto.dr.test.location\",\"type\":\"record\"}"
}

func (r *Location) SchemaName() string {
	return "de.otto.dr.test.location"
}

func (_ *Location) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Location) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Location) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Location) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Location) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Location) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Location) SetString(v string)   { panic("Unsupported operation") }
func (_ *Location) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Location) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Double{Target: &r.Lat}
	case 1:
		return &types.Double{Target: &r.Lon}
	}
	panic("Unknown field index")
}

func (r *Location) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Location) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Location) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Location) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Location) Finalize()                        {}

func (_ *Location) AvroCRC64Fingerprint() []byte {
	return []byte(LocationAvroCRC64Fingerprint)
}
