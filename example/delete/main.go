package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const url = "http://localhost:8088/ksql"

func main() {
	queries := []string{
		"DROP STREAM articles_versions;",
		"DROP STREAM articles;",
		"DROP STREAM versions;",
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
