package main

import (
	"fmt"
	"math"
)

func givensMethod(matrix [][]float64, vector []float64) ([]float64, error) {
	n := len(matrix)
	new_matrix := make([][]float64, n)
	for i := range new_matrix {
		new_matrix[i] = make([]float64, n+1)
		copy(new_matrix[i], matrix[i])
		new_matrix[i][n] = vector[i]
	}
	//прямой ход
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if math.Abs(new_matrix[j][i]) < 1e-15 {
				continue
			}
			r := math.Sqrt(new_matrix[i][i]*new_matrix[i][i] + new_matrix[j][i]*new_matrix[j][i])
			if r < 1e-15 {
				continue
			}
			c := new_matrix[i][i] / r
			s := new_matrix[j][i] / r
			for k := 0; k <= n; k++ {
				previous_i := new_matrix[i][k]
				previous_j := new_matrix[j][k]
				new_matrix[i][k] = c*previous_i + s*previous_j
				new_matrix[j][k] = -s*previous_i + c*previous_j
			}
		}
		fmt.Printf("После обнуления столбца %d:\n", i+1)
		print(new_matrix)
	}
	//обратный ход
	solution := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += new_matrix[i][j] * solution[j]
		}
		solution[i] = (new_matrix[i][n] - sum) / new_matrix[i][i]
	}
	return solution, nil
}

func print(a [][]float64) {
	for i := range a {
		for j := range a[i] {
			fmt.Printf("%12.3f", a[i][j])
		}
		fmt.Println()
	}
}
