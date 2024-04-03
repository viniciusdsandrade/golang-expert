package main

import (
	"strings"
)

func monthName(month int) string {

	if month < 1 || month > 12 {
		return "Mês inválido"
	}

	months := []string{
		"Janeiro", "Fevereiro", "Março", "Abril",
		"Maio", "Junho", "Julho", "Agosto",
		"Setembro", "Outubro", "Novembro", "Dezembro",
	}
	return months[month-1]
}

func removeAccents(s string) string {
	accentMap := map[rune]rune{
		'á': 'a', 'à': 'a', 'ã': 'a', 'â': 'a',
		'é': 'e', 'è': 'e', 'ê': 'e',
		'í': 'i', 'ì': 'i', 'î': 'i',
		'ó': 'o', 'ò': 'o', 'õ': 'o', 'ô': 'o',
		'ú': 'u', 'ù': 'u', 'û': 'u',
		'ç': 'c',
		'Á': 'A', 'À': 'A', 'Ã': 'A', 'Â': 'A',
		'É': 'E', 'È': 'E', 'Ê': 'E',
		'Í': 'I', 'Ì': 'I', 'Î': 'I',
		'Ó': 'O', 'Ò': 'O', 'Õ': 'O', 'Ô': 'O',
		'Ú': 'U', 'Ù': 'U', 'Û': 'U',
		'Ç': 'C',
	}

	runes := []rune(s)
	for i, r := range runes {
		if val, ok := accentMap[r]; ok {
			runes[i] = val
		}
	}
	return string(runes)
}

func isPalindromeString(s string) bool {

	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	s = removeAccents(s)

	// Verificar se a string original é igual à string revertida
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func reverseString1(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseString2(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}

func reverseString3(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
