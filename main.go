package main

import (
	"fmt"

	"github.com/Ssnakerss/leecode/funcs"
)

func main() {
	// start := time.Now()
	fmt.Println(funcs.CountSpecialNumbers(10000) + 1)
	// // //597810
	// // // 553.5011ms
	// end := time.Since(start)
	// fmt.Println(end)
	// // num := make([]rune, 3)
	// start = time.Now()
	// fmt.Println(funcs.CountSpecialNumbers2(251_100_554))
	// //597810
	// //29.7091ms
	// end = time.Since(start)
	// fmt.Println(end)

	fmt.Println(countNumbersWithUniqueDigits(4))

	fmt.Println(countNumbersWithUniqueDigits2(4))
}

const (
	f9 = 362880
)

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	cnt := 0
	for i := 0; i <= n; i++ {
		d := enum(i)
		fmt.Println(i, " -> ", d)
		cnt += enum(i)
	}
	return cnt
}

func enum(n int) int {
	return f9/Factorial(9-n) + (n-1)*f9/Factorial(9-n+1)
}

// Factorial of n
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func countNumbersWithUniqueDigits2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}

	result := 10
	current := 9

	for i := 2; i <= n; i++ {
		current *= (10 - (i - 1))
		result += current
	}

	return result
}
