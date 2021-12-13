package example_prefixSum

func maxSubArray(nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	minNum := 0
	ans := nums[0]
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
		diff := prefixSum[i] - minNum
		if diff > ans {
			ans = diff
		}
		if prefixSum[i] < minNum {
			prefixSum[i] = minNum
		}
	}
	return ans
}
