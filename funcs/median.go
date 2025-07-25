package funcs

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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

				return float64(prev+next) / 2
			} else {

				return float64(next)
			}
		}

	}

	return 0
}
