package main

import (
	"context"
	"log"

	"github.com/protsack-stephan/schema-registry-example/pkg/ksqldb"
)

const url = "http://localhost:8088/"
const query = `SELECT * from articles;`

func main() {
	cl := ksqldb.NewClient(url)
	ctx := context.Background()

	hr, rows, err := cl.Pull(ctx, &ksqldb.QueryRequest{SQL: query})

	if err != nil {
		log.Panic(err)
	}

	log.Println(hr)

	for _, row := range rows {
		log.Println(row.String(1))
	}
}
