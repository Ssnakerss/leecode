package funcs

import "fmt"

func CountAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	return rle(CountAndSay(n - 1))
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
