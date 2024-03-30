package main

import (
	"fmt"
	_ "fmt"
	"math"
	"time"
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

	fmt.Println("Digite seu 10 numeros quaisquer")
	var numeros []int

	for i := 0; i < 10; i++ {
		var num int
		fmt.Printf("Digite o %dº número: ", i+1)
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		numeros = append(numeros, num)
	}

	var primos []int

	for _, num := range numeros {
		if ehPrimo(num) {
			primos = append(primos, num)
		}
	}

	// Imprime os números primos encontrados
	fmt.Println("Números primos digitados:")
	fmt.Println(primos)

	fmt.Println("Vinícius dos Santos Andrade")

	// Declaração de variáveis
	var nome = "Vinícius"
	var idade = 25
	var versao float32 = 1.1

	fmt.Println("Olá, sr.", nome, "sua idade é", idade, "e a versão do programa é", versao)
	fmt.Println("O tipo da variável nome é", fmt.Sprintf("%T", nome))

	// Declaração listas

	var lista = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lista de números", lista)

	// Declaração de listas sem tamanho
	var lista2 = []int{1, 2, 3, 4, 5}
	fmt.Println("Lista de números", lista2)

	// Declaração de mapas
	mapa := map[string]string{
		"nome":       "Vinícius",
		"sobrenome":  "Andrade",
		"nascimento": "2001-12-06",
	}

	fmt.Println(mapa)

	nascimentoStr := mapa["nascimento"]

	// Converte a string de nascimento para um objeto time. Time
	nascimento, err := time.Parse("06/12/2001", nascimentoStr)
	if err != nil {
		fmt.Println("Erro ao converter data:", err)
		return
	}

	// Formata a data no formato dd/mm/aaaa
	nascimentoFormatado := nascimento.Format("06/12/2001")

	fmt.Println("Data de nascimento:", nascimentoFormatado)
}
