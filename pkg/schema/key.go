package schema

import "github.com/hamba/avro"

// Available message keys.
const (
	KeyTypeArticle = "article"
	KeyTypeVersion = "version"
)

// KeySchema avro schema for message key.
const KeySchema = `{
  "type": "record",
  "name": "Key",
  "namespace": "wme_poc.general.schema",
  "fields": [
    {
      "name": "identifier",
      "type": "string"
    },
    {
      "name": "type",
      "type": "string"
    }
  ]
}`

// NewKeySchema create new message key avro schema.
func NewKeySchema() (avro.Schema, error) {
	return avro.Parse(KeySchema)
}

// NewKey create new message key.
func NewKey(identifier string, keyType string) *Key {
	return &Key{
		Identifier: identifier,
		Type:       keyType,
	}
}

// Key schema for kafka message keys.
type Key struct {
	Identifier string `json:"identifier" avro:"identifier"`
	Type       string `json:"entity" avro:"type"`
}
