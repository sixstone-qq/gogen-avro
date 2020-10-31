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

type NestedRecord struct {
	StringField string `json:"StringField"`

	BoolField bool `json:"BoolField"`

	BytesField Bytes `json:"BytesField"`
}

const NestedRecordAvroCRC64Fingerprint = "\x81\x8d\xc3K?\xe83\xcc"

func NewNestedRecord() *NestedRecord {
	return &NestedRecord{}
}

func DeserializeNestedRecord(r io.Reader) (*NestedRecord, error) {
	t := NewNestedRecord()
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

func DeserializeNestedRecordFromSchema(r io.Reader, schema string) (*NestedRecord, error) {
	t := NewNestedRecord()

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

func writeNestedRecord(r *NestedRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.StringField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.BoolField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.BytesField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *NestedRecord) Serialize(w io.Writer) error {
	return writeNestedRecord(r, w)
}

func (r *NestedRecord) Schema() string {
	return "{\"aliases\":[\"aliasedRecord\"],\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"NestedRecord\",\"type\":\"record\"}"
}

func (r *NestedRecord) SchemaName() string {
	return "NestedRecord"
}

func (_ *NestedRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *NestedRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *NestedRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *NestedRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *NestedRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *NestedRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *NestedRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *NestedRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedRecord) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.StringField}
	case 1:
		return &types.Boolean{Target: &r.BoolField}
	case 2:
		return &BytesWrapper{Target: &r.BytesField}
	}
	panic("Unknown field index")
}

func (r *NestedRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NestedRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *NestedRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *NestedRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *NestedRecord) Finalize()                        {}

func (_ *NestedRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NestedRecordAvroCRC64Fingerprint)
}
