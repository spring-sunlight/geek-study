package seventh_week

//参考题解

// 思路 动态规划法 dp[i] 表示 i下标能不能到达
// dp[i] 能不能到达 取决于 dp[0:i-1] 中能到达的地方能不能到达dp[i]
//  dp[j] 为 dp[i]  前的某个位置  那么当 dp[j]=true 并且 nums[j]>=i-j 那么 i就能到达
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	// 第一个肯定能到达(起始条件)
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}
