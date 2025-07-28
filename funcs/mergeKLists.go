package funcs

import (
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

func MergeKLists(lists []*ListNode) *ListNode {
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

func MergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return Merge2List(lists[0], lists[1])
	} else {
		return Merge2List(MergeKLists2(lists[:len(lists)/2]), MergeKLists2(lists[len(lists)/2:]))
	}
}

func Merge2List(list1 *ListNode, list2 *ListNode) *ListNode {
	var first *ListNode
	var curItem *ListNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			if first == nil {
				first = list1
				curItem, list1 = list1, list1.Next
			} else {
				curItem.Next, curItem, list1 = list1, list1, list1.Next
			}
		} else {
			if first == nil {
				first = list2
				curItem, list2 = list2, list2.Next
			} else {
				curItem.Next, curItem, list2 = list2, list2, list2.Next
			}
		}
	}
	if list1 != nil {
		if first == nil {
			first = list1
		} else {
			curItem.Next = list1
		}
	}
	if list2 != nil {
		if first == nil {
			first = list2
		} else {
			curItem.Next = list2
		}
	}

	return first
}
