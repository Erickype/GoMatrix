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

	fmt.Println(len(matrix1) + len(matrix2))

	fmt.Printf("Time elapsed: %v", time.Since(start))
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
