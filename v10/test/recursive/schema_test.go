package avro

import (
	"io"
	"testing"

	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTripExactBytes(t,
		func() container.AvroRecord { return &RecursiveUnionTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeRecursiveUnionTestRecord(r)
			return &record, err
		})
}
