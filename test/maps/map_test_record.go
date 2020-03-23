// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type MapTestRecord struct {

	
	
		IntField *MapInt
	

	
	
		LongField *MapLong
	

	
	
		DoubleField *MapDouble
	

	
	
		StringField *MapString
	

	
	
		FloatField *MapFloat
	

	
	
		BoolField *MapBool
	

	
	
		BytesField *MapBytes
	

}

const MapTestRecordAvroCRC64Fingerprint = "3c3f18a007df5e9e"

func NewMapTestRecord() (*MapTestRecord) {
	return &MapTestRecord{}
}

func DeserializeMapTestRecord(r io.Reader) (*MapTestRecord, error) {
	t := NewMapTestRecord()
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

func DeserializeMapTestRecordFromSchema(r io.Reader, schema string) (*MapTestRecord, error) {
	t := NewMapTestRecord()

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

func writeMapTestRecord(r *MapTestRecord, w io.Writer) error {
	var err error
	
	err = writeMapInt( r.IntField, w)
	if err != nil {
		return err			
	}
	
	err = writeMapLong( r.LongField, w)
	if err != nil {
		return err			
	}
	
	err = writeMapDouble( r.DoubleField, w)
	if err != nil {
		return err			
	}
	
	err = writeMapString( r.StringField, w)
	if err != nil {
		return err			
	}
	
	err = writeMapFloat( r.FloatField, w)
	if err != nil {
		return err			
	}
	
	err = writeMapBool( r.BoolField, w)
	if err != nil {
		return err			
	}
	
	err = writeMapBytes( r.BytesField, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *MapTestRecord) Serialize(w io.Writer) error {
	return writeMapTestRecord(r, w)
}

func (r *MapTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":{\"type\":\"map\",\"values\":\"int\"}},{\"name\":\"LongField\",\"type\":{\"type\":\"map\",\"values\":\"long\"}},{\"name\":\"DoubleField\",\"type\":{\"type\":\"map\",\"values\":\"double\"}},{\"name\":\"StringField\",\"type\":{\"type\":\"map\",\"values\":\"string\"}},{\"name\":\"FloatField\",\"type\":{\"type\":\"map\",\"values\":\"float\"}},{\"name\":\"BoolField\",\"type\":{\"type\":\"map\",\"values\":\"boolean\"}},{\"name\":\"BytesField\",\"type\":{\"type\":\"map\",\"values\":\"bytes\"}}],\"name\":\"MapTestRecord\",\"type\":\"record\"}"
}

func (r *MapTestRecord) SchemaName() string {
	return "MapTestRecord"
}

func (_ *MapTestRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MapTestRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MapTestRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MapTestRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MapTestRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MapTestRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MapTestRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *MapTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MapTestRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.IntField = NewMapInt()
	
		
		
			return r.IntField
		
	
	case 1:
		
			r.LongField = NewMapLong()
	
		
		
			return r.LongField
		
	
	case 2:
		
			r.DoubleField = NewMapDouble()
	
		
		
			return r.DoubleField
		
	
	case 3:
		
			r.StringField = NewMapString()
	
		
		
			return r.StringField
		
	
	case 4:
		
			r.FloatField = NewMapFloat()
	
		
		
			return r.FloatField
		
	
	case 5:
		
			r.BoolField = NewMapBool()
	
		
		
			return r.BoolField
		
	
	case 6:
		
			r.BytesField = NewMapBytes()
	
		
		
			return r.BytesField
		
	
	}
	panic("Unknown field index")
}

func (r *MapTestRecord) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *MapTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MapTestRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *MapTestRecord) Finalize() { }


func (_ *MapTestRecord) AvroCRC64Fingerprint() string {
  return MapTestRecordAvroCRC64Fingerprint
}
