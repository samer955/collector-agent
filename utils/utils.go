package utils

import (
	"encoding/json"
	"log"
)

func FromStructToByte(data any) ([]byte, error) {

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Unable to convert %s to bytes \n", data)
	}
	return bytes, err
}

func FromBytesToStruct(from []byte, to any) error {

	err := json.Unmarshal(from, to)
	if err != nil {
		log.Printf("Unable to Unmarshal Bytes to %s \n", to)
	}
	return err
}
