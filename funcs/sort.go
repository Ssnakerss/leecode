package funcs

func SortArrBubble(arr []int) []int {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}
	return arr
}

// sortArr sorts an array of integers
func SortArrMin(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
func SortArrInsert(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	sarr := make([]int, 0)
	sarr = append(sarr, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] <= sarr[0] {
			sarr = ArrIns(sarr, 0, arr[i])
		} else if arr[i] >= sarr[len(sarr)-1] {
			sarr = ArrIns(sarr, len(sarr), arr[i])
		} else {
			mid := len(sarr) / 2
			if arr[i] > sarr[mid] {
				for j := mid; j < len(sarr); j++ {
					if arr[i] <= sarr[j] {
						sarr = ArrIns(sarr, j, arr[i])
						break
					}
				}

			} else {
				for j := mid; j >= 0; j-- {
					if arr[i] >= sarr[j] {
						sarr = ArrIns(sarr, j+1, arr[i])
						break
					}
				}
			}

		}
	}
	return sarr
}

// QuickSort
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less, greater []int
		for _, v := range arr[1:] {
			if v <= pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
		return append(append(QuickSort(less), pivot), QuickSort(greater)...)
	}
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		mid := len(arr) / 2
		return Merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
	}
}
