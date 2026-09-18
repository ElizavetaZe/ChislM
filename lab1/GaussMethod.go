package main

import (
	"fmt"
	"math"
)

func gaussMethod(matrix [][]float64, vector []float64) ([]float64, error) {
	n := len(matrix)
	//расширенная матрица
	new_matrix := make([][]float64, n)
	for i := range new_matrix {
		new_matrix[i] = make([]float64, n+1)
		copy(new_matrix[i], matrix[i])
		new_matrix[i][n] = vector[i]
	}
	//прямой ход Гаусса
	for k := 0; k < n; k++ {
		maxRow := k
		maxValue := math.Abs(new_matrix[k][k])
		for i := k + 1; i < n; i++ {
			if math.Abs(new_matrix[i][k]) > maxValue {
				maxValue = math.Abs(new_matrix[i][k])
				maxRow = i
			}
		}
		//перестановка строк
		if maxRow != k {
			new_matrix[k], new_matrix[maxRow] = new_matrix[maxRow], new_matrix[k]
		}
		//нормализация ведущей строки
		mainElement := new_matrix[k][k]
		for j := k; j <= n; j++ {
			new_matrix[k][j] /= mainElement
		}
		//исключение x_k
		for i := k + 1; i < n; i++ {
			value := new_matrix[i][k]
			for j := k; j <= n; j++ {
				new_matrix[i][j] -= value * new_matrix[k][j]
			}
		}
	}
	//обратный ход Гаусса
	solution := make([]float64, n)
	solution[n-1] = new_matrix[n-1][n]
	for i := n - 2; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += new_matrix[i][j] * solution[j]
		}
		solution[i] = new_matrix[i][n] - sum
	}
	return solution, nil
}

// проверка подстановкой
func checkGauss(matrix [][]float64, vector []float64, solution []float64) {
	fmt.Println("\nПроверка решения методом Гаусса подстановкой:")
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
