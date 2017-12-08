/*
Valid Parentheses
Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
	determine if the input string is valid.

The brackets must close in the correct order,
	"()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/

package main

func isValid(s string) bool {
	stack, top := make([]rune, len(s)/2), -1
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			top++
			if top == len(stack) {
				return false
			}
			stack[top] = v
		case ')':
			if top == -1 || stack[top] != '(' {
				return false
			}
			top--
		case '}', ']':
			if top == -1 || stack[top] != v-2 {
				return false
			}
			top--
		}
	}
	return top == -1
}
