package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type UnionNullBytesTypeEnum int

const (
	UnionNullBytesTypeEnumBytes UnionNullBytesTypeEnum = 1
)

type UnionNullBytes struct {
	Null      *types.NullVal
	Bytes     []byte
	UnionType UnionNullBytesTypeEnum
}

func writeUnionNullBytes(r *UnionNullBytes, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBytesTypeEnumBytes:
		return vm.WriteBytes(r.Bytes, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBytes")
}

func NewUnionNullBytes() *UnionNullBytes {
	return &UnionNullBytes{}
}

func (_ *UnionNullBytes) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullBytes) SetLong(v int64) {
	r.UnionType = (UnionNullBytesTypeEnum)(v)
}
func (r *UnionNullBytes) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		return &types.Bytes{Target: (&r.Bytes)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullBytes) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBytes) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBytes) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBytes) Finalize()                        {}

func (r *UnionNullBytes) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionNullBytesTypeEnumBytes:
		return json.Marshal(map[string]interface{}{"bytes": r.Bytes})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBytes")
}

func (r *UnionNullBytes) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["bytes"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Bytes)
	}
	return fmt.Errorf("invalid value for *UnionNullBytes")
}
