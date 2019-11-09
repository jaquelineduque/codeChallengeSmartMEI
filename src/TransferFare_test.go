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
