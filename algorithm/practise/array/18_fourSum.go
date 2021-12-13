package array

import (
	"sort"
)

//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//[-2,-1,0,0,1,2]
func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ans = append(ans, threeSum(nums[i+1:], target-nums[i], target)...)
	}
	return ans
}

func threeSum(nums []int, threeTarget, FourTarget int) [][]int {
	ans := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ans = append(ans, twoSum(nums[i+1:], threeTarget-nums[i], threeTarget, FourTarget)...)
	}
	return ans
}

func twoSum(nums []int, twoTarget, threeTarget, FourTarget int) [][]int {
	ans := [][]int{}
	for left, right := 0, len(nums)-1; left < right; {
		if left > 0 && nums[left] == nums[left-1] {
			left++
			continue
		}
		if right < len(nums)-1 && nums[right] == nums[right+1] {
			right--
			continue
		}
		if nums[left]+nums[right] < twoTarget {
			left++
		} else if nums[left]+nums[right] > twoTarget {
			right--
		} else {
			ans = append(ans, []int{nums[left], nums[right], threeTarget - twoTarget, FourTarget - threeTarget})
			left++
			right--
		}
	}
	return ans
}
