package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("1", reflect.DeepEqual(generateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"}))
	fmt.Println("2", reflect.DeepEqual(generateParenthesis(1), []string{"()"}))
}

func generateParenthesis(n int) []string {
	// opened and closed parenthesis
	// there should never be more closing than open
	// there should never be more opened than n
	// exit when open == closed == n
	// add open or closed to a stack and join strings together and return as res

	var res, stack []string
	var backtrack func(opened, closed int)
	backtrack = func(opened, closed int) {
		fmt.Println(stack)
		// return condition
		// append open or close to stack
		// call backtrack to drill down on decision tree
		// pop since stack is global var and it needs to be shared with other instances
		if opened == closed && closed == n {
			res = append(res, strings.Join(stack, ""))
		}
		if n > opened {
			stack = append(stack, "(")
			backtrack(opened+1, closed)
			pop(&stack)
		}
		if opened > closed {
			stack = append(stack, ")")
			backtrack(opened, closed+1)
			pop(&stack)
		}
	}
	backtrack(0, 0)
	fmt.Println(res)
	return res
}
func pop(list *[]string) {
	*list = (*list)[:len(*list)-1]
}
