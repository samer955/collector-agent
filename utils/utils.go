package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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

func LatencyCalc(actual, before time.Time) int64 {

	if before.After(actual) {
		log.Println("Negative Latency. Something went wrong. 0 default value is set.")
		return 0
	}

	return actual.UnixMilli() - before.UnixMilli()

}

func HandlePanicError() {

	if r := recover(); r != nil {
		fmt.Println("Recovered. Error:\n", r)
	}
}
