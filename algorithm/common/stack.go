package common

import "fmt"

type stack struct {
	nums []interface{}
}

func NewStack() *stack {
	return &stack{
		nums: []interface{}{},
	}
}

func (stack *stack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.nums[len(stack.nums)-1]
}

func (stack *stack) IsEmpty() bool {
	return len(stack.nums) == 0
}

func (stack *stack) Push(val interface{}) {
	stack.nums = append(stack.nums, val)
}

func (stack *stack) Pop() {
	stack.nums = stack.nums[:len(stack.nums)-1]
}

func (stack *stack) PrintStack() {
	fmt.Println("==================")
	fmt.Println("stack 的长度为:", len(stack.nums))
	fmt.Println("stack 的值为:", stack.nums)
}
