package main

import (
	"fmt"
	"log"

	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
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
	fmt.Fprintln(w, "Boas-vindas!")
}

func FilterFields(graphQuery string, currencyObj CurrencyOutput) []byte {
	var currenciesOptions = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "currenciesOptions",
			Fields: graphql.Fields{
				"BRL": &graphql.Field{
					Type: graphql.Float,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return currencyObj.CurrenciesOptions.BRL, nil
					},
				},
				"EUR": &graphql.Field{
					Type: graphql.Float,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return currencyObj.CurrenciesOptions.EUR, nil
					},
				},
				"USD": &graphql.Field{
					Type: graphql.Float,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return currencyObj.CurrenciesOptions.USD, nil
					},
				},
			},
		},
	)

	fields := graphql.Fields{
		"successful": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return currencyObj.Successful, nil
			},
		},
		"message": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return currencyObj.Message, nil
			},
		},
		"date": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return currencyObj.Date, nil
			},
		},
		"fareDescription": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return currencyObj.FareDescripton, nil
			},
		},
		"currenciesOptions": &graphql.Field{
			Type: currenciesOptions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return currencyObj.CurrenciesOptions, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	params := graphql.Params{Schema: schema, RequestString: graphQuery}
	rqryP := graphql.Do(params)
	if len(rqryP.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", rqryP.Errors)
	}
	rJSON, _ := json.Marshal(rqryP)
	return rJSON
}

func ConversaoMoeda(w http.ResponseWriter, r *http.Request) {
	currenciesValues, err := GetCurrencyValue("BRL")
	if err != nil {
		FormatError(w, err.Error())
		return
	}

	cOutput := GetTransferValue(currenciesValues)

	rJson := FilterFields(r.URL.Query().Get("query"), cOutput)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	w.Write(rJson)
}
