package main

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

//прибавление к матрице С матрицу аЕ
func addAlfhaE(C [][]float64, alpha float64) [][]float64 {
	n := len(C)
	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
		copy(result[i], C[i])
		result[i][i] += alpha
	}
	return result
}

func regularationMethod(A [][]float64, b []float64, alpha float64, x0 []float64) ([]float64, float64, error) {
	n := len(A)
	C := symmetricMatrix(A)
	new_b := multiplyVector(A, b)
	if x0 != nil {
		for i := 0; i < n; i++ {
			new_b[i] += alpha * x0[i]
		}
	}
	C_reg := addAlfhaE(C, alpha)
	det := determinant(C_reg)
	solution, err := gaussMethod(C_reg, new_b)
	if err != nil {
		return nil, det, err
	}
	return solution, det, nil
}
