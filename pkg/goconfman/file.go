package goconfman

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadFromFile(config interface{}, filePath string) error {
	configFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(configBytes, config)
	if err != nil {
		return err
	}

	return nil
}