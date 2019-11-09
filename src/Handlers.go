package main

import (
	"fmt"
	//"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	//"strconv"
)

func FormatError(w http.ResponseWriter, message string) {
	var cOutput CurrencyOutput
	cOutput.Successful = false
	cOutput.Message = message
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(500)

	if err := json.NewEncoder(w).Encode(cOutput); err != nil {
		http.Error(w, "Json não pôde ser parseado: "+err.Error(), 500)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Boas vindas!")
}

func ConversaoMoeda(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//brlVal := vars["brl_val"]
	currenciesValues, err := GetCurrencyValue("BRL")
	if err != nil {
		FormatError(w, err.Error())
		return
	}

	cOutput := GetTransferValue(currenciesValues)
	jsonProp, err := json.Marshal(cOutput)
	if err != nil {
		FormatError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	w.Write(jsonProp)
}
