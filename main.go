package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	const n = 100

	matrix1 := populateRandomValues(n)
	matrix2 := populateRandomValues(n)

	result := multiplyMatrices(matrix1, matrix2)

	fmt.Printf("Length: %v\n", len(matrix1))

	fmt.Printf("Result: \n %v \n", result)

	fmt.Printf("Time elapsed: %v", time.Since(start))
}

func multiplyMatrices(matrix1 [][]int, matrix2 [][]int) [][]int {

	result := make([][]int, len(matrix1))

	return result
}

func populateRandomValues(size int) [][]int {

	m := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m[i] = append(m[i], rand.Intn(10)-rand.Intn(9))
		}
	}
	return m
}
