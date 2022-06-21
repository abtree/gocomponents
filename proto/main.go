package main

import (
	_ "components/proto/pb"
	"reflect"

	"github.com/golang/protobuf/proto"
)

func main() {
	typ := proto.MessageType("pb.MsgCfgTest")
	reflect.New(typ)

}
