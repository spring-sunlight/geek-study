package example_array

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if index1 < 0 {
			nums1[i] = nums2[index2]
			index2--
			continue
		}
		if index2 < 0 {
			break
		}
		if nums1[index1] < nums2[index2] {
			nums1[i] = nums2[index2]
			index2--
		} else {
			nums1[i] = nums1[index1]
			index1--
		}
	}
	fmt.Println(nums1)
}

func testMerge() {
	nums1 := []int{2, 3, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge(nums1, 3, nums2, 3)
}
