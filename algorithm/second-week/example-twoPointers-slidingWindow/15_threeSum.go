package example_twoPointers_slidingWindow

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		otherTwoSum := innerTwoSum(nums[i+1:], -nums[i])
		if len(otherTwoSum) != 0 {
			ans = append(ans, otherTwoSum...)
		}
	}
	return ans
}

func innerTwoSum(nums []int, target int) [][]int {
	res := [][]int{}
	for left, right := 0, len(nums)-1; left < right; {
		if left > 0 && right < len(nums)-1 {
			if nums[left] == nums[left-1] {
				left++
				continue
			}
			if nums[right] == nums[right+1] {
				right--
				continue
			}
		}
		if nums[left]+nums[right] < target {
			left++
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			res = append(res, []int{-target, nums[left], nums[right]})
			left++
			right--
		}
	}
	return res
}

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
