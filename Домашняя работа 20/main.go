package main

import "fmt"

func main() {
	task_1()
}

func task_1() {

	var result [5][3]int
	var transpose [3][3]int
	var transposeResult [3][5]int
	
	matrix := [3][3]int{
		{1, 7, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	//Транспонирование матрицы
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			transpose[i][j] = matrix[j][i]
		}
	}
	//Добавление первых двух столбцов из исходной матрицы
	for i, j := 0, 0; i < len(result); i, j = i+1, j+1 {
		if j >= len(transpose) {
			j = 0
		}
		result[i] = transpose[j]
	}
	//Транспонирование матрицы
	for i := 0; i < len(result[0]); i++ {
		for j := 0; j < len(result); j++ {
			transposeResult[i][j] = result[j][i]
		}
	}
	var det int
	// Вычисление определителя
	for j := 0; j < len(transposeResult); j++ {
		x := 1
		y := 1
		for i := 0; i < len(transposeResult); i++ {
			x = x * transposeResult[i][i+j]
			y = y * transposeResult[i][len(transposeResult[i])-1-i-j]			
		}
		det += x - y
		
	}
	fmt.Println(det)
}

func task_2() {
	var result [3][4]int

	matrixOne := [3][5]int{
		{1, 2, 3, 4, 5},
		{7, 8, 9, 10, 11},
		{7, 8, 21, 34, 12},
	}
	matrixTwo := [5][4]int{
		{1, 2, 3, 23},
		{4, 5, 6, 12},
		{4, 5, 6, 7},
		{4, 5, 6, 12},
		{4, 5, 6, 11},
	}
	for i := 0; i < len(matrixOne); i++ {
		for j := 0; j < len(matrixTwo[0]); j++ {
			result[i][j] = 0
			for k := 0; k < 5; k++ {
				result[i][j] += matrixOne[i][k] * matrixTwo[k][j]
			}
		}
	}
	fmt.Println(result)

}
