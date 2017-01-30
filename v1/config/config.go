package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	DbServer string
	DbName   string
}

var AppConfiguration Configuration = Configuration{}

func ReadConfig() error {

	file, err := os.Open("../eventservice/conf.json")
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(&AppConfiguration)
}
