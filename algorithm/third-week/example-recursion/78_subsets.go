package example_recursion

func subsets(nums []int) [][]int {
	ans := [][]int{}
	share := []int{}

	var recur func(i int, nums []int)
	recur = func(i int, nums []int) {
		//1.确定递归边界
		if i == len(nums) {
			ans = append(ans, append([]int{}, share...))
			return
		}

		//2.相似子问题,第 i 个数选或者不选
		//选:放进共享数组,到下一个数
		//不选:直接到下一个数

		//不选
		recur(i+1, nums)

		//选
		share = append(share, nums[i])
		recur(i+1, nums)

		//回退现场,每个分支走完后把 share 还原
		//share里添加了一个值,回退就要减少一个值
		share = share[:len(share)-1]
	}
	recur(0, nums)
	return ans
}
