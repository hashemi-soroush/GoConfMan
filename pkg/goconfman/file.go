package goconfman

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadFromFile(config interface{}, filePath string) error {

	fileExt := filepath.Ext(filePath)
	configFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	if fileExt == ".json" {
		defer configFile.Close()
		configBytes, err := ioutil.ReadAll(configFile)
		if err != nil {
			return err
		}

		err = json.Unmarshal(configBytes, config)
		if err != nil {
			return err
		}
	} else if fileExt == ".yml" || fileExt == ".yaml" {
		yamlFile, err := ioutil.ReadAll(configFile)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(yamlFile, config); err != nil {
			return err
		}
	}
	return nil
}
