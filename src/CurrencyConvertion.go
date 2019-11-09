package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
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

func BRLToEUR(valBRL float64, currenciesValues CurrenciesValues) float64 {
	return (currenciesValues.Rates.EUR * valBRL)
}

func BRLToUSD(valBRL float64, currenciesValues CurrenciesValues) float64 {
	return (currenciesValues.Rates.USD * valBRL)
}

func GetBRLTransferFare() float64 {
	return 7
}

func GetTransferValue(currenciesValues CurrenciesValues) CurrencyOutput {
	var currencyOutput CurrencyOutput
	currencyOutput.Date = time.Now()
	currencyOutput.FareDescripton = "Descrição fixa"
	currencyOutput.CurrenciesOptions.BRL = GetBRLTransferFare()
	currencyOutput.CurrenciesOptions.EUR = BRLToEUR(currencyOutput.CurrenciesOptions.BRL, currenciesValues)
	currencyOutput.CurrenciesOptions.USD = BRLToUSD(currencyOutput.CurrenciesOptions.BRL, currenciesValues)
	return currencyOutput

}
