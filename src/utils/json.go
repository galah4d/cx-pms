package utils

import (
	"io/ioutil"
	"os"
)

func ReadJson(f string) ([]byte, error) {
	jsonFile, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}

func DumpJson(b []byte, f string) error {
	return ioutil.WriteFile(f, b, 0644)
}
