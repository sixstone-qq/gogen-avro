// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type UnionRecord struct {
	Id string `json:"id"`

	Age int32 `json:"age"`
}

const UnionRecordAvroCRC64Fingerprint = "XXv\f\x93)\xf7\xe2"

func NewUnionRecord() *UnionRecord {
	return &UnionRecord{}
}

func DeserializeUnionRecord(r io.Reader) (*UnionRecord, error) {
	t := NewUnionRecord()
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

func DeserializeUnionRecordFromSchema(r io.Reader, schema string) (*UnionRecord, error) {
	t := NewUnionRecord()

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

func writeUnionRecord(r *UnionRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Age, w)
	if err != nil {
		return err
	}
	return err
}

func (r *UnionRecord) Serialize(w io.Writer) error {
	return writeUnionRecord(r, w)
}

func (r *UnionRecord) Schema() string {
	return "{\"fields\":[{\"default\":\"test_id\",\"name\":\"id\",\"type\":\"string\"},{\"default\":100,\"name\":\"age\",\"type\":\"int\"}],\"name\":\"UnionRecord\",\"type\":\"record\"}"
}

func (r *UnionRecord) SchemaName() string {
	return "UnionRecord"
}

func (_ *UnionRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *UnionRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *UnionRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *UnionRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *UnionRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *UnionRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *UnionRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *UnionRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *UnionRecord) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Id}
	case 1:
		return &types.Int{Target: &r.Age}
	}
	panic("Unknown field index")
}

func (r *UnionRecord) SetDefault(i int) {
	switch i {
	case 0:
		r.Id = "test_id"
		return
	case 1:
		r.Age = 100
		return
	}
	panic("Unknown field index")
}

func (r *UnionRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *UnionRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionRecord) Finalize()                        {}

func (_ *UnionRecord) AvroCRC64Fingerprint() []byte {
	return []byte(UnionRecordAvroCRC64Fingerprint)
}
