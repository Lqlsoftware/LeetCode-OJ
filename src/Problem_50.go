/*
Pow(x, n)
Implement pow(x, n).

Example 1:
Input: 2.00000, 10
Output: 1024.00000

Example 2:
Input: 2.10000, 3
Output: 9.26100
*/

package main

// 递归
func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(1 / x, -n)
	} else if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n % 2 == 1 {
		return x * myPow(x * x,n / 2)
	}
	return myPow(x * x,n / 2)
}

// 循环 减少了递归堆栈的空间
func myPow1(x float64, n int) float64 {
	if n < 0 {
		return myPow(1 / x, -n)
	}
	res, current := float64(1), x
	for n > 0 {
		if n % 2 == 1 {
			res *= current
		}
		n /= 2
		current *= current
	}
	return res
}