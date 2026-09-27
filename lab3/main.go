package main

import (
	"fmt"
	"math"
)

func printSystem(matrix [][]float64, vector []float64) {
	fmt.Println("Исходная система:")
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%10.3f ", matrix[i][j])
		}
		fmt.Printf("| %8.3f\n", vector[i])
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
		{1.08, 0.996},
		{0.991, 0.944},
	}
	vector := []float64{0.502, 0.482}
	printSystem(matrix, vector)
	det := determinant(matrix)
	fmt.Println("Определитель матрицы: ", det)

	fmt.Println("Метод Регуляризации Тихонова:")
	x0 := []float64{0.0, 0.0}
	alphas := []float64{0.0, 0.001, 0.01, 0.1, 1.0}
	fmt.Printf("%-10s %-15s %-15s %-15s %-15s\n", "a", "x1", "x2", "det(C_reg)", "||r||")
	for _, alpha := range alphas {
		solution, detC_reg, err := regularationMethod(matrix, vector, alpha, x0)
		if err != nil {
			fmt.Printf("%-10.3f Ошибка: %v\n", alpha, err)
			continue
		}
		n := len(matrix)
		//ищем евклидову норму (длину) вектора
		rNorm := 0.0
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				sum += matrix[i][j] * solution[j]
			}
			rNorm += (sum - vector[i]) * (sum - vector[i])
		}
		rNorm = math.Sqrt(rNorm)

		fmt.Printf("%-10.3f %-15.3f %-15.3f %-15.3f %-15.3f\n", alpha, solution[0], solution[1], detC_reg, rNorm)
	}

	fmt.Println("\nМетод вращения Гивенса:")
	solutionGivens, err := givensMethod(matrix, vector)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Println("Решение системы методом Гивенса:")
	for i, value := range solutionGivens {
		fmt.Printf("x%d = %.3f\n", i+1, value)
	}
	check(matrix, vector, solutionGivens)
}
