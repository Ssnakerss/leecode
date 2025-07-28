package main

import (
	"fmt"

	"github.com/Ssnakerss/leecode/funcs"
)

func main() {
	n3 := &funcs.ListNode{Val: 5, Next: nil}
	n2 := &funcs.ListNode{Val: 4, Next: n3}
	n1 := &funcs.ListNode{Val: 1, Next: n2}

	n33 := &funcs.ListNode{Val: 4, Next: nil}
	n22 := &funcs.ListNode{Val: 3, Next: n33}
	n11 := &funcs.ListNode{Val: 1, Next: n22}

	n222 := &funcs.ListNode{Val: 6, Next: nil}
	n111 := &funcs.ListNode{Val: 2, Next: n222}

	arr := []*funcs.ListNode{n1, n11, n111}

	fmt.Println(n1.String())
	fmt.Println(n11.String())
	fmt.Println(n111.String())

	res := funcs.MergeKLists2(arr)
	if res != nil {
		fmt.Println(res.String())
	}

}
