package funcs

import (
	"errors"
	"unicode"
)

type stack struct {
	Arr []Key
}

func newStack() *stack {
	return &stack{
		Arr: []Key{},
	}
}

func (s *stack) Push(k Key) {
	s.Arr = append(s.Arr, k)
}

func (s *stack) TryPop() (Key, error) {
	if len(s.Arr) == 0 {
		return Key{0, -1}, errors.New("stack is empty")
	}
	val := s.Arr[len(s.Arr)-1]
	s.Arr = s.Arr[:len(s.Arr)-1]
	return val, nil
}
func (s *stack) IsEmpty() bool {
	return len(s.Arr) == 0
}

type Key struct {
	Value rune
	Idx   int
}

func MergeKeysLists(r1 []Key, r2 []Key) []rune {
	m := make([]rune, 0)
	for len(r1) > 0 || len(r2) > 0 {
		if len(r1) > 0 && len(r2) > 0 {
			if r1[0].Idx < r2[0].Idx {
				m = append(m, r1[0].Value)
				r1 = r1[1:]
			} else {
				m = append(m, r2[0].Value)
				r2 = r2[1:]
			}
		} else if len(r1) > 0 {
			m = append(m, r1[0].Value)
			r1 = r1[1:]
		} else if len(r2) > 0 {
			m = append(m, r2[0].Value)
			r2 = r2[1:]
		}
	}
	return m
}

func BrockenKeysStack(Keys string) string {
	lowers := newStack()
	uppers := newStack()
	for pos, k := range Keys {
		switch {
		case k == 'b':
			lowers.TryPop()
		case k == 'B':
			uppers.TryPop()
		default:
			if unicode.IsUpper(rune(k)) {
				uppers.Push(Key{Value: k, Idx: pos})
			} else {
				lowers.Push(Key{Value: k, Idx: pos})
			}
		}
	}
	return string(MergeKeysLists(lowers.Arr, uppers.Arr))
}
