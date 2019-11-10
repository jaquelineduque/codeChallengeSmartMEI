package main

import (
	"testing"
)

func TestHasPrincipalCurrencies(t *testing.T) {
	expected := true
	if observed := HasPrincipalCurrencies("BRL"); observed != expected {
		t.Fatalf("HasPrincipalCurrencies() = %v, want %v", observed, expected)
	}
}


func TestGetBRLTransferFare(t *testing.T) {
	expected := 7.0
	observed, _ := GetBRLTransferFare("https://www.smartmei.com.br/#planos-e-tarifas")
	
	if observed != expected {
		t.Fatalf("GetBRLTransferFare() = %v, want %v", observed, expected)
	}
}

