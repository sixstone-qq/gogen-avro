package avro

import (
	"io"
	"testing"

	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTripExactBytes(t, func() container.AvroRecord { return &RecursiveUnionTestRecord{} }, func(r io.Reader) (interface{}, error) {
		return DeserializeRecursiveUnionTestRecord(r)
	})
}
