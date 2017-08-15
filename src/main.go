package main

import (
	"./protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	hw := &protobuf.Helloworld{
		Id:  1,
		Str: "2",
		Opt: 3,
	}
	data, err := proto.Marshal(hw)
	if err != nil {
		fmt.Println("marshaling error")
	}
	fmt.Println(hw)
	fmt.Println(data)
}
