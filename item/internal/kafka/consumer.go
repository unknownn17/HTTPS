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
		// kgo.ConsumerGroup("franz-group"),
		// kgo.ConsumeResetOffset(kgo.NewOffset().At(13)),
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
				Adjust(record)
			}
		})
	}
}

func (u *Consumer17) Adjust(record *kgo.Record) error {

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

func (u *Consumer17) Update(req []byte)error{
	
}
