/*
Reverse Integer
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
Input: 123
Output:  321

Example 2:
Input: -123
Output: -321

Example 3:
Input: 120
Output: 21

Note:
Assume we are dealing with an environment
which could only hold integers within the 32-bit signed integer range.
For the purpose of this problem, assume that your function returns 0
when the reversed integer overflows.
*/

package main

import "math"

func reverse(x int) int {
	symbol := (x>>31)<<1 + 1
	abs := x * symbol
	res := 0
	for abs != 0 {
		res = res*10 + abs%10
		abs /= 10
	}
	if res > math.MaxInt32 {
		return 0
	}
	return res * symbol
}
