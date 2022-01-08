package array

import "fmt"

/*
	题目大意:原地删除有序数组中的重复元素,使得每个元素出现最多不超过两次

	例子:输入：nums = [0,0,1,1,1,1,2,3,3]
		输出：7, nums = [0,0,1,1,2,3,3]
		解释：函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。 不需要考虑数组中超出新长度后面的元素。

	思路: 算出从当前开始多少个数相同,如果一个数相同拼接一个数,两个及两个数以上就拼接两个数,
		cur表示字符串开始拼接的位置,start,end 表示每一个相同的子串的起始位置和结束位置,count 是一样的数组出现的次数
	    注意最后 end 出界后还需要判定一次
*/

func removeDuplicates(nums []int) int {
	cur, start, end, count := 0, 0, 0, 0
	for end <= len(nums) {
		if end == len(nums) || nums[end] != nums[start] {
			count = end - start
			if count == 1 {
				nums[cur] = nums[end-1]
				cur++
			} else if count >= 2 {
				nums[cur], nums[cur+1] = nums[start], nums[start]
				cur += 2
			}
			start = end
		}
		end++
	}

	fmt.Println(nums[0:cur])
	return cur
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 3, 3}
	fmt.Println(removeDuplicates(nums))
}
