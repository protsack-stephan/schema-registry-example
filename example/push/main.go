package main

import (
	"context"
	"log"

	"github.com/protsack-stephan/schema-registry-example/pkg/ksqldb"
)

const url = "http://localhost:8088/"
const query = `SELECT * from articles EMIT CHANGES;`

func main() {
	cl := ksqldb.NewClient(url)

	err := cl.Push(context.Background(), &ksqldb.QueryRequest{SQL: query}, func(qr *ksqldb.QueryResponse, row ksqldb.Row) {
		// log.Println(qr.ColumnTypes)
		// log.Println(qr.ColumnNames)
		log.Println(row.Int(3))
		log.Println(row.Time(3))
		log.Println("---------------------------------------------------------------------------------------")
	})

	if err != nil {
		log.Panic(err)
	}
}
