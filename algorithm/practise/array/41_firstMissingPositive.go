package array

import "fmt"

/*
	题目:给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

	请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

	示例 1：

	输入：nums = [1,2,0]
	输出：3
	示例 2：

	输入：nums = [3,4,-1,1]
	输出：2
	示例 3：

	输入：nums = [7,8,9,11,12]
	输出：1

	提示：

	1 <= nums.length <= 5 * 105
	-231 <= nums[i] <= 231 - 1

	思路:如果数组中有 负数,0,大于数组长度的数,缺失的数肯定在 [1,n+1]之间,n 是数组长度
		[3,4,-1,-2,1,5,16,0,2,0]
		1.先查看 1 在不在数组中,不在的话返回 1,是为了下面的操作做准备
		2.将数组中负数,0,大于数组长度的数置为 1,数组长度为10,此时数组中都是正数
		[3,4,1,1,1,5,1,1,2,1]
		3.从头开始把数组中nums[i]绝对值位置上的数置为负数
		[3,-4,-1,-1,-1,5,1,1,2,1]
		4.因为已经算过 1 必定在数组中了,因此第 0 个位置上的数就不用算了, 从数组第 1 个位置上的数开始算起第一个为正数的位置就是答案
		5.特殊情况[1,2,3,4,5],替换后为[-1,-2,-3,-4,-5],都是负数,返回数组长度+1 即可
*/

func firstMissingPositive(nums []int) int {
	//看 1 在不在数组中,将数组中负数,0,大于数组长度的数置为 1
	flag := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			flag = true
		} else if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}
	if !flag {
		return 1
	}
	//从头开始把数组中nums[i]绝对值位置上的数置为负数
	for i := 0; i < len(nums); i++ {
		nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
	}
	fmt.Println(nums)
	//从数组第 1 个位置上的数开始算起第一个为正数的位置就是答案
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	//特殊情况替换后都是负数,返回数组长度+1 即可
	return len(nums) + 1
}

func main() {
	nums := []int{1, 2, 0}
	nums1 := []int{3, 4, -1, 1}
	nums2 := []int{7, 8, 9, 11, 12}
	nums3 := []int{0, 2, 2, 1, 1}
	fmt.Println(firstMissingPositive(nums))
	fmt.Println(firstMissingPositive(nums1))
	fmt.Println(firstMissingPositive(nums2))
	fmt.Println(firstMissingPositive(nums3))
}
