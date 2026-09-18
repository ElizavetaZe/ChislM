package main

import "fmt"

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

func main() {
	matrix := [][]float64{
		{2.36, 2.37, 2.13},
		{2.51, 2.40, 2.10},
		{2.59, 2.41, 2.06},
	}
	vector := []float64{1.48, 1.92, 2.16}
	printSystem(matrix, vector)

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
