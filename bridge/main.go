package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	eventstream "github.com/protsack-stephan/mediawiki-eventstream-client"
	"github.com/protsack-stephan/schema-registry-example/pkg/schema"
)

const database = "enwiki"
const url = "https://en.wikipedia.org"
const bootstrapServers = "localhost:29092"

func main() {
	ctx := context.Background()
	since := time.Now()
	streams := eventstream.NewClient()
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})

	if err != nil {
		log.Panic(err)
	}

	defer producer.Close()

	artSch, err := schema.NewArticleSchema()

	if err != nil {
		log.Panic(err)
	}

	verSch, err := schema.NewVersionSchema()

	if err != nil {
		log.Panic(err)
	}

	keySch, err := schema.NewKeySchema()

	if err != nil {
		log.Panic(err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Println(ev.TopicPartition.Error)
				}
			}
		}
	}()

	stream := streams.RevisionCreate(ctx, since, func(evt *eventstream.RevisionCreate) error {
		if evt.Data.Database == database {
			art := new(schema.Article)
			art.Name = evt.Data.PageTitle
			art.Identifier = evt.Data.PageID
			art.DateModified = &evt.Data.RevTimestamp
			art.URL = fmt.Sprintf("%s/wiki/%s", url, evt.Data.PageTitle)
			art.Version = &schema.Version{
				Identifier: evt.Data.RevID,
			}

			ver := new(schema.Version)
			ver.Comment = evt.Data.Comment
			ver.Identifier = evt.Data.RevID

			artKey, err := schema.Marshal(1, keySch, schema.NewKey(fmt.Sprintf("/articles/%s/%s", evt.Data.Database, evt.Data.PageTitle), schema.KeyTypeArticle))

			if err != nil {
				return err
			}

			artVal, err := schema.Marshal(3, artSch, art)

			if err != nil {
				return err
			}

			artMsg := &kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &schema.TopicArticles, Partition: kafka.PartitionAny},
				Key:            artKey,
				Value:          artVal,
			}

			if err := producer.Produce(artMsg, nil); err != nil {
				return err
			}

			verKey, err := schema.Marshal(1, keySch, schema.NewKey(fmt.Sprintf("/versions/%s/%d", evt.Data.Database, evt.Data.RevID), schema.KeyTypeVersion))

			if err != nil {
				return err
			}

			verVal, err := schema.Marshal(2, verSch, ver)

			if err != nil {
				return err
			}

			verMsg := &kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &schema.TopicVersions, Partition: kafka.PartitionAny},
				Key:            verKey,
				Value:          verVal,
			}

			return producer.Produce(verMsg, nil)
		}

		return nil
	})

	for err := range stream.Sub() {
		if err != nil {
			log.Println(err)
		}
	}
}
