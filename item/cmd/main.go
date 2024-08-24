package main

import (
	"fmt"
	"item/internal/connections"
	"item/internal/protos/item"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	server := connections.NewGrpc()
	item.RegisterItemServiceServer(s, server)
	reflection.Register(s)
	a := connections.NewConsumer()
	go func() {
		a.Consumer()
	}()
	fmt.Printf("server started on the port 8080")

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
