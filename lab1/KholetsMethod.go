package main

import (
	"fmt"
	"math"
)

// проверка на семметричность
func isSymmetric(matrix [][]float64, eps float64) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if math.Abs(matrix[i][j]-matrix[j][i]) > eps {
				return false
			}
		}
	}
	return true
}

// приведение к симметричной матрице
func symmetricMatrix(matrixA [][]float64) [][]float64 {
	n := len(matrixA)
	matrixC := make([][]float64, n)
	for i := range matrixC {
		matrixC[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < n; k++ {
				sum += matrixA[k][i] * matrixA[k][j]
			}
			matrixC[i][j] = sum
		}
	}
	return matrixC
}

// умножение транспонируемой матрицы на вектор
func multiplyVector(matrixA [][]float64, b []float64) []float64 {
	n := len(matrixA)
	new_b := make([]float64, n)
	for i := 0; i < n; i++ {
		sum := 0.0
		for k := 0; k < n; k++ {
			sum += matrixA[k][i] * b[k]
		}
		new_b[i] = sum
	}
	return new_b
}

// разложение Холецкого C=L*L^T
func kholetsDecompose(matrixC [][]float64) ([][]float64, error) {
	n := len(matrixC)
	matrixL := make([][]float64, n)
	for i := range matrixL {
		matrixL[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			sum := 0.0
			for k := 0; k < j; k++ {
				sum += matrixL[i][k] * matrixL[j][k]
			}
			if i == j {
				matrixL[i][j] = math.Sqrt(matrixC[i][i] - sum)
			} else {
				matrixL[i][j] = (matrixC[i][j] - sum) / matrixL[j][j]
			}
		}
	}
	return matrixL, nil
}

// прямой ход: решение L*y=b
func forward(matrixL [][]float64, b []float64) []float64 {
	n := len(matrixL)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < i; j++ {
			sum += matrixL[i][j] * y[j]
		}
		y[i] = (b[i] - sum) / matrixL[i][i]
	}
	return y
}

// обратный ход: решение L^(T)*x=y
func backward(matrixL [][]float64, y []float64) []float64 {
	n := len(matrixL)
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += matrixL[j][i] * x[j]
		}
		x[i] = (y[i] - sum) / matrixL[i][i]
	}
	return x
}

// метод Холецкого
func kholetsMethod(matrix [][]float64, vector []float64) ([]float64, error) {
	n := len(matrix)
	esp := 1e-9
	var C [][]float64
	var b []float64
	//1 шаг
	if isSymmetric(matrix, esp) {
		fmt.Println("Матрица симметрична")
		C, b = matrix, vector
	} else {
		fmt.Println("Матрица не симметрична. Приводим ее к симметрии умножением на транспонируемую (вектор тоже)")
		C = symmetricMatrix(matrix)
		b = multiplyVector(matrix, vector)
		fmt.Println("\nПреобразованная симметричная матрица:")
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf("%10.4f ", C[i][j])
			}
			fmt.Printf("| %10.4f\n", b[i])
		}
		fmt.Println()
	}
	//2 шаг
	L, err := kholetsDecompose(C)
	if err != nil {
		return nil, err
	}
	fmt.Println("Матрица L:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%10.4f ", L[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	//3 шаг
	y := forward(L, b)
	fmt.Println("Вектор y:")
	for i, val := range y {
		fmt.Printf("y%d = %10.4f\n", i+1, val)
	}
	fmt.Println()
	//4 шаг
	x := backward(L, y)
	return x, nil
}

// проверка подстановкой
func checkKholets(matrix [][]float64, vector []float64, solution []float64) {
	fmt.Println("\nПроверка решения методом Холецкого подстановкой:")
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
