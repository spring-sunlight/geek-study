package string

import "fmt"

func convert(s string, numRows int) string {
	//输入：s = "PAYPALISHIRING", numRows = 4
	//输出："PINALSIGYAHRPI"
	//解释：
	//P     I    N
	//A   L S  I G
	//Y A   H R
	//P     I

	//P A Y P A L I S H I R  I  N  G
	//0 1 2 3 4 5 6 7 8 9 10 11 12 13

	//思路:
	//1.开辟 numRows个数组存储每行的结果
	//2.每 numRows-1个元素是一个方向,用 flag 记录
	//3.flag = 1 代表向下,正着存
	//  flag = 0 代表向上,倒着存
	//  每numRows-1个元素换一个方向
	//4.把每个数组结果存储起来

	if numRows == 1 {
		return s
	}
	res := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = []byte{}
	}
	flag := 1 //up
	row := 0
	for i := 0; i < len(s); i++ {
		if flag == 1 {
			res[row] = append(res[row], s[i])
			row++
			if row == numRows-1 {
				flag = 0
			}
		} else {
			res[row] = append(res[row], s[i])
			row--
			if row == 0 {
				flag = 1
			}
		}
	}
	ans := ""
	for i := 0; i < len(res); i++ {
		ans += string(res[i])
		//fmt.Println(string(res[i]))
	}

	return ans
}

// @solution-sync:end

func main() {
	var s = "AB"
	var numRows = 2

	var result = convert(s, numRows)
	fmt.Printf("%v\n", result)
}
