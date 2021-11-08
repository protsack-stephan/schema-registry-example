package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/protsack-stephan/schema-registry-example/pkg/schema"
)

const url = "http://localhost:8088/ksql"

func main() {
	queries := []string{
		fmt.Sprintf("CREATE STREAM articles WITH (KAFKA_TOPIC='%s', KEY_FORMAT='AVRO', VALUE_FORMAT='AVRO');", schema.TopicArticles),
		fmt.Sprintf("CREATE STREAM versions WITH (KAFKA_TOPIC='%s', KEY_FORMAT='AVRO', VALUE_FORMAT='AVRO');", schema.TopicVersions),
		"CREATE STREAM articles_versions AS SELECT * FROM articles INNER JOIN versions WITHIN 24 HOURS ON articles.version->identifier = versions.identifier;",
	}

	for _, q := range queries {
		res, err := http.Post(url, "application/json", strings.NewReader(fmt.Sprintf(`{"ksql": "%s"}`, q)))

		if err != nil {
			log.Panic(err)
		}

		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			log.Println(err)
		}

		if res.StatusCode != http.StatusOK {
			log.Println(string(data))
		}
	}
}
