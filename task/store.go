package task

import (
	"encoding/json"
	"os"
)

func LoadFileContent[T any](filename string) (T, error) {
	var data T

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	errDecode := decoder.Decode(&data)

	return data, errDecode
}

func WriteFile[T any](content T, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		encoder := json.NewEncoder(file)
		err = encoder.Encode(content)

	}
	defer file.Close()

	return err
}
