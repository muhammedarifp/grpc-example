package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJson(msg proto.Message) ([]byte, error) {
	marshler := protojson.MarshalOptions{
		Multiline:       true,
		EmitUnpopulated: true,
	}
	return marshler.Marshal(msg)
}
