package main

import (
	"fmt"
	"math"
)

func printSystem(matrix [][]float64, vector []float64) {
	fmt.Println("Исходная система:")
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%10.4f ", matrix[i][j])
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

func check(matrix [][]float64, vector []float64, solution []float64) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += matrix[i][j] * solution[j]
		}
		diff := math.Abs(sum - vector[i])
		fmt.Printf("Уравнение %d: вычислено %10.4f, ожидается %10.4f, отклонение %.4e\n", i+1, sum, vector[i], diff)
	}
}

func main() {
	matrix := [][]float64{
		{6.9000, 0.0319, 0.0390, 0.0461},
		{0.0191, 6.0000, 0.0333, 0.0405},
		{0.0134, 0.0205, 5.1000, 0.0348},
		{0.0077, 0.0149, 0.0220, 4.2000},
	}
	vector := []float64{5.6632, 6.1119, 6.2000, 5.9275}
	eps := 0.001
	printSystem(matrix, vector)
	det := determinant(matrix)
	fmt.Println("Определитель матрицы: ", det)

	fmt.Println("Метод Якоби:")
	solutionYakobi, iterationY := yakobiMethod(matrix, vector, eps)
	fmt.Printf("\nРешение системы методом Якоби (e = %.4f):\n", eps)
	for i, value := range solutionYakobi {
		fmt.Printf("x%d = %.4f\n", i+1, value)
	}
	fmt.Printf("Количество итераций: %d\n", iterationY)
	fmt.Println("\nПроверка решения методом Якоби подстановкой:")
	check(matrix, vector, solutionYakobi)

	fmt.Println("\nМетод Зейделя:")
	solutionZeidel, iterationZ := zeidelMethod(matrix, vector, eps)
	fmt.Printf("\nРешение системы методом Зейделя (e = %.4f):\n", eps)
	for i, value := range solutionZeidel {
		fmt.Printf("x%d = %.4f\n", i+1, value)
	}
	fmt.Printf("Количество итераций: %d\n", iterationZ)
	fmt.Println("\nПроверка решения методом Зейделя подстановкой:")
	check(matrix, vector, solutionZeidel)
}
