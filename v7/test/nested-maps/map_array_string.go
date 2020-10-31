// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

func writeMapArrayString(r map[string][]string, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = writeArrayString(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapArrayStringWrapper struct {
	Target *map[string][]string
	keys   []string
	values [][]string
}

func (_ *MapArrayStringWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapArrayStringWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapArrayStringWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapArrayStringWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapArrayStringWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v []string
	v = make([]string, 0)

	r.values = append(r.values, v)
	return &ArrayStringWrapper{Target: &r.values[len(r.values)-1]}
}

func (_ *MapArrayStringWrapper) AppendArray() types.Field { panic("Unsupported operation") }
