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

func NewRecordFooWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewRecordFoo()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type RecordFooReader struct {
	r io.Reader
	p *vm.Program
}

func NewRecordFooReader(r io.Reader) (*RecordFooReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewRecordFoo()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &RecordFooReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r RecordFooReader) Read() (*RecordFoo, error) {
	t := NewRecordFoo()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
