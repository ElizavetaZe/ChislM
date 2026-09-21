package main

import (
	"fmt"
	"math"
)

func zeidelMethod(matrix [][]float64, vector []float64, eps float64) ([]float64, int) {
	n := len(matrix)
	//начальное приближение
	x0 := make([]float64, n)
	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = vector[i] / matrix[i][i]
	}
	iteration := 0
	for {
		iteration++
		copy(x0, x)
		//новое приближение
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if i != j {
					sum += matrix[i][j] * x[j]
				}
			}
			x[i] = (vector[i] - sum) / matrix[i][i]
		}
		//проверка условия остановки
		maxDiff := 0.0
		for i := 0; i < n; i++ {
			diff := math.Abs(x[i] - x0[i])
			if diff > maxDiff {
				maxDiff = diff
			}
		}
		//вывод текущей итерации
		fmt.Printf("Итерация %d: ", iteration)
		for i := 0; i < n; i++ {
			fmt.Printf("x%d = %6.4f ", i+1, x[i])
		}
		fmt.Printf("Максимальная разница: %.4f\n", maxDiff)
		if maxDiff < eps {
			break
		}
		copy(x0, x)
	}
	return x, iteration
}
