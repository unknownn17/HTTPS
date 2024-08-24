package service

import (
	"context"
	interface17 "item/internal/interface"
	"item/internal/models"
)

type Service struct {
	I interface17.Item
}

func (u *Service) ItemCreate(ctx context.Context, req *models.CreateItemRequest) (*models.DeleteResponse, error) {
	return u.I.ItemCreate(ctx, req)
}

func (u *Service) ItemGet(ctx context.Context, req *models.GetItemRequest) (*models.GeneralItem, error) {
	return u.I.ItemGet(ctx, req)
}

func (u *Service) ItemGets(ctx context.Context, req *models.GetItemsRequest) ([]*models.GeneralItem, error) {
	return u.I.ItemGets(ctx, req)
}

func (u *Service) ItemUpdate(ctx context.Context, req *models.GeneralItem) (*models.DeleteResponse, error) {
	return u.I.ItemUpdate(ctx, req)
}

func (u *Service) ItemDelete(ctx context.Context, req *models.GetItemRequest) (*models.DeleteResponse, error) {
	return u.I.ItemDelete(ctx, req)
}

func (u *Service) ItemLastInserted(ctx context.Context) (*models.GeneralItem, error) {
	return u.I.ItemLastInserted(ctx)
}
