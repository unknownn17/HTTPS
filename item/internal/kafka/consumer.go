package kafkaconsumer

import (
	"context"
	"encoding/json"
	"item/internal/models"
	"item/internal/protos/item"
	servicemethods "item/internal/service/methods"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer17 struct {
	C   *servicemethods.GrpcService
	Ctx context.Context
}

func (u *Consumer17) Consumer() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("item17"),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	for {
		fetches := client.PollFetches(ctx)
		if err := fetches.Errors(); len(err) > 0 {
			log.Fatal(err)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				if err := u.Adjust(record); err != nil {
					log.Println(err)
				}
			}
		})
	}
}

func (u *Consumer17) Adjust(record *kgo.Record) error {
	switch string(record.Key) {
	case "create":
		if err := u.Create(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "update":
		if err := u.Update(record.Value); err != nil {
			log.Println(err)
			return err
		}
	case "delete":
		if err := u.Delete(record.Value); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (u *Consumer17) Create(req []byte) error {
	var req1 models.CreateItemRequest

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = item.CreateItemRequest{
		Username: req1.Username,
		Name:     req1.Name,
		Type:     req1.Type,
		Amount:   req1.Amount,
	}
	_, err := u.C.ItemCreate(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer17) Update(req []byte) error {
	var req1 models.GeneralItem

	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}
	var newreq = item.GeneralItem{
		Id:       req1.ID,
		Username: req1.Username,
		Name:     req1.Name,
		Type:     req1.Type,
		Amount:   req1.Amount,
	}
	_, err := u.C.ItemUpdate(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer17) Delete(req []byte) error {
	var req1 models.GetItemRequest
	if err := json.Unmarshal(req, &req1); err != nil {
		log.Println(err)
		return err
	}

	var newreq = item.GetItemRequest{
		Id: req1.ID,
	}

	_, err := u.C.ItemDelete(u.Ctx, &newreq)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
