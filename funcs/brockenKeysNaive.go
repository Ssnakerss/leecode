package funcs

import (
	"unicode"
)

func BrockenKeysNaive(keys string) string {
	res := make([]rune, 0)
	lowerArr := make([]int, 0)
	upperArr := make([]int, 0)
	shift := 0

	// fmt.Println(keys)

	for pos, r := range keys {
		if r == 'b' {
			if len(lowerArr) == 0 {
				continue
			}
			if lowerArr[len(lowerArr)-1]+1 <= len(res) {
				res = append(res[:lowerArr[len(lowerArr)-1]], res[lowerArr[len(lowerArr)-1]+1:]...)
			} else {
				res = res[:lowerArr[len(lowerArr)-1]]
			}
			idx := len(upperArr) - 1
			for idx >= 0 && upperArr[idx] > lowerArr[len(lowerArr)-1] {
				upperArr[idx] -= 1
				idx -= 1
			}
			lowerArr = lowerArr[:len(lowerArr)-1]
			shift = shift - 2
		} else if r == 'B' {
			if len(upperArr) == 0 {
				continue
			}
			if upperArr[len(upperArr)-1]+1 <= len(res) {
				res = append(res[:upperArr[len(upperArr)-1]], res[upperArr[len(upperArr)-1]+1:]...)
			} else {
				res = res[:upperArr[len(upperArr)-1]]
			}
			idx := len(lowerArr) - 1
			for idx >= 0 && lowerArr[idx] > upperArr[len(upperArr)-1] {
				lowerArr[idx] -= 1
				idx -= 1
			}
			upperArr = upperArr[:len(upperArr)-1]
			shift = shift - 2
		} else {
			if unicode.IsUpper(r) {
				upperArr = append(upperArr, pos+shift)
			} else {
				lowerArr = append(lowerArr, pos+shift)
			}
			res = append(res, r)
		}
	}
	return string(res)
}
