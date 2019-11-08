package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetCurrencyValue(currencyBase string) (CurrenciesValues, error) {

	var currenciesValues CurrenciesValues
	response, err := http.Get("https://api.exchangeratesapi.io/latest?base=" + currencyBase)
	if err != nil {
		return currenciesValues, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return currenciesValues, err
	}
	if err := json.Unmarshal(responseData, &currenciesValues); err != nil {
		return currenciesValues, err
	}
	return currenciesValues, nil

}
