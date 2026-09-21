package main

import (
	"fmt"
	"math"
)

func printSystem(matrix [][]float64, vector []float64) {
	fmt.Println("Исходная система:")
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%8.4f ", matrix[i][j])
		}
		fmt.Printf("| %8.4f\n", vector[i])
	}
	fmt.Println()
}

func determinant(matrix [][]float64) float64 {
	n := len(matrix)
	a := make([][]float64, n)
	for i := range a {
		a[i] = make([]float64, n)
		copy(a[i], matrix[i])
	}
	det := 1.0
	for k := 0; k < n; k++ {
		maxRow := k
		maxValue := math.Abs(a[k][k])
		for i := k + 1; i < n; i++ {
			if math.Abs(a[i][k]) > maxValue {
				maxValue = math.Abs(a[i][k])
				maxRow = i
			}
		}
		if maxRow != k {
			a[k], a[maxRow] = a[maxRow], a[k]
			det = -det
		}
		det *= a[k][k]
		for i := k + 1; i < n; i++ {
			factor := a[i][k] / a[k][k]
			for j := k; j < n; j++ {
				a[i][j] -= factor * a[k][j]
			}
		}
	}
	return det
}

func main() {
	matrix := [][]float64{
		{2.36, 2.37, 2.13},
		{2.51, 2.40, 2.10},
		{2.59, 2.41, 2.06},
	}
	vector := []float64{1.48, 1.92, 2.16}
	printSystem(matrix, vector)
	det := determinant(matrix)
	fmt.Println("Определитель матрицы: ", det)

	fmt.Println("Метод Гаусса:")
	solutionGauss, err := gaussMethod(matrix, vector)
	if err != nil {
		fmt.Println("Ошибка ", err)
		return
	}
	fmt.Println("Решение системы методом Гаусса:")
	for i, value := range solutionGauss {
		fmt.Printf("x%d = %.4f\n", i+1, value)
	}
	checkGauss(matrix, vector, solutionGauss)

	fmt.Println()

	fmt.Println("Метод Холецкого:")
	solutionKholets, err := kholetsMethod(matrix, vector)
	if err != nil {
		fmt.Println("Ошибка ", err)
		return
	}
	fmt.Println("Решение системы методом Холецкого:")
	for i, value := range solutionKholets {
		fmt.Printf("x%d = %.4f\n", i+1, value)
	}
	checkKholets(matrix, vector, solutionKholets)
}
