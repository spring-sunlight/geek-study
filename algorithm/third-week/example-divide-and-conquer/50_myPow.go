package example_divide_and_conquer

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	ans := float64(0)
	temp := myPow(x, n/2)
	ans = temp * temp
	if n%2 == 1 {
		ans = ans * x
	}
	return ans
}

func main() {
	x := float64(2)
	fmt.Println(myPow(x, 10))
}
