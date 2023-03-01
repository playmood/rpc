package pb_test

import (
	"fmt"
	"github.com/playmood/rpc/protobuf/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)

	str := &pb.String{Value: "hello"}

	// object --> protobuf --> []byte
	pbBytes, err := proto.Marshal(str)
	if should.NoError(err) {
		fmt.Println(pbBytes)
	}

	// []byte --> protobuf --> object
	obj := pb.String{}
	err = proto.Unmarshal(pbBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}
