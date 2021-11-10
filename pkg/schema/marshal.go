package schema

import (
	"encoding/binary"

	"github.com/hamba/avro"
)

func Marshal(id int, sch avro.Schema, v interface{}) ([]byte, error) {
	data, err := avro.Marshal(sch, v)

	if err != nil {
		return []byte{}, err
	}

	idb := make([]byte, 4)
	binary.BigEndian.PutUint32(idb, uint32(id))

	return append(append(append([]byte{}, byte(0)), idb...), data...), nil
}
