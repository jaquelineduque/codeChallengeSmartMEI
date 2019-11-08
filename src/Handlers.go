package main

import (
	"fmt"
	//"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	//"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Boas vindas!")
}

func ConversaoMoeda(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//brlVal := vars["brl_val"]
	currenciesValues, _ := GetCurrencyValue("BRL")
	jsonProp, _ := json.Marshal(currenciesValues)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	w.Write(jsonProp)
}
