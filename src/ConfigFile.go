package main

import (
	"encoding/json"
	"os"
)

func FileToConfig(filename string) (Config, int, error) {

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return Config{}, 11012, err
	}

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return Config{}, 11013, err
	}
	return configuration, 0, nil
}

func GetURLConsult() (string, int, error) {
	configuration, idError, err := FileToConfig(PathConfigJson)
	if err != nil {
		return "", idError, err
	}

	return configuration.URLCurrencyValue, 0, nil

}

func GetURLFares() (string, int, error) {
	configuration, idError, err := FileToConfig(PathConfigJson)
	if err != nil {
		return "", idError, err
	}

	return configuration.URLFares, 0, nil

}
