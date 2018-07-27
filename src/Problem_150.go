/*
Evaluate Reverse Polish Notation
Evaluate the value of an arithmetic expression in Reverse Polish Notation.
Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Note:
Division between two integers should truncate toward zero.
The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.

Example 1:
Input: ["2", "1", "+", "3", "*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Example 2:
Input: ["4", "13", "5", "/", "+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

Example 3:
Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
Output: 22
Explanation:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/

package main

// Use stack. Data-structure traditional problem.
func evalRPN(tokens []string) int {
	var Stack []int
	for _, s := range tokens {
		if s != "+" && s != "-" && s != "*" && s != "/" {
			num, _ := strconv.Atoi(s)
			Stack = append(Stack, num)
			continue
		}
		var1, var2 := Stack[len(Stack) - 1], Stack[len(Stack) - 2]
		Stack = Stack[:len(Stack) - 2]
		switch s {
		case "+": Stack = append(Stack, var2 + var1)
		case "-": Stack = append(Stack, var2 - var1)
		case "*": Stack = append(Stack, var2 * var1)
		case "/": Stack = append(Stack, var2 / var1)
		}
	}
	return Stack[0]
}