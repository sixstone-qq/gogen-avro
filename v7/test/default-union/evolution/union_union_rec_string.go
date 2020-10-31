// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type UnionUnionRecStringTypeEnum int

const (
	UnionUnionRecStringTypeEnumUnionRec UnionUnionRecStringTypeEnum = 0

	UnionUnionRecStringTypeEnumString UnionUnionRecStringTypeEnum = 1
)

type UnionUnionRecString struct {
	UnionRec  *UnionRec
	String    string
	UnionType UnionUnionRecStringTypeEnum
}

func writeUnionUnionRecString(r *UnionUnionRecString, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionUnionRecStringTypeEnumUnionRec:
		return writeUnionRec(r.UnionRec, w)
	case UnionUnionRecStringTypeEnumString:
		return vm.WriteString(r.String, w)
	}
	return fmt.Errorf("invalid value for *UnionUnionRecString")
}

func NewUnionUnionRecString() *UnionUnionRecString {
	return &UnionUnionRecString{}
}

func (r *UnionUnionRecString) Serialize(w io.Writer) error {
	return writeUnionUnionRecString(r, w)
}

func DeserializeUnionUnionRecString(r io.Reader) (*UnionUnionRecString, error) {
	t := NewUnionUnionRecString()
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

func (r *UnionUnionRecString) Schema() string {
	return "[{\"fields\":[{\"name\":\"a\",\"type\":\"int\"}],\"name\":\"unionRec\",\"type\":\"record\"},\"string\"]"
}

func (_ *UnionUnionRecString) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionUnionRecString) SetLong(v int64) {
	r.UnionType = (UnionUnionRecStringTypeEnum)(v)
}
func (r *UnionUnionRecString) Get(i int) types.Field {
	switch i {
	case 0:
		r.UnionRec = NewUnionRec()
		return r.UnionRec
	case 1:
		return &types.String{Target: (&r.String)}
	}
	panic("Unknown field index")
}
func (_ *UnionUnionRecString) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionUnionRecString) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionUnionRecString) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionUnionRecString) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionUnionRecString) Finalize()                        {}

func (r *UnionUnionRecString) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionUnionRecStringTypeEnumUnionRec:
		return json.Marshal(map[string]interface{}{"unionRec": r.UnionRec})
	case UnionUnionRecStringTypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	}
	return nil, fmt.Errorf("invalid value for *UnionUnionRecString")
}

func (r *UnionUnionRecString) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["unionRec"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.UnionRec)
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.String)
	}
	return fmt.Errorf("invalid value for *UnionUnionRecString")
}
