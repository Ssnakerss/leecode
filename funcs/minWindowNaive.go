package funcs

import "maps"

type result struct {
	window string
	length int
}

func makeKeysMap(t string) map[string]int {
	if len(t) == 0 {
		return nil
	}
	tKey := map[string]int{}
	for _, r := range t {
		if _, ok := tKey[string(r)]; !ok {
			tKey[string(r)] = 1
		} else {
			tKey[string(r)]++
		}
	}
	return tKey
}
func findWindow(s string, keys map[string]int) result {
	start := -1
	end := -1
	lens := len(s)
	for i, r := range s {
		if i > lens-len(keys) {
			break
		}
		//match kekys list
		if v, ok := keys[string(r)]; ok {
			if v > 0 {
				keys[string(r)]--
				v--
				if v == 0 {
					delete(keys, string(r))
				}
				if start == -1 {
					start = i
				}
				end = i
			}
		}
		if len(keys) == 0 {
			break
		}
	}
	if start == -1 || end == -1 || len(keys) > 0 {
		return result{"", 0}
	}
	return result{s[start : end+1], end - start + 1}
}

// Minimum window substring
func MinWindowNaive(s string, t string) string {
	tKey := makeKeysMap(t)
	res := make(chan *result)
	goCnt := 0
	lens := len(s)
	lent := len(t)

	for i, r := range s {
		if i > lens-lent {
			break
		}
		if _, ok := tKey[string(r)]; ok {
			tKeyCopy := map[string]int{}
			maps.Copy(tKeyCopy, tKey)
			goCnt++
			go func() {
				r := findWindow(s[i:], tKeyCopy)
				res <- &r
			}()
		}
	}
	minWin := result{}
	for goCnt > 0 {
		r := <-res
		// fmt.Printf("%s \n", r.window)
		if (r.length < minWin.length && r.length > 0) || minWin.length == 0 {
			minWin = *r
		}
		goCnt--
	}
	close(res)
	return minWin.window
}
