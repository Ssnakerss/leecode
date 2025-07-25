package main

import "fmt"

const (
	limit   = 100000
	arrSize = 50
)

func main() {
	// r0 := funcs.GenRandArray(arrSize, limit)

	// fmt.Printf("%v\n", r0)
	// fmt.Println("starting...")
	// wg := sync.WaitGroup{}
	// // wg.Add(1)
	// // go printResult("bubble", sortArrBubble, append(make([]int, 0), r0...), &wg)
	// // wg.Add(1)
	// // go printResult("min", sortArrMin, append(make([]int, 0), r0...), &wg)

	// wg.Add(1)
	// go funcs.PrintResult("insert", funcs.SortArrInsert, append(make([]int, 0), r0...), true, &wg)
	// // wg.Add(1)
	// // go funcs.PrintResult("quick", funcs.QuickSort, append(make([]int, 0), r0...), &wg)
	// // wg.Add(1)
	// // go funcs.PrintResult("merge", funcs.MergeSort, append(make([]int, 0), r0...), &wg)
	// wg.Wait()

	// r1 := []int{1, 3, 3, 3, 4, 5, 7, 9, 9, 9, 10, 10, 10, 10}
	// r2 := []int{2, 4, 6, 8, 9, 10}
	// fmt.Printf("%v \n", funcs.Merge(r1, r2))
	// fmt.Printf("%v \n", funcs.MergeSortedLists(r1, r2))

	// idxs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//       01234567890123456789
	// keys := funcs.GetRandString(100000) //"AaCaCaaAbbbAaB"
	// start := time.Now()
	// funcs.BrockenKeys(keys)
	// duration := time.Since(start)
	// fmt.Println(duration)
	// fmt.Println(string(res))

	// keys := funcs.GetRandString(10_000_000) //"AaCBDEbFB" //"AaCcDdeEbbbFGfB" //
	// if len(keys) < 300 {

	// 	fmt.Println(keys)
	// }

	// start := time.Now()
	// st := funcs.BrockenKeysStack(keys)
	// end := time.Since(start)
	// fmt.Println("stack", end)

	// start = time.Now()
	// na := funcs.BrockenKeysNaive(keys)
	// end = time.Since(start)
	// fmt.Println("naive", end)

	// start = time.Now()
	// rv := funcs.BrockenKeysRev(keys)
	// end = time.Since(start)
	// fmt.Println("rev", end)

	// fmt.Println("st-na", st == na, "na-rv", na == rv, "st-rv", st == rv)
	// if len(keys) < 300 {
	// 	fmt.Println(st)
	// 	fmt.Println(na)
	// 	fmt.Println(rv)
	// }
	// fmt.Println(string(funcs.Reverse([]rune("ABCDE"))))
	nums1 := []int{1, 3, 4, 6}
	nums2 := []int{2, 5, 7, 8}
	i, j := 0, 0
	prev, next := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j >= len(nums2) {
			prev, next = next, nums1[i]
			i++
		} else if j < len(nums2) && i >= len(nums1) {
			prev, next = next, nums2[j]
			j++
		} else {
			if nums1[i] < nums2[j] {
				prev, next = next, nums1[i]
				i++
			} else {
				prev, next = next, nums2[j]
				j++
			}
		}

		if i+j > (len(nums1)+len(nums2))/2 {
			if (len(nums1)+len(nums2))%2 == 0 {
				fmt.Printf(">>%d: %f \n", i+j, float64(prev+next)/2)
				return
			} else {
				fmt.Printf(">>%d: %d \n", i+j, next)
				return
			}
		}
		fmt.Printf("%d: %d %d\n", i+j, prev, next)
	}
	fmt.Println((len(nums1)+len(nums2))/2, (len(nums1)+len(nums2))%2)
}
