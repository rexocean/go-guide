package main

import (
	"google.golang.org/grpc"
	"grpc-demo/service"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	service.RegisterProductServiceServer(grpcServer, &service.ProductService{})

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("grpc server listening on :8888")
	_ = grpcServer.Serve(listen)
}
