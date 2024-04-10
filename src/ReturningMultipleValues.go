package main

func somaAndBool(a, b int) (int, bool) {
	if a > b {
		return a + b, true
	}

	return a - b, false
}
