package funcs

import (
	"unicode"
)

func BrockenKeysRev(keys string) string {
	capsB := newStack()
	smallB := newStack()
	curBigB := Key{Idx: -1}
	curSmallB := Key{Idx: -1}

	res := make([]rune, 0)
	runes := []rune(keys)
	for i := len(runes) - 1; i >= 0; i-- {
		switch keys[i] {
		case 'B':
			capsB.Push(Key{'B', i})
		case 'b':
			smallB.Push(Key{'b', i})
		default:
			if curBigB.Idx == -1 {
				curBigB, _ = capsB.TryPop()
			}
			if curSmallB.Idx == -1 {
				curSmallB, _ = smallB.TryPop()

			}
			if unicode.IsUpper(runes[i]) {
				if i > curBigB.Idx {

					res = append(res, runes[i])
				} else {
					curBigB, _ = capsB.TryPop()
				}
			} else {
				if i > curSmallB.Idx {

					res = append(res, runes[i])
				} else {
					curSmallB, _ = smallB.TryPop()
				}
			}
		}
	}
	return string(Reverse(res))
}

func Reverse(s []rune) []rune {
	if len(s) == 0 {
		return s
	} else {
		for i := 0; i < len(s)/2; i++ {
			if i != len(s)-i-1 {
				s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
			}
		}
	}
	return s
}
