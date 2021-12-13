package example_prefixSum

//输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
//输出：[10,55,45,25,25]
//解释：
//航班编号        1   2   3   4   5
//预订记录 1 ：   10  10
//预订记录 2 ：       20  20
//预订记录 3 ：       25  25  25  25
//总座位数：      10  55  45  25  25
//因此，answer = [10,55,45,25,25]

//差分
//差分数组的前缀和就是原数组
func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		ans[bookings[i][0]-1] += bookings[i][2]
		if bookings[i][1] != n {
			ans[bookings[i][1]] -= bookings[i][2]
		}
	}
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}

func main() {
	bookings := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}
	corpFlightBookings(bookings, 5)
}
