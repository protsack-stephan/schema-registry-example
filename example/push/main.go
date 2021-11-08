package main

import (
	"context"
	"log"

	"github.com/protsack-stephan/schema-registry-example/pkg/ksqldb"
)

const url = "http://localhost:8088/"
const query = `SELECT * from articles_versions EMIT CHANGES;`

func main() {
	cl := ksqldb.NewClient(url)

	err := cl.Push(context.Background(), &ksqldb.QueryRequest{SQL: query}, func(qr *ksqldb.QueryResponse, row ksqldb.Row) {
		log.Println(qr.ColumnNames)
		log.Println(row.String(2))
		log.Println("---------------------------------------------------------------------------------------")
	})

	if err != nil {
		log.Panic(err)
	}
}
