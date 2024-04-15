package main

import (
	"math"
	"math/big"
)

func allDivisors(num int) []int {
	if num == 0 {
		return []int{}
	}

	var divisors []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i, -i)
		}
	}
	return divisors
}

func numDivisors(num int) int {
	if num == 0 {
		return 0
	}

	divisors := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisors += 2
		}
	}
	return divisors
}

func quantityOfDigits(num int) int {
	if num == 0 {
		return 1
	}
	if num < 0 {
		num = -num
	}

	digits := 0
	for num > 0 {
		num = num / 10
		digits++
	}

	return digits
}

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

func fatorial1(num int64) *big.Int {
	if num == 0 {
		return big.NewInt(1)
	}
	if num == 1 {
		return big.NewInt(1)
	}

	bigNum := big.NewInt(num)
	return bigNum.Mul(bigNum, fatorial1(num-1))
}

func fatorial2(num int64) *big.Int {
	if num == 0 {
		return big.NewInt(1)
	}
	if num == 1 {
		return big.NewInt(1)
	}
	fat := big.NewInt(1)
	for i := int64(2); i <= num; i++ {
		fat.Mul(fat, big.NewInt(i))
	}
	return fat
}
