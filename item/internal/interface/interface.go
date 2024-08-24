package interface17

import (
	"context"
	"item/internal/models"
)


type Item interface{
	ItemCreate(ctx context.Context,req *models.CreateItemRequest)(*models.DeleteResponse,error)
	ItemGet(ctx context.Context,req *models.GetItemRequest)(*models.GeneralItem,error)
	ItemGets(ctx context.Context,req *models.GetItemsRequest)([]*models.GeneralItem,error)
	ItemUpdate(ctx context.Context,req *models.GeneralItem)(*models.DeleteResponse,error)
	ItemDelete(ctx context.Context,req *models.GetItemRequest)(*models.DeleteResponse,error)
}