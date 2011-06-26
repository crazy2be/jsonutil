package jsonutil

import (
	"fmt"
	"json"
	"os"
)

func DecodeFromFile(filename string, object interface{}) (err os.Error) {
	var file *os.File
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open %s, %s.\n", filename, err)
		return
	}
	enc := json.NewDecoder(file)
	err = enc.Decode(object)
	if err != nil {
		fmt.Printf("Failed to decode JSON object: %s from file %s. (Object end up as %#v).\n", err, filename, object)
		return
	}
	return
}

func EncodeToFile(filename string, object interface{}) (err os.Error) {
	var file *os.File
	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("Failed to open %s, %s.\n", filename, err)
		return
	}
	enc := json.NewEncoder(file)
	err = enc.Encode(object)
	if err != nil {
		fmt.Printf("Failed to encode JSON object: %s. (Object represented as %#v).\n", err, object)
		return
	}
	return
}