package main

import (
	"fmt"
	"math"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	} else {
		return strconv.Itoa(l.Val) + "->" + l.Next.String()
	}
}

func main() {
	n3 := &ListNode{5, nil}
	n2 := &ListNode{4, n3}
	n1 := &ListNode{1, n2}

	n33 := &ListNode{4, nil}
	n22 := &ListNode{3, n33}
	n11 := &ListNode{1, n22}

	n222 := &ListNode{6, nil}
	n111 := &ListNode{2, n222}

	arr := []*ListNode{n1, n11, n111}

	res := mergeKLists(arr)

	fmt.Println(n1.String())
	fmt.Println(n11.String())
	fmt.Println(n111.String())

	if res != nil {
		fmt.Println(res.String())
	}

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	idxs := map[int]bool{}
	for idx, l := range lists {
		if l != nil {
			idxs[idx] = true
		}
	}
	if len(idxs) == 0 {
		return nil
	}

	var first *ListNode
	var curItem *ListNode
	for {
		minVal := math.MaxInt
		minIdx := -1
		for i := range idxs {
			if lists[i].Val < minVal {
				minVal = lists[i].Val
				minIdx = i
			}
		}
		if first == nil {
			first = lists[minIdx]
		}
		if curItem != nil {
			curItem.Next = lists[minIdx]
		}
		curItem = lists[minIdx]
		lists[minIdx] = lists[minIdx].Next

		if lists[minIdx] == nil {
			if len(idxs) > 1 {
				delete(idxs, minIdx)
			} else {
				break
			}
		}
	}
	return first
}
