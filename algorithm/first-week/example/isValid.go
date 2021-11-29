package example

import "fmt"

type stack struct {
	nums []string
}

func newStack() *stack {
	return &stack{
		nums: []string{},
	}
}

func isEmpty(s *stack) bool {
	return len(s.nums) == 0
}

func pushStack(s *stack, val string) {
	s.nums = append(s.nums, val)
}

func popStack(s *stack) bool {
	if isEmpty(s) {
		return false
	}
	s.nums = s.nums[:len(s.nums)-1]
	return true
}

func getTop(s *stack) string {
	if isEmpty(s) {
		return ""
	}
	return s.nums[len(s.nums)-1]
}

func check(s1, s2 string) bool {
	fmt.Println(s1, s2)
	return (s1 == "(" && s2 == ")") || (s1 == "[" && s2 == "]") || (s1 == "{" && s2 == "}")
}

func isValid(s string) bool {
	stack := newStack()
	for i := 0; i < len(s); i++ {
		if isEmpty(stack) {
			pushStack(stack, string(s[i]))
		} else {
			top := getTop(stack)
			if !check(top, string(s[i])) {
				pushStack(stack, string(s[i]))
			} else {
				popStack(stack)
			}
		}
	}
	if isEmpty(stack) {
		return true
	}
	return false
}

func TestIsValid() {
	fmt.Println(isValid("(){}[]"))
}
