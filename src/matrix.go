package main

import "fmt"

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			print(matrix[i][j], " ")
		}
		println()
	}
}

func printVector(vector []int) {
	fmt.Print("[")
	for i, v := range vector {
		fmt.Print(v)
		if i < len(vector)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func vectorSum(vector []int) int {
	sum := 0
	for _, v := range vector {
		sum += v
	}
	return sum
}

func matrixSum(matrix [][]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum += matrix[i][j]
		}
	}
	return sum
}
