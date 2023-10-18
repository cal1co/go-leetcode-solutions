package main

import "fmt"

func main() {
	fmt.Println("1", isValid("()"))
	fmt.Println("2", isValid("()[]{}"))
	fmt.Println("3", isValid("(]"))
}

func isValid(s string) bool {
	// have a map of the brackets {')': '('}
	// loop through string
	// if not present in bracket map, add to stack and continue
	// if is in map and top of stack != it, return false
	// pop from stack
	pMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	for _, p := range []byte(s) {
		val, ok := pMap[p]
		if !ok {
			stack = append(stack, p)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if val != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]

	}
	return len(stack) == 0
}
