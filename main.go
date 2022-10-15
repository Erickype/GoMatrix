package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	const n = 3

	matrix1 := populateRandomValues(n)
	matrix2 := populateRandomValues(n)

	result := multiplyMatrices(matrix1, matrix2)

	fmt.Printf("Length: %v\n", len(matrix1))

	fmt.Printf("M1: %v\n", matrix1)
	fmt.Printf("M2: %v\n", matrix2)

	fmt.Printf("Result: \n %v \n", result)

	fmt.Printf("Time elapsed: %v", time.Since(start))
}

func multiplyMatrices(matrix1 [][]int, matrix2 [][]int) [][]int {

	result := make([][]int, len(matrix1))

	for k := 0; k < len(matrix1); k++ {
		for f := 0; f < len(matrix1); f++ {
			sum := 0

			for i := 0; i < len(matrix1); i++ {
				for j := 0; j < len(matrix2); j++ {
					sum = sum + matrix1[k][i]*matrix2[j][f]
				}
			}

			result[k] = append(result[k], sum)
		}
	}

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
