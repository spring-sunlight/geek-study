package example_hash_set_map

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := m[target-nums[i]]; ok {
			return []int{i, index}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

func TestTwoSum() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(twoSum(nums, 9))
}
