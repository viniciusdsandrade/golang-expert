package main

import (
	"fmt"
	_ "fmt"
	"math"
)

func ehPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	phrases := []string{
		"arara",
		"Osso",
		"Subi no Ônibus",
		"Anotaram a data da maratona",
		"O galo ama o lago",
		"Ovo",
		"Essa frase não é um palíndromo",
	}

	for _, phrase := range phrases {
		fmt.Printf("\"%s\" é um palíndromo? %t\n", phrase, isPalindrome(phrase))
	}

	var vector = []int{1, 2, 3, 4, 5}
	var sumVector = vectorSum(vector)
	printVector(vector)

	fmt.Println("Soma do vetor:", sumVector)

	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	printMatrix(matrix)
	var sumMatrix = matrixSum(matrix)
	fmt.Println("Soma da matriz:", sumMatrix)

	var myName = "Vinícius dos Santos Andrade"
	var myNameReversed1 = reverseString1(myName)
	var myNameReversed2 = reverseString2(myName)
	var myNameReversed3 = reverseString3(myName)

	fmt.Println(myNameReversed1)
	fmt.Println(myNameReversed2)
	fmt.Println(myNameReversed3)

	fmt.Println("Digite seu 10 numeros quaisquer")
	var numeros []int
	var primos []int

	for i := 0; i < 10; i++ {
		var num int
		fmt.Printf("Digite o %dº número: ", i+1)
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		numeros = append(numeros, num)
	}

	for _, num := range numeros {
		if ehPrimo(num) {
			primos = append(primos, num)
		}
	}

	fmt.Println("Números digitados:", numeros)
}
