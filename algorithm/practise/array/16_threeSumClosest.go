package array

import (
	"fmt"
	"sort"
)

//双指针扫描
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	if len(nums) == 3 || nums[0]+nums[1]+nums[2] > target {
		return nums[0] + nums[1] + nums[2]
	}
	length := len(nums)
	if nums[length-1]+nums[length-2]+nums[length-3] < target {
		return nums[length-1] + nums[length-2] + nums[length-3]
	}
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < length-2; i++ {
		left := i + 1
		right := length - 1
		twoSumTarget := target - nums[i]
		for left < right {
			twoSum := nums[left] + nums[right]
			if twoSum < twoSumTarget {
				left++
			} else if twoSum > twoSumTarget {
				right--
			} else {
				return target
			}
			if abs(ans-target) > abs(nums[i]+twoSum-target) {
				ans = nums[i] + twoSum
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{1, 2, 5, 10, 11}
	//nums = []int{0, 0, 0}
	fmt.Println(threeSumClosest(nums, 12))
}
