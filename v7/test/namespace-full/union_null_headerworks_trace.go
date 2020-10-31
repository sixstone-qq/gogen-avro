// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
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

type UnionNullHeaderworksTraceTypeEnum int

const (
	UnionNullHeaderworksTraceTypeEnumHeaderworksTrace UnionNullHeaderworksTraceTypeEnum = 1
)

type UnionNullHeaderworksTrace struct {
	Null             *types.NullVal
	HeaderworksTrace *HeaderworksTrace
	UnionType        UnionNullHeaderworksTraceTypeEnum
}

func writeUnionNullHeaderworksTrace(r *UnionNullHeaderworksTrace, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullHeaderworksTraceTypeEnumHeaderworksTrace:
		return writeHeaderworksTrace(r.HeaderworksTrace, w)
	}
	return fmt.Errorf("invalid value for *UnionNullHeaderworksTrace")
}

func NewUnionNullHeaderworksTrace() *UnionNullHeaderworksTrace {
	return &UnionNullHeaderworksTrace{}
}

func (r *UnionNullHeaderworksTrace) Serialize(w io.Writer) error {
	return writeUnionNullHeaderworksTrace(r, w)
}

func DeserializeUnionNullHeaderworksTrace(r io.Reader) (*UnionNullHeaderworksTrace, error) {
	t := NewUnionNullHeaderworksTrace()
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

func (r *UnionNullHeaderworksTrace) Schema() string {
	return "[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"Trace\",\"type\":\"record\"}]"
}

func (_ *UnionNullHeaderworksTrace) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullHeaderworksTrace) SetLong(v int64) {
	r.UnionType = (UnionNullHeaderworksTraceTypeEnum)(v)
}
func (r *UnionNullHeaderworksTrace) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.HeaderworksTrace = NewHeaderworksTrace()
		return r.HeaderworksTrace
	}
	panic("Unknown field index")
}
func (_ *UnionNullHeaderworksTrace) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksTrace) Finalize()                        {}

func (r *UnionNullHeaderworksTrace) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionNullHeaderworksTraceTypeEnumHeaderworksTrace:
		return json.Marshal(map[string]interface{}{"headerworks.Trace": r.HeaderworksTrace})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullHeaderworksTrace")
}

func (r *UnionNullHeaderworksTrace) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["headerworks.Trace"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.HeaderworksTrace)
	}
	return fmt.Errorf("invalid value for *UnionNullHeaderworksTrace")
}
