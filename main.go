package main

import (
	"fmt"
)

func main() {
	// s := funcs.S
	// t := funcs.T
	// //  time: 5.3204844s
	// // s := "ADOBECODEBANC"
	// // t := "ABC"

	// start := time.Now()
	// minWin := funcs.MinWindow(s, t)
	// end := time.Since(start)
	// fmt.Printf("Min window: %s \n time: %v", minWin, end)

	fmt.Printf("%s \n", countAndSay(10))

}

func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	return rle(countAndSay(n - 1))
}

func rle(s string) string {
	res := ""
	cnt := 1

	i := 1
	for ; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			res += fmt.Sprintf("%d%s", cnt, string(s[i-1]))
			cnt = 1
		}
	}
	res += fmt.Sprintf("%d%s", cnt, string(s[i-1]))
	return res
}
