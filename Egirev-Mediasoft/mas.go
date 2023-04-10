package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Задаем размерность массива
	rows := 5
	cols := 10

	// Создаем двумерный массив со случайными уникальными числами
	rand.Seed(time.Now().UnixNano()) // Задаем seed для генерации случайных чисел
	arr := make([][]int, rows)
	used := make(map[int]bool)
	for i := 0; i < rows; i++ {
		arr[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			num := rand.Intn(rows * cols)
			for used[num] {
				num = rand.Intn(rows * cols)
			}
			arr[i][j] = num
			used[num] = true
		}
	}

	// Находим строку с самым большим числом
	maxSum := 0
	maxRow := 0
	for i := 0; i < rows; i++ {
		rowSum := 0
		for j := 0; j < cols; j++ {
			rowSum += arr[i][j]
		}
		if rowSum > maxSum {
			maxSum = rowSum
			maxRow = i
		}
	}

	// Выводим результаты
	fmt.Println("Массив:")
	for i := 0; i < rows; i++ {
		fmt.Println(arr[i])
	}
	fmt.Printf("Строка с самой большой суммой (%d): %v\n", maxSum, arr[maxRow])
}
