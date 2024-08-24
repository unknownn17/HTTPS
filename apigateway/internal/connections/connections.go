package connections

import (
	"api/internal/api/handler"
	itemclient "api/internal/clients/item"
	"api/internal/kafka/adjust"
	"context"
)

func NewAdjust() *adjust.Adjust {
	a := itemclient.Itemclient()
	ctx := context.Background()
	return &adjust.Adjust{Client: a, Ctx: ctx}
}

func NewHandler()*handler.Handler{
	a:=NewAdjust()
	return &handler.Handler{A: a}
}