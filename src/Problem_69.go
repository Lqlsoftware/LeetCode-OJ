/*
Sqrt(x)
Implement int sqrt(int x).
Compute and return the square root of x.
x is guaranteed to be a non-negative integer.

Example 1:
Input: 4
Output: 2

Example 2:
Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we want to return an integer, the decimal part will be truncated.
*/

package main

func mySqrt(x int) int {
	i := 1
	for i * i <= x {
		i++
	}
	return i - 1
}

// BS
func mySqrt1(x int) int {
	if x <= 1 {
		return x
	}
	start, end := 0, x
	for start < end {
		mid := (start + end) / 2
		sqrtX := mid * mid
		if sqrtX < x {
			start = mid + 1
		} else if sqrtX > x {
			end = mid
		} else {
			return mid
		}
	}
	return end - 1
}