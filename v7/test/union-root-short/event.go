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

type Event struct {
	// Unique ID for this event.
	Id string `json:"id"`
	// Start IP of this observation's IP range.
	Start_ip Ip_address `json:"start_ip"`
	// End IP of this observation's IP range.
	End_ip Ip_address `json:"end_ip"`
}

const EventAvroCRC64Fingerprint = "\xebZ\xc0m\xf9OV\x97"

func NewEvent() *Event {
	return &Event{}
}

func DeserializeEvent(r io.Reader) (*Event, error) {
	t := NewEvent()
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

func DeserializeEventFromSchema(r io.Reader, schema string) (*Event, error) {
	t := NewEvent()

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

func writeEvent(r *Event, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeIp_address(r.Start_ip, w)
	if err != nil {
		return err
	}
	err = writeIp_address(r.End_ip, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Event) Serialize(w io.Writer) error {
	return writeEvent(r, w)
}

func (r *Event) Schema() string {
	return "{\"fields\":[{\"doc\":\"Unique ID for this event.\",\"metadata\":{\"key\":\"field1\"},\"name\":\"id\",\"type\":\"string\"},{\"doc\":\"Start IP of this observation's IP range.\",\"metadata\":{\"key\":\"field2\"},\"name\":\"start_ip\",\"type\":{\"key\":\"value\",\"metadata\":{\"a\":\"b\",\"c\":123},\"name\":\"ip_address\",\"size\":16,\"type\":\"fixed\"}},{\"doc\":\"End IP of this observation's IP range.\",\"metadata\":{\"key\":\"field3\"},\"name\":\"end_ip\",\"type\":\"ip_address\"}],\"metadata\":{\"key\":\"value\"},\"name\":\"event\",\"subject\":\"event\",\"type\":\"record\"}"
}

func (r *Event) SchemaName() string {
	return "event"
}

func (_ *Event) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Event) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Event) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Event) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Event) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Event) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Event) SetString(v string)   { panic("Unsupported operation") }
func (_ *Event) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Event) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Id}
	case 1:
		return &Ip_addressWrapper{Target: &r.Start_ip}
	case 2:
		return &Ip_addressWrapper{Target: &r.End_ip}
	}
	panic("Unknown field index")
}

func (r *Event) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Event) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Event) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Event) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Event) Finalize()                        {}

func (_ *Event) AvroCRC64Fingerprint() []byte {
	return []byte(EventAvroCRC64Fingerprint)
}
