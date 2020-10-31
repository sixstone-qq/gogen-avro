// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/vm"
)

func NewEnumTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewEnumTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type EnumTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewEnumTestRecordReader(r io.Reader) (*EnumTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewEnumTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &EnumTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r EnumTestRecordReader) Read() (*EnumTestRecord, error) {
	t := NewEnumTestRecord()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
