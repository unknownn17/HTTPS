package itemclient

import (
	"api/internal/protos/item"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Itemclient() item.ItemServiceClient {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := item.NewItemServiceClient(conn)
	return client
}
