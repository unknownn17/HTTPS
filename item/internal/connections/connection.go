package connections

import (
	"database/sql"
	"item/internal/database/methods"
	"item/internal/database/service"
	interface17 "item/internal/interface"
	"log"
	_"github.com/lib/pq"
)


func NewDatabase()interface17.Item{
	db,err:=sql.Open("postgres","postgres://postgres:2005@localhost/test?sslmode=disable")
	if err!=nil{
		log.Println(err)
	}
	if err:=db.Ping();err!=nil{
		log.Println(err)
	}
	return &methods.Database{Db: db}
}

func NewService()*service.Service{
	a:=NewDatabase()
	return &service.Service{I: a}
}