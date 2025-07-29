package funcs

import (
	"fmt"
	"math"
)

func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	tMap := make([]int, 128)
	cnt := len(t)
	start, end := 0, 0
	minLen, startIndex := math.MaxInt, -1
	for _, char := range t {
		tMap[char]++
	}
	for end < len(s) {
		if tMap[s[end]] > 0 {
			cnt--
		}
		tMap[s[end]]--
		end++
		for cnt == 0 {
			if end-start < minLen {
				startIndex = start
				minLen = end - start
			}
			if tMap[s[start]] == 0 {
				cnt++
			}
			tMap[s[start]]++
			start++
		}
	}

	if startIndex == -1 {
		return ""
	}

	fmt.Printf("%v \n", tMap)

	return s[startIndex : startIndex+minLen]
}
