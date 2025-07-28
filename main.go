package main

import (
	"fmt"
	"maps"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	tKey := makeKeysMap(t)

	for i, r := range s {
		if i >= len(s)-len(t) {
			break
		}
		if _, ok := tKey[string(r)]; ok {
			tKeyCopy := map[string]int{}
			maps.Copy(tKeyCopy, tKey)
			w, l := findWindow(s[i:], tKeyCopy)
			fmt.Printf("%s: %d\n", w, l)
		}
	}
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
func findWindow(s string, keys map[string]int) (string, int) {
	start := -1
	end := -1
	for i, r := range s {
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
	}
	if start == -1 || end == -1 || len(keys) > 0 {
		return "", 0
	}
	return s[start : end+1], end - start + 1
}

// // Minimum window substring
// func minWindow(s string, t string) string {
// 	return ""
// }
