package example_array

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j = j + 1
			nums[j] = nums[i]
		}
	}
	nums = nums[:j+1]
	fmt.Println(nums)
	return len(nums)
}

func testRemoveDuplicates() {
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
}
