package array

import (
	"fmt"
	"sort"
)

/*
	题意:求数组的子集,数组中可能有重复元素,不限制顺序
	例子:[1,2,2]
    输出:[[],[1],[2],[1,2],[2,2],[1,2,2]]

 	思路:1.每一个数有选和不选两个选项
		2.如果当前数字与前一个数字相同,
        3.前一个数选了,这个数可以选或者不选,因为求子集,必须有两个数都选的情况,所以前一个数必选
        4.前一个数没选,这个数可以选或者不选,都是相同的结果与上一步重合,因此这种情况舍弃
		5.最后注意还原边界
*/

func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	share := []int{}
	used := make([]int, len(nums))
	sort.Ints(nums)
	var recur func(i int)

	recur = func(i int) {
		if i == len(nums) {
			ans = append(ans, append([]int{}, share...))
			return
		}

		recur(i + 1)

		if i > 0 && used[i-1] == 0 && nums[i] == nums[i-1] {
			return
		}
		if used[i] == 0 {
			share = append(share, nums[i])
			used[i] = 1
			recur(i + 1)
			used[i] = 0
			share = share[:len(share)-1]
		}
	}
	recur(0)
	return ans
}

func main() {
	nums := []int{1, 2, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
