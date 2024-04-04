package serializer_test

import (
	"testing"

	"github.com/muhammedarifp/grpc-example/sample"
	"github.com/muhammedarifp/grpc-example/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	serializer.WriteProtobufToJSONFile(laptop1, "../tmp/laptop.json")
}
