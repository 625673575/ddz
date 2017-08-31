package message

import (
	"testing"
	"log"
	"github.com/golang/protobuf/proto"
)

func TestJsonMessage_Marshal(t *testing.T) {
	me := &LoginRequest{Username: "liuhongwei", Password: "4324ffdds"}
	b, _ := proto.Marshal(me)
	newme := new(LoginRequest)
	proto.Unmarshal(b, newme)
	log.Println(newme)
}
