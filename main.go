package main

import (
	"fmt"
	"github.com/revisao-1/bonus"
	"github.com/revisao-1/q1"
	"github.com/revisao-1/q2"
	"github.com/revisao-1/q3"
	"github.com/revisao-1/q4"
	"github.com/revisao-1/q5"
)

func main() {
	discount, err := q1.CalculateDiscount(5000, []float64{1000, 2050, 3500, 4000, 5500})
	fmt.Printf("Q1:\tdiscount: %f,\terr: %v\n", discount, err)

	averageLettersPerWord, err := q2.AverageLettersPerWord("Olá, meu nome é Gopher!")
	fmt.Printf("Q2:\taverageLettersPerWord: %f,\terr: %v\n", averageLettersPerWord, err)

	min, max, avg, err := q3.FindMinMaxAverage([]int{1, 2, 3, 4, 5})
	fmt.Printf("Q3:\tmin: %d,\tmax: %d,\tavg: %f,\terr: %v\n", min, max, avg, err)

	price, err := q4.CalculateFinalPrice(100, "SP", 1)
	fmt.Printf("Q4:\tprice: %f,\terr: %v\n", price, err)

	temperature, err := q5.ConvertTemperature(100, "C", "F")
	fmt.Printf("Q5:\ttemperature: %f,\terr: %v\n", temperature, err)

	damage, err := bonus.CalculateDamage(100, 100)
	fmt.Printf("Bonus:\tdamage: %d,\terr: %v\n", damage, err)
}
