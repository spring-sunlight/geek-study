package example_stack_queue

import "strconv"

type evalRPNStack struct {
	stack []string
}

func newEvalRPNStack() *evalRPNStack {
	return &evalRPNStack{
		stack: []string{},
	}
}

func push(e *evalRPNStack, str string) {
	e.stack = append(e.stack, str)
}

func pop(e *evalRPNStack) {
	e.stack = e.stack[:len(e.stack)-1]
}

func top(e *evalRPNStack) string {
	return e.stack[len(e.stack)-1]
}

func evalRPN(tokens []string) int {
	stack := newEvalRPNStack()

	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			push(stack, tokens[i])
		} else {
			first := top(stack)
			pop(stack)
			second := top(stack)
			pop(stack)
			push(stack, calculate(first, second, tokens[i]))
		}
	}
	res, _ := strconv.Atoi(top(stack))
	return res
}

func calculate(first, second, operation string) string {
	n1, _ := strconv.Atoi(first)
	n2, _ := strconv.Atoi(second)
	switch operation {
	case "+":
		return strconv.Itoa(n1 + n2)
	case "-":
		return strconv.Itoa(n2 - n1)
	case "*":
		return strconv.Itoa(n1 * n2)
	case "/":
		return strconv.Itoa(n2 / n1)
	}
	return ""
}
