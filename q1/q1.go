package q1

import (
	"errors"
)

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, errors.New("Valor da compra inválido")
	}

	totalCompras := somarCompras(purchaseHistory)
	mediaCompras := totalCompras / float64(len(purchaseHistory))

	var desconto float64

	if totalCompras > 1000 {
		desconto = currentPurchase * 0.1
	} else if totalCompras <= 1000 && totalCompras > 500 {
		desconto = currentPurchase * 0.05
	} else if totalCompras <= 500 {
		desconto = currentPurchase * 0.02
	}

	if len(purchaseHistory) == 0 {
		desconto = currentPurchase * 0.1
	}

	if mediaCompras > 1000 && desconto < currentPurchase*0.2 {
		desconto = currentPurchase * 0.2
	}

	return desconto, nil
}

func somarCompras(compras []float64) float64 {
	var total float64
	for _, valor := range compras {
		total += valor
	}
	return total
}
