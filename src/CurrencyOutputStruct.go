package main

import (
	"time"
)

//CurrencyOutput type to define return structure
type CurrencyOutput struct {
	Successful        bool              `json:"successful"`
	Message           string            `json:"message,omitempty"`
	Date              time.Time         `json:"date,omitempty"`
	FareDescripton    string            `json:"fareDescription,omitempty"`
	CurrenciesOptions CurrenciesOptions `json:"currenciesOptions,omitempty"`
}

//CurrenciesOptions stores all currencies supported
type CurrenciesOptions struct {
	BRL float64 `json:"BRL,omitempty"`
	EUR float64 `json:"EUR,omitempty"`
	USD float64 `json:"USD,omitempty"`
}
