package schema

import "github.com/hamba/avro"

// VersionSchema avro schema for version.
const VersionSchema = `{
  "type": "record",
  "name": "Version",
  "namespace": "wme_poc.general.schema",
  "fields": [
    {
      "name": "identifier",
      "type": "int"
    },
    {
      "name": "parent_identifier",
      "type": "int"
    },
    {
      "name": "comment",
      "type": "string"
    }
  ]
}`

// NewVersionSchema creates new version avro schema.
func NewVersionSchema() (avro.Schema, error) {
	return avro.Parse(VersionSchema)
}

// Version representation for the article.
// Mainly modeled after https://schema.org/Thing.
type Version struct {
	Identifier       int    `json:"identifier,omitempty" avro:"identifier"`
	ParentIdentifier int    `json:"parent_identifier,omitempty" avro:"parent_identifier"`
	Comment          string `json:"comment,omitempty" avro:"comment"`
}
