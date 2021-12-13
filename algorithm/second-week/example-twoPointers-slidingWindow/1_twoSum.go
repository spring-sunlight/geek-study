package example_twoPointers_slidingWindow

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	index := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		index[i] = []int{nums[i], i}
	}
	sort.Slice(index, func(i, j int) bool {
		return index[i][0] < index[j][0]
	})
	for left, right := 0, len(nums)-1; left < right; {
		if index[left][0]+index[right][0] < target {
			left++
		} else if index[left][0]+index[right][0] > target {
			right--
		} else {
			return []int{index[left][1], index[right][1]}
		}
	}
	return []int{}
}
