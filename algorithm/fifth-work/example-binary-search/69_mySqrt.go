package example_binary_search

//输入：x = 8
//输出：2
//解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

//寻找最后一个 ans*ans <= target 的数
func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := (left + right + 1) >> 1
		if mid <= x/mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}
