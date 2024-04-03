package main

import "math"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 5; i <= sqrtNum; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func isPalindromeNumber(num int) bool {
	if num < 0 {
		return false
	}
	return num == reverseNumber(num)
}

func reverseNumber(num int) int {
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num = num / 10
	}
	return reversed
}
