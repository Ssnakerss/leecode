package main

import (
	"fmt"
	"strconv"
	"sync"
)

type digit struct {
	Used bool
	Val  string
}

func main() {
	// start := time.Now()
	// fmt.Println(countSpecialNumbers(8_087_006))
	// //597810
	// // 553.5011ms
	// end := time.Since(start)
	// fmt.Println(end)

	// num := make([]rune, 3)

	digits := map[int]*digit{}
	for i := 0; i < 10; i++ {
		d := &digit{false, strconv.Itoa(i)}
		digits[i] = d
	}

	fmt.Println(10 + makeSpecialNumbers(3, "", digits, 135))

	// fmt.Println(digits)
	// var di, dj *digit

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i: ", i)
	// 	di = digits[i]
	// 	di.Used = true
	// 	for j := i; j < 10; j++ {
	// 		fmt.Println("  j: ", j)
	// 		dj = digits[j]
	// 		dj.Used = true
	// 		for _, d := range digits {
	// 			if !d.Used {
	// 				n, err := strconv.Atoi(di.Val + dj.Val + d.Val)
	// 				fmt.Println("    ", di.Val+dj.Val+d.Val, "  -  ", checkIfSpecial(n), err)
	// 			}
	// 		}
	// 	}
	// 	di.Used = false
	// 	dj.Used = false
	// }

}

func makeSpecialNumbers(size int, s string, digits map[int]*digit, target int) int {
	cnt := 0
	for _, d := range digits {
		if !d.Used {
			s1 := s + d.Val
			if size > 1 {
				d.Used = true
				cnt += makeSpecialNumbers(size-1, s1, digits, target)
				d.Used = false
			} else {
				n, _ := strconv.Atoi(s1)
				if n <= target {
					cnt++
					fmt.Println(s1, "  -  ", checkIfSpecial(n), cnt)
				}
			}
		}
	}
	return cnt
}

func countSpecialNumbers(n int) int {
	if n < 10 {
		return n
	}
	res := 9
	resC := make(chan bool)
	wg := &sync.WaitGroup{}
	go func() {
		for b := range resC {
			if b {
				res++
			}
		}
	}()

	for i := 10; i <= n; i++ {
		wg.Add(1)
		go func() {
			resC <- checkIfSpecial2(i)
			wg.Done()
		}()
	}
	wg.Wait()
	close(resC)
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

func checkIfSpecial2(n int) bool {
	m := map[rune]bool{}
	ns := strconv.Itoa(n)
	for _, r := range ns {
		if m[r] {
			return false
		} else {
			m[r] = true
		}
	}
	return true
}
