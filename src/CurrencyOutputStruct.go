package main

import (
	"time"
)

//CurrencyOutput type to define return structure
type CurrencyOutput struct {
	Date              time.Time         `json:"date"`
	FareDescripton    string            `json:"fareDescription"`
	CurrenciesOptions CurrenciesOptions `json:"currenciesOptions"`
}

//CurrenciesOptions stores all currencies supported
type CurrenciesOptions struct {
	BRL float64 `json:"BRL"`
	EUR float64 `json:"EUR"`
	USD float64 `json:"USD"`
}
