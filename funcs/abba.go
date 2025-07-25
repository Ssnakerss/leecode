package funcs

import "errors"

type stackOfRune struct {
	Stack []rune
}

func (s *stackOfRune) Push(r rune) {
	s.Stack = append(s.Stack, r)
}

func newStackOfRune() *stackOfRune {
	return &stackOfRune{
		Stack: make([]rune, 0),
	}
}

func (s *stackOfRune) Pop() (rune, error) {
	if len(s.Stack) == 0 {
		return 0, errors.New("empty stack")
	} else {
		ret := s.Stack[len(s.Stack)-1]
		s.Stack = s.Stack[:len(s.Stack)-1]
		return ret, nil
	}
}

func (s *stackOfRune) String() string {
	if len(s.Stack) == 0 {
		return "empty stack"
	} else {
		return string(s.Stack)
	}
}

func (s *stackOfRune) Peek() rune {
	if len(s.Stack) == 0 {
		return -1
	} else {
		return (s.Stack[len(s.Stack)-1])
	}
}

func countScore(what string, score int, where string) (int, string) {
	stack := newStackOfRune()
	sum := 0
	for _, r := range where {
		lastRune := stack.Peek()
		if string(lastRune)+string(r) == what {
			sum += score
			stack.Pop()
			continue
		}
		stack.Push(r)
	}
	return sum, stack.String()
}
func maximumGain(s string, x int, y int) int {
	sum := 0
	if x > y {
		sum, s = countScore("ab", x, s)
		sum1, _ := countScore("ba", y, s)
		sum += sum1
	} else {
		sum, s = countScore("ba", y, s)
		sum1, _ := countScore("ab", x, s)
		sum += sum1
	}
	return sum
}
