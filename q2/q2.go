package q2

import (
	"errors"
	"strings"
	"unicode"
)

func AverageLettersPerWord(text string) (float64, error) {
	palavras := strings.Fields(text)
	numPalavras := len(palavras)
	if numPalavras == 0 {
		return 0, errors.New("Texto vazio")
	}

	totalLetras := 0
	numPalavrasComLetras := 0
	for _, palavra := range palavras {
		numLetras := contarLetras(palavra)
		if numLetras > 0 {
			totalLetras += numLetras
			numPalavrasComLetras++
		}
	}

	if numPalavrasComLetras == 0 {
		return 0, errors.New("não há palavras com letras no texto")
	}

	media := float64(totalLetras) / float64(numPalavrasComLetras)
	return media, nil
}

func contarLetras(palavra string) int {
	numLetras := 0
	for _, char := range palavra {
		if unicode.IsLetter(char) {
			numLetras++
		}
	}
	return numLetras
}
