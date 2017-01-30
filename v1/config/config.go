package config

import (
	"encoding/json"
	"os"
)

/*
Global Configuration struct
*/
type Configuration struct {
	DbServer string
	DbName   string
}

/*
Global Configuration instance
*/
var AppConfiguration Configuration = Configuration{}

/*
Read Global Configuration
*/
func ReadConfig() error {
	file, err := os.Open("../eventLogger/conf.json")
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(&AppConfiguration)
}
