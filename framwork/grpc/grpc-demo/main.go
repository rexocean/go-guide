package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"grpc-demo/service"
)

func main() {
	user := &service.User{
		Username: "zhangsan",
		Age:      12,
	}

	marshal, err := proto.Marshal(user)
	if err != nil {
		fmt.Println("marshal err:", err)
		panic(err)
	}
	fmt.Printf("serialized content, length: %d, content:%s", len(string(marshal)), string(marshal))

	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("proto serialized content: %s\n", newUser.String())

}
