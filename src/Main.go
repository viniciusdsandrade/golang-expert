package main

import (
	"fmt"
	_ "fmt"
	"strconv"
)

func main() {

	fmt.Print("Digite um número: ")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Erro: entrada inválida. Por favor, digite um número.")
		return
	}

	fmt.Println("Quantidade de divisores de", num, ":", numDivisors(num))
	fmt.Println("Divisores de", num, ":", allDivisors(num))

	fmt.Print("Digite um número para verificar a quantidade de dígitos: ")
	var numerinho int
	_, err = fmt.Scan(&numerinho)
	if err != nil {
		fmt.Println("Erro: entrada inválida. Por favor, digite um número.")
		return
	}

	fmt.Println("O número", numerinho, "possui", quantityOfDigits(numerinho), "dígitos.")

	fmt.Println("Digite um numero para encontrar seu fatorial:")
	var numFat int64
	_, errror := fmt.Scan(&numFat)
	if errror != nil {
		fmt.Println("Erro: entrada inválida. Por favor, digite um número.")
		return
	}

	fmt.Println("Fatorial de", numFat, "é", fatorial1(numFat), "usando fatorial1")
	fmt.Println("Fatorial de", numFat, "é", fatorial2(numFat), "usando fatorial2")

	var isPalindromeNum int

	fmt.Println("Digite um número para verificar se é palíndromo:")
	_, err = fmt.Scan(&isPalindromeNum)
	if err != nil {
		fmt.Println("Erro: entrada inválida. Por favor, digite um número.")
		return
	}

	if isPalindromeNumber(isPalindromeNum) {
		fmt.Println("O número", isPalindromeNum, "é palíndromo.")
	} else {
		fmt.Println("O número", isPalindromeNum, "não é palíndromo.")
	}

	var mesStr string
	var mes int
	for {
		fmt.Println("Digite um número de 1 a 12:")
		_, err := fmt.Scan(&mesStr)
		if err != nil {
			fmt.Println("Erro: entrada inválida. Por favor, digite um número.")
			continue
		}

		mes, err = strconv.Atoi(mesStr) // Atribuição direta à variável dexterna 'mes'
		if err != nil {
			fmt.Println("Erro: entrada inválida. Por favor, digite um número.")
			continue
		}

		break
	}
	fmt.Println("O mês correspondente ao número", mes, "é", monthName(mes))

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
		fmt.Printf("\"%s\" é um palíndromo? %t\n", phrase, isPalindromeString(phrase))
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
		if isPrime(num) {
			primos = append(primos, num)
		}
	}

	fmt.Println("Números digitados:", numeros)
}
