package string

import "fmt"

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		if res < (-1<<31) || res > 1<<31-1 {
			return 0
		}
		x = x / 10
	}
	return res
}

// @solution-sync:end

func main() {
	var x = 0

	var result = reverse(x)
	fmt.Printf("%v\n", result)
}
