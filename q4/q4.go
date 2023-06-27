package q4

import (
	"fmt"
)

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	if taxCode <= 1 && taxCode > 3 {
		return 0, fmt.Errorf("imposto não encontrado")
	}

	if basePrice <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}

	aliquota := 0.0

	if taxCode == 1 {
		aliquota = 0.05
	} else if taxCode == 2 {
		aliquota = 0.1
	} else {
		aliquota = 0.15
	}

	frete := 0.0

	if state == "SP" {
		frete = 0.1
	} else if state == "RJ" {
		frete = 0.15
	} else if state == "MG" {
		frete = 0.2
	} else if state == "ES" {
		frete = 0.25
	} else {
		frete = 0.3
	}

	finalPrice := basePrice + basePrice*aliquota + basePrice*frete

	return finalPrice, nil
}
