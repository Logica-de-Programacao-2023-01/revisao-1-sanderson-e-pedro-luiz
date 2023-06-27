package q3

import (
	"errors"
	"sort"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}
	sort.Ints(numbers)
	maior := numbers[len(numbers)-1]
	menor := numbers[0]
	soma := 0
	for i := 1; i < len(numbers)-1; i++ {
		soma += numbers[i]
	}
	media := float64(soma) / float64(len(numbers)-2)

	return maior, menor, media, nil
}
