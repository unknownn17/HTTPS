package servicemethods

import (
	"context"
	"item/internal/protos/item"
	"item/internal/service/adjust"
	"log"
)

type GrpcService struct {
	item.UnimplementedItemServiceServer
	A *adjust.Adjust
}

func (u *GrpcService) ItemCreate(ctx context.Context, req *item.CreateItemRequest) (*item.DeleteResponse, error) {
	res, err := u.A.Create(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *GrpcService) ItemGet(ctx context.Context, req *item.GetItemRequest) (*item.GeneralItem, error) {
	res, err := u.A.Get(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *GrpcService) ItemUpdate(ctx context.Context, req *item.GeneralItem) (*item.DeleteResponse, error) {
	res, err := u.A.Update(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *GrpcService) ItemsGet(ctx context.Context, req *item.GetItemsRequest) (*item.GetItemsResponse, error) {
	res, err := u.A.Gets(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *GrpcService) ItemDelete(ctx context.Context, req *item.GetItemRequest) (*item.DeleteResponse, error) {
	res, err := u.A.Delete(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (u *GrpcService) ItemLastInserted(ctx context.Context, req *item.GetItemsRequest) (*item.GeneralItem, error) {
	res, err := u.A.ItemLastInserted(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
