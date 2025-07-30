package funcs

import "math"

func divide(dividend int, divisor int) int {
	negative := false
	if dividend < 0 {
		negative = true
		dividend = -dividend
	}

	if divisor < 0 {
		if negative {
			negative = false
		} else {
			negative = true
		}
		divisor = -divisor
	}
	cnt := 0
	divisorX2 := divisor
	mul := 1
	mulMap := map[int]int{}
	divMap := map[int]int{}
	for dividend >= divisor {
		if dividend >= divisorX2 {
			dividend -= divisorX2
			cnt += mul
			mulMap[mul+mul] = mul
			mul += mul
			divMap[divisorX2+divisorX2] = divisorX2
			divisorX2 += divisorX2
		} else {
			mul = mulMap[mul]
			divisorX2 = divMap[divisorX2]
		}
	}
	if negative {
		return -cnt
	}
	if cnt > math.MaxInt32 {
		return math.MaxInt32
	}
	return cnt
}
