package practise

import "fmt"

//输入：[1, 2, 2, 3, 1]
//输出：2
//解释：
//输入数组的度是2，因为元素1和2的出现频数最大，均为2.
//连续子数组里面拥有相同度的有如下所示:
//[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
//最短连续子数组[2, 2]的长度为2，所以返回2.

func findShortestSubArray(nums []int) int {
	numsCountMap := make(map[int]int)
	numsIndexMap := make(map[int][]int)
	max := 0
	for i := 0; i < len(nums); i++ {
		numsIndexMap[nums[i]] = append(numsIndexMap[nums[i]], i)
		numsCountMap[nums[i]]++
		if numsCountMap[nums[i]] > max {
			max = numsCountMap[nums[i]]
		}
	}
	ans := len(nums)
	for num, index := range numsIndexMap {
		if numsCountMap[num] == max {
			length := index[len(index)-1] - index[0] + 1
			if ans > length {
				ans = length
			}
		}
	}
	return ans
}
func main() {
	nums := []int{}
	fmt.Println(findShortestSubArray(nums))
}
