package example_stack_queue

/*
type calculateStack struct {
	stack []string
}

func newCalculateStack() *calculateStack {
	return &calculateStack{stack: []string{}}
}

func push(c *calculateStack, val string) {
	c.stack = append(c.stack, val)
}

func pop(c *calculateStack) {
	c.stack = c.stack[:len(c.stack)-1]
}

func top(c *calculateStack) string {
	return c.stack[len(c.stack)-1]
}

func lenStack(c *calculateStack) int {
	return len(c.stack)
}

func isEmpty224(c *calculateStack) bool {
	return len(c.stack) == 0
}

func isNum(s string) bool {
	return s != "+" && s != "-" && s != "(" && s != ")"
}

func calculate(s string) int {
	stack := newCalculateStack()
	symbols := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			if isEmpty(stack) {
				stack.stack = append(stack.stack,string(s[i]))
			}else if isNum(top(stack)){
				stack.stack[len(stack.stack)-1] += string(s[i])
			} else {
				push(stack, string(s[i]))
			}
			//fmt.Println("stack50: ", stack.stack)
		} else {
			push(stack, string(s[i]))
			//fmt.Println("stack53: ", stack.stack)
			if s[i] == '(' {
				symbols = append(symbols, lenStack(stack)-1)
			} else if s[i] == ')' {
				//fmt.Println("symbols: ", symbols)
				nums := stack.stack[symbols[len(symbols)-1]:]
				res := getResult(nums)
				//fmt.Println("res: ", res)
				stack.stack = stack.stack[:symbols[len(symbols)-1]]
				symbols = symbols[:len(symbols)-1]
				push(stack, res)
				//fmt.Println("stack64: ", stack.stack)
			}
		}
	}

	res,_:= strconv.Atoi(getResult(stack.stack))

	return res
}

func getResult(nums []string) string {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return nums[0]+nums[1]
	}

	if nums[0] != "(" {
		nums = append([]string{"("}, nums...)
		nums = append(nums, ")")
	}
	if nums[1] == "-" || nums[1] == "+" {
		nums = append(append(nums[:1], "0"), nums[1:]...)
	}
	res, _ := strconv.Atoi(nums[1])
	for i := 3; i < len(nums); i += 2 {
		num, _ := strconv.Atoi(nums[i])
		if nums[i-1] == "+" {
			res = res + num
		} else {
			res = res - num
		}
	}
	return strconv.Itoa(res)
}

*/
