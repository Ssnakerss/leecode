package funcs

import "strconv"

func CountSpecialNumbers(n int) int {
	len := len(strconv.Itoa(n))
	digits := [10]bool{}
	cnt, _ := makeSpecialNumbers(true, len, 0, digits[:], n)
	return cnt
}

func makeSpecialNumbers(canUseZero bool, size int, s int, digits []bool, target int) (int, bool) {
	cnt := 0
	for i := 0; i < len(digits); i++ {
		if !digits[i] {
			s1 := s*10 + i
			if size > 1 {
				if !(canUseZero && i == 0) {
					digits[i] = true
				}
				if canUseZero {
					canUseZero = i == 0
				}
				c, finish := makeSpecialNumbers(canUseZero, size-1, s1, digits, target)
				cnt += c
				if finish {
					return cnt, true
				}
				digits[i] = false
			} else {
				if s1 <= target {
					if s1 != 0 {
						cnt++
					}
				} else {
					return cnt, true
				}
			}
		}
	}
	return cnt, false
}

func CountSpecialNumbers2(n int) int {
	used := [10]bool{}
	cleanUsed := func() {
		for i := range used {
			used[i] = false
		}
	}
	count := 0
	for i := 1; i <= 9; i++ {
		used[i] = true
		count += dfs(used, i, n)
		cleanUsed()
	}
	return count
}

func dfs(used [10]bool, current int, limit int) int {
	// fmt.Println(current)
	if current > limit {
		return 0
	}
	answ := 1
	for i := 0; i <= 9; i++ {
		if !used[i] {
			used[i] = true
			answ += dfs(used, current*10+i, limit)
			used[i] = false
		}
	}
	return answ
}

// func countSpecialNumbers(n int) int {
// 	if n < 10 {
// 		return n
// 	}
// 	res := 9
// 	resC := make(chan bool)
// 	wg := &sync.WaitGroup{}
// 	go func() {
// 		for b := range resC {
// 			if b {
// 				res++
// 			}
// 		}
// 	}()

// 	for i := 10; i <= n; i++ {
// 		wg.Add(1)
// 		go func() {
// 			resC <- checkIfSpecial2(i)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	close(resC)
// 	return res
// }

// func checkIfSpecial(n int) bool {
// 	if n < 10 {
// 		return true
// 	}
// 	m := map[int]bool{}
// 	for n > 0 {
// 		if m[n%10] {
// 			return false
// 		}
// 		m[n%10] = true
// 		n = n / 10
// 	}
// 	return true
// }

// func checkIfSpecial2(n int) bool {
// 	m := map[rune]bool{}
// 	ns := strconv.Itoa(n)
// 	for _, r := range ns {
// 		if m[r] {
// 			return false
// 		} else {
// 			m[r] = true
// 		}
// 	}
// 	return true
// }
