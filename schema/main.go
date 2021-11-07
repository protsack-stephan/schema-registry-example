package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/protsack-stephan/schema-registry-example/pkg/schema"
)

const url = "http://localhost:8085"

type Reference struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Version int    `json:"version"`
}

type Subject struct {
	Schema     string       `json:"schema"`
	SchemaType string       `json:"schemaType"`
	References []*Reference `json:"references"`
}

func main() {
	subjects := map[string]*Subject{
		fmt.Sprintf("%s-key", schema.TopicVersions): {
			Schema:     schema.KeySchema,
			SchemaType: "AVRO",
		},
		fmt.Sprintf("%s-value", schema.TopicVersions): {
			Schema:     schema.VersionSchema,
			SchemaType: "AVRO",
		},
		fmt.Sprintf("%s-key", schema.TopicArticles): {
			Schema:     schema.KeySchema,
			SchemaType: "AVRO",
		},
		fmt.Sprintf("%s-value", schema.TopicArticles): {
			Schema:     schema.ArticleSchema,
			SchemaType: "AVRO",
			References: []*Reference{
				{
					Name:    "Version",
					Subject: fmt.Sprintf("%s-value", schema.TopicVersions),
					Version: 1,
				},
			},
		},
	}

	for name, subject := range subjects {
		body, err := json.Marshal(subject)

		if err != nil {
			log.Panic(err)
		}

		res, err := http.Post(fmt.Sprintf("%s/subjects/%s/versions", url, name), "application/json", bytes.NewReader(body))

		if err != nil {
			log.Panic(err)
		}

		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Panic(err)
		}

		log.Println(name)
		log.Println(string(data))
	}
}
