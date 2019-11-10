package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HasPrincipalCurrencies(currencyBase string) bool {
	currencyValues, err := GetCurrencyValue(currencyBase)
	if err != nil {
		return false
	}
	if (currencyValues.Rates.USD != 0) && (currencyValues.Rates.EUR != 0) {
		return true
	} else {
		return false
	}

}

func GetCurrencyValue(currencyBase string) (CurrenciesValues, error) {

	var currenciesValues CurrenciesValues

	urlConsult, _, err := GetURLConsult()
	if err != nil {
		return currenciesValues, err
	}
	response, err := http.Get(urlConsult + "?base=" + currencyBase)
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

func GetBRLTransferFare(URL string) (float64, error) {
	response, err := http.Get(URL)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	//Transform into a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	pageContent := string(dataInBytes)

	//get Transfer value by substring
	TransferFareStartIndex := strings.Index(pageContent, "tarifas-2-2-2\">")
	if TransferFareStartIndex == -1 {
		err := errors.New("No tarifas-2-2-2 element found")
		return 0, err
	}
	TransferFareStartIndex += 15
	tam := len(pageContent)
	//final := (tam - TransferFareStartIndex)

	pageContent = (pageContent[TransferFareStartIndex:tam])

	// Find the index of the closing tag
	TransferFareEndIndex := strings.Index(pageContent, "</div>")
	if TransferFareEndIndex == -1 {
		err := errors.New("No closing tag for tarifas-2-2-2 found")
		return 0, err
	}
	pageTransferFare := (pageContent[0:TransferFareEndIndex])
	valueTransferFare := strings.Replace(pageTransferFare, "R$", "", -1)
	valueTransferFare = strings.Trim(valueTransferFare, "\n")
	valueTransferFare = strings.Trim(valueTransferFare, " ")
	valueTransferFare = strings.Replace(valueTransferFare, ",", ".", -1)

	f, err := strconv.ParseFloat(valueTransferFare, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func GetTransferValue(currenciesValues CurrenciesValues) CurrencyOutput {
	urlFares, _, _ := GetURLFares()

	var currencyOutput CurrencyOutput
	currencyOutput.Date = time.Now()
	currencyOutput.FareDescripton = "Descrição fixa"
	currencyOutput.CurrenciesOptions.BRL, _ = GetBRLTransferFare(urlFares)
	currencyOutput.CurrenciesOptions.EUR = BRLToEUR(currencyOutput.CurrenciesOptions.BRL, currenciesValues)
	currencyOutput.CurrenciesOptions.USD = BRLToUSD(currencyOutput.CurrenciesOptions.BRL, currenciesValues)
	currencyOutput.Successful = true
	return currencyOutput

}
