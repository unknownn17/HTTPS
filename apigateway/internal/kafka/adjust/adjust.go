package adjust

import (
	"api/internal/kafka/producer"
	"api/internal/models"
	"api/internal/protos/item"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

type Adjust struct {
	Client item.ItemServiceClient
	Ctx    context.Context
}

func (u *Adjust) Broadcast(key string, req interface{}) error {
	newreq, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := producer.Producer(key, newreq); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Adjust) Create() (*models.GeneralItem, error) {
	res, err := u.Client.ItemLastInserted(u.Ctx, &item.GetItemsRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralItem{ID: res.Id, Username: res.Username, Name: res.Name, Type: res.Type, Amount: res.Amount}, nil
}

func (u *Adjust) Get(id string) (*models.GeneralItem, error) {
	newid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := u.Client.ItemGet(u.Ctx, &item.GetItemRequest{Id: int32(newid)})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &models.GeneralItem{ID: res.Id, Username: res.Username, Name: res.Name, Type: res.Type, Amount: res.Amount}, nil
}

func (u *Adjust) Gets() ([]*models.GeneralItem, error) {
	res, err := u.Client.ItemsGet(u.Ctx, &item.GetItemsRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var newres []*models.GeneralItem

	for _, v := range res.Items {
		if v.Username==""{
			return nil,errors.New("there is no any items yet")
		}
		var all = models.GeneralItem{
			ID:       v.Id,
			Username: v.Username,
			Name:     v.Name,
			Type:     v.Type,
			Amount:   v.Amount,
		}
		newres = append(newres, &all)
	}
	return newres, nil
}
