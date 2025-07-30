package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(countSpecialNumbers(8087006))
	end := time.Since(start)
	fmt.Println(end)
}

func countSpecialNumbers(n int) int {
	if n < 10 {
		return n
	}
	res := 9
	for i := 10; i <= n; i++ {
		if checkIfSpecial(i) {
			res++
		}
	}
	return res
}

func checkIfSpecial(n int) bool {
	if n < 10 {
		return true
	}
	m := map[int]bool{}

	for n > 0 {
		if m[n%10] {
			return false
		}
		m[n%10] = true
		n = n / 10
	}
	return true
}
