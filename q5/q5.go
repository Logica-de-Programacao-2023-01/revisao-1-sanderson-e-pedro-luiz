package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	if fromScale != "C" && fromScale != "F" && fromScale != "K" {
		return 0, fmt.Errorf("escala inválida")
	}
	if toScale != "C" && toScale != "F" && toScale != "K" {
		return 0, fmt.Errorf("escala inválida")
	}

	tempconv := 0.0

	if fromScale == "C" && toScale == "F" {
		tempconv = temp*9/5 + 32
	} else if fromScale == "C" && toScale == "K" {
		tempconv = temp + 273.15
	} else if fromScale == "F" && toScale == "C" {
		tempconv = (temp - 32) * 5 / 9
	} else if fromScale == "F" && toScale == "K" {
		tempconv = (temp-32)*5/9 + 273.15
	} else if fromScale == "K" && toScale == "C" {
		tempconv = temp - 273.15
	} else if fromScale == "K" && toScale == "F" {
		tempconv = (temp-273.15)*9/5 + 32
	}

	return tempconv, nil
}
