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

type Parent struct {
	Children []*Child `json:"Children"`
}

const ParentAvroCRC64Fingerprint = "p\x9fn\xf8\x11i\x98\x9b"

func NewParent() *Parent {
	return &Parent{}
}

func DeserializeParent(r io.Reader) (*Parent, error) {
	t := NewParent()
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

func DeserializeParentFromSchema(r io.Reader, schema string) (*Parent, error) {
	t := NewParent()

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

func writeParent(r *Parent, w io.Writer) error {
	var err error
	err = writeArrayChild(r.Children, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Parent) Serialize(w io.Writer) error {
	return writeParent(r, w)
}

func (r *Parent) Schema() string {
	return "{\"fields\":[{\"default\":[{\"name\":\"record1\"},{\"name\":\"record2\"}],\"name\":\"Children\",\"type\":{\"items\":{\"fields\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Child\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Parent\",\"type\":\"record\"}"
}

func (r *Parent) SchemaName() string {
	return "Parent"
}

func (_ *Parent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Parent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Parent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Parent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Parent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Parent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Parent) SetString(v string)   { panic("Unsupported operation") }
func (_ *Parent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Parent) Get(i int) types.Field {
	switch i {
	case 0:
		r.Children = make([]*Child, 0)

		return &ArrayChildWrapper{Target: &r.Children}
	}
	panic("Unknown field index")
}

func (r *Parent) SetDefault(i int) {
	switch i {
	case 0:
		r.Children = make([]*Child, 2)
		r.Children[0] = NewChild()
		r.Children[0].Name = "record1"

		r.Children[1] = NewChild()
		r.Children[1].Name = "record2"

		return
	}
	panic("Unknown field index")
}

func (r *Parent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Parent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Parent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Parent) Finalize()                        {}

func (_ *Parent) AvroCRC64Fingerprint() []byte {
	return []byte(ParentAvroCRC64Fingerprint)
}
