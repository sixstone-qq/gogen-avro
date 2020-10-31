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

type UnionNullBodyworksDataTypeEnum int

const (
	UnionNullBodyworksDataTypeEnumBodyworksData UnionNullBodyworksDataTypeEnum = 1
)

type UnionNullBodyworksData struct {
	Null          *types.NullVal
	BodyworksData *BodyworksData
	UnionType     UnionNullBodyworksDataTypeEnum
}

func writeUnionNullBodyworksData(r *UnionNullBodyworksData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBodyworksDataTypeEnumBodyworksData:
		return writeBodyworksData(r.BodyworksData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksData")
}

func NewUnionNullBodyworksData() *UnionNullBodyworksData {
	return &UnionNullBodyworksData{}
}

func (r *UnionNullBodyworksData) Serialize(w io.Writer) error {
	return writeUnionNullBodyworksData(r, w)
}

func DeserializeUnionNullBodyworksData(r io.Reader) (*UnionNullBodyworksData, error) {
	t := NewUnionNullBodyworksData()
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

func (r *UnionNullBodyworksData) Schema() string {
	return "[\"null\",{\"doc\":\"Common information related to the event which must be included in any clean event\",\"fields\":[{\"default\":null,\"doc\":\"Unique identifier for the event used for de-duplication and tracing.\",\"name\":\"uuid\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"bodyworks.datatype\",\"type\":\"record\"}]},{\"default\":null,\"doc\":\"Fully qualified name of the host that generated the event that generated the data.\",\"name\":\"hostname\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"Trace information not redundant with this object\",\"name\":\"trace\",\"type\":[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"Trace\",\"type\":\"record\"}]}],\"name\":\"Data\",\"namespace\":\"bodyworks\",\"type\":\"record\"}]"
}

func (_ *UnionNullBodyworksData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullBodyworksData) SetLong(v int64) {
	r.UnionType = (UnionNullBodyworksDataTypeEnum)(v)
}
func (r *UnionNullBodyworksData) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.BodyworksData = NewBodyworksData()
		return r.BodyworksData
	}
	panic("Unknown field index")
}
func (_ *UnionNullBodyworksData) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBodyworksData) Finalize()                        {}

func (r *UnionNullBodyworksData) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionNullBodyworksDataTypeEnumBodyworksData:
		return json.Marshal(map[string]interface{}{"bodyworks.Data": r.BodyworksData})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBodyworksData")
}

func (r *UnionNullBodyworksData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["bodyworks.Data"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.BodyworksData)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksData")
}
