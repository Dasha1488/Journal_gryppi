package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Чтение матрицы заданного размера
func inputMatrix(size int) [][]float64 {
	matrix := make([][]float64, size)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите элементы матрицы %dx%d построчно через пробел:\n", size, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Строка %d: ", i+1)
		scanner.Scan()
		line := scanner.Text()
		elements := strings.Fields(line)
		if len(elements) != size {
			fmt.Println("Некорректное количество элементов, попробуйте снова.")
			i--
			continue
		}
		row := make([]float64, size)
		for j, val := range elements {
			num, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Println("Некорректное число, попробуйте снова.")
				i--
				break
			}
			row[j] = num
		}
		matrix[i] = row
	}
	return matrix
}

// Вывод матрицы
func printMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%8.2f ", val)
		}
		fmt.Println()
	}
}

// Сложение двух матриц
func addMatrices(a, b [][]float64) [][]float64 {
	size := len(a)
	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

// Умножение матрицы на число
func multiplyMatrixByNumber(matrix [][]float64, number float64) [][]float64 {
	size := len(matrix)
	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = matrix[i][j] * number
		}
	}
	return result
}

// Умножение двух матриц
func multiplyMatrices(a, b [][]float64) [][]float64 {
	size := len(a)
	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			sum := 0.0
			for k := 0; k < size; k++ {
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Выберите размер матрицы (2 или 3):")
	scanner.Scan()
	sizeInput := scanner.Text()
	size, err := strconv.Atoi(sizeInput)
	if err != nil || (size != 2 && size != 3) {
		fmt.Println("Некорректный ввод. Завершение.")
		return
	}

	// Ввод первой матрицы
	fmt.Println("Введите первую матрицу:")
	matrixA := inputMatrix(size)

	// Ввод второй матрицы
	fmt.Println("Введите вторую матрицу:")
	matrixB := inputMatrix(size)

	for {
		fmt.Println("\nВыберите операцию:")
		fmt.Println("1 - Сложение матриц")
		fmt.Println("2 - Умножение матрицы на число")
		fmt.Println("3 - Умножение двух матриц")
		fmt.Println("4 - Выход")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			result := addMatrices(matrixA, matrixB)
			fmt.Println("Результат сложения:")
			printMatrix(result)
		case "2":
			fmt.Println("Введите число для умножения:")
			scanner.Scan()
			numStr := scanner.Text()
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Println("Некорректный ввод.")
				continue
			}
			result := multiplyMatrixByNumber(matrixA, num)
			fmt.Println("Результат умножения матрицы на число:")
			printMatrix(result)
		case "3":
			result := multiplyMatrices(matrixA, matrixB)
			fmt.Println("Результат умножения двух матриц:")
			printMatrix(result)
		case "4":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Некорректный выбор, попробуйте снова.")
		}
	}
}
