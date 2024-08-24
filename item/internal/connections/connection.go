package connections

import (
	"context"
	"database/sql"
	"item/internal/database/methods"
	"item/internal/database/service"
	interface17 "item/internal/interface"
	kafkaconsumer "item/internal/kafka"
	"item/internal/service/adjust"
	servicemethods "item/internal/service/methods"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() interface17.Item {
	db, err := sql.Open("postgres", "postgres://postgres:2005@localhost/test?sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	return &methods.Database{Db: db}
}

func NewService() *service.Service {
	a := NewDatabase()
	return &service.Service{I: a}
}

func NewAdjust() *adjust.Adjust {
	a := NewService()
	return &adjust.Adjust{S: a}
}

func NewGrpc() *servicemethods.GrpcService {
	a := NewAdjust()
	return &servicemethods.GrpcService{A: a}
}

func NewConsumer()*kafkaconsumer.Consumer17{
	a:=NewGrpc()
	ctx:=context.Background()
	return &kafkaconsumer.Consumer17{C: a,Ctx: ctx}
}