package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shon-phand/bookstore_users-api/domains/configuration"
)

func LoadConfiguration(filename string) (configuration.Config, error) {
	var config configuration.Config

	configFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("There is error in loading config file")
		return config, err
	}

	fmt.Println("File imported : ", configFile)

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}
