package schema

import (
	"time"

	"github.com/hamba/avro"
)

// ArticleSchema avro schema for the article.
const ArticleSchema = `{
  "type": "record",
  "name": "Article",
  "namespace": "wme_poc.general.schema",
  "fields": [
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "identifier",
      "type": "int"
    },
    {
      "name": "date_modified",
      "type": [
        "null",
        {
          "type": "long",
          "logicalType": "timestamp-micros"
        }
      ]
    },
    {
      "name": "version",
      "type": [
        "null",
        "Version"
      ]
    },
    {
      "name": "url",
      "type": "string"
    }
  ]
}`

// NewArticleSchema create new article avro schema.
func NewArticleSchema() (avro.Schema, error) {
	for _, schema := range []string{
		VersionSchema,
	} {
		if _, err := avro.Parse(schema); err != nil {
			return nil, err
		}
	}

	return avro.Parse(ArticleSchema)
}

// Article schema for wikipedia article.
// Tries to compliant with https://schema.org/Article.
type Article struct {
	Name         string     `json:"name" avro:"name"`
	Identifier   int        `json:"identifier,omitempty" avro:"identifier"`
	DateModified *time.Time `json:"date_modified,omitempty" avro:"date_modified"`
	Version      *Version   `json:"version,omitempty" avro:"version"`
	URL          string     `json:"url,omitempty" avro:"url"`
}
