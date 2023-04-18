package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	var soma float64
	var contador int
	var resul float64
	for _, X := range purchaseHistory {
		soma += X
		contador++
	}
	med := soma / float64(contador)
	if med > 1000 {
		resul = currentPurchase * 0.2
		fmt.Println("O valor do produto final é de: ", resul)
	} else if soma > 1000 {
		resul = currentPurchase * 0.1
		fmt.Println("O valor do desconto do cliente é de 10%: ", resul)
	} else if soma <= 1000 {
		resul = currentPurchase * 0.05
		fmt.Println("O valor do desconto do cliente é de 5%: ", resul)
	} else if soma <= 500 {
		resul = currentPurchase * 0.02
		fmt.Println("O valor do desconto do cliente é de 2%: ", resul)
	} else if len(purchaseHistory) == 0 {
		resul = currentPurchase * 0.1
		fmt.Println("Essa é sua primeira compra, teve um desconto de 10%: ", resul)
	}
	return resul, nil
}
