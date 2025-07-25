package funcs

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unicode"
)

const (
	limit      = 100000
	arrSize    = 50
	lowerChars = "abcdefghijklmnopqrstuvwxyz"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetRandString(n int) string {
	res := make([]rune, 0)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		r := rnd.Intn(26)
		b := rnd.Intn(10)
		c := rnd.Intn(2)
		switch b {
		case 5:
			if c == 1 {
				res = append(res, rune('b'))
			} else {
				res = append(res, rune('B'))
			}
		default:
			if c == 1 {
				res = append(res, rune(upperChars[r]))
			} else {
				res = append(res, rune(lowerChars[r]))
			}
		}

	}
	return string(res)
}

func GenRandArray(size int, limit int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// size := rnd.Intn(n)
	// size := n
	if size == 0 {
		return []int{}
	} else {
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = rnd.Intn(limit)
		}
		return arr
	}
}
func ArrIns(arr []int, pos int, i int) []int {
	arr = append(arr[:pos], append([]int{i}, arr[pos:]...)...)
	return arr
}

func Merge(arr1 []int, arr2 []int) []int {
	res := make([]int, 0)
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			res = append(res, arr1[0])
			arr1 = arr1[1:]
		} else {
			res = append(res, arr2[0])
			arr2 = arr2[1:]
		}
	}
	res = append(res, arr1...)
	res = append(res, arr2...)
	return res
}

func appendDistinct(arr []int, i int) []int {
	if len(arr) == 0 {
		return []int{i}
	} else {
		if arr[len(arr)-1] != i {
			return append(arr, i)
		} else {
			return arr
		}
	}
}
func MergeSortedLists(r1 []int, r2 []int) []int {
	m := make([]int, 0)
	for len(r1) > 0 || len(r2) > 0 {
		if len(r1) > 0 && len(r2) > 0 {
			if r1[0] < r2[0] {
				m = appendDistinct(m, r1[0])
				r1 = r1[1:]
			} else {
				m = appendDistinct(m, r2[0])
				r2 = r2[1:]
			}
		} else if len(r1) > 0 {
			m = appendDistinct(m, r1[0])
			r1 = r1[1:]
		} else if len(r2) > 0 {
			m = appendDistinct(m, r2[0])
			r2 = r2[1:]
		}
	}
	return m
}

func BrockenKeyboard(keys string) string {
	bIndex := 0
	BIndex := 0
	for pos, char := range keys {
		if unicode.IsUpper(char) {
			BIndex = pos
		} else {
			bIndex = pos
		}
	}
	fmt.Println(bIndex, "  ", BIndex)
	return keys
}

func PrintResult(name string, f func([]int) []int, arr []int, printArr bool, wg *sync.WaitGroup) {
	fmt.Printf("%s started \n", name)
	start := time.Now()
	arr = f(arr)
	end := time.Since(start)
	fmt.Printf("%s done in %d Microseconds \n", name, end.Microseconds())
	if printArr {
		fmt.Printf("%s: %v\n", name, arr)
	}
	wg.Done()
}
