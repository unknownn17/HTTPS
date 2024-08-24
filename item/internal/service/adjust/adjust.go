package adjust

import (
	"context"
	"item/internal/database/service"
	"item/internal/models"
	"item/internal/protos/item"
	"log"
)

type Adjust struct {
	S *service.Service
}

func (u *Adjust) Create(ctx context.Context, req *item.CreateItemRequest) (*item.DeleteResponse, error) {
	var newReq = models.CreateItemRequest{
		Username: req.Username,
		Name:     req.Name,
		Type:     req.Type,
		Amount:   req.Amount,
	}
	res, err := u.S.ItemCreate(ctx, &newReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &item.DeleteResponse{Message: res.Message}, nil
}

func (u *Adjust) Get(ctx context.Context, req *item.GetItemRequest) (*item.GeneralItem, error) {
	res, err := u.S.ItemGet(ctx, &models.GetItemRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &item.GeneralItem{
		Id:       res.ID,
		Username: res.Username,
		Name:     res.Name,
		Type:     res.Type,
		Amount:   res.Amount,
	}, nil
}

func (u *Adjust) Gets(ctx context.Context, req *item.GetItemsRequest) (*item.GetItemsResponse, error) {
	res, err := u.S.ItemGets(ctx, &models.GetItemsRequest{})
	if err != nil {
		log.Println(err)
	}
	var items []*item.GeneralItem

	for _, v := range res {
		var all = item.GeneralItem{
			Id:       v.ID,
			Username: v.Username,
			Name:     v.Name,
			Type:     v.Type,
			Amount:   v.Amount,
		}
		items = append(items, &all)
	}
	return &item.GetItemsResponse{Items: items}, nil
}

func (u *Adjust) Update(ctx context.Context, req *item.GeneralItem) (*item.DeleteResponse, error) {
	var newReq = models.GeneralItem{
		ID:       req.Id,
		Username: req.Username,
		Name:     req.Name,
		Type:     req.Type,
		Amount:   req.Amount,
	}
	res, err := u.S.ItemUpdate(ctx, &newReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &item.DeleteResponse{Message: res.Message}, err
}

func (u *Adjust) Delete(ctx context.Context, req *item.GetItemRequest) (*item.DeleteResponse, error) {
	res, err := u.S.ItemDelete(ctx, &models.GetItemRequest{ID: req.Id})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &item.DeleteResponse{Message: res.Message}, nil
}

func (u *Adjust) ItemLastInserted(ctx context.Context) (*item.GeneralItem, error) {
	res, err := u.S.ItemLastInserted(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &item.GeneralItem{
		Id:       res.ID,
		Username: res.Username,
		Name:     res.Name,
		Type:     res.Type,
		Amount:   res.Amount,
	}, nil
}
