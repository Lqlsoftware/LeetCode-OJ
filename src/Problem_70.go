/*
Climbing Stairs
You are climbing a stair case. It takes n steps to reach to the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
Note: Given n will be a positive integer.

Example 1:
Input: 2
Output:  2
Explanation:  There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: 3
Output:  3
Explanation:  There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

package main

// Recursion
func climbStairs(n int) int {
	memo := make([]int,n + 1)
	return solve(&memo,n)
}

func solve(memo *[]int, n int) int {
	if n == 1 || n == 2 {
		return n
	} else if (*memo)[n] != 0 {
		return (*memo)[n]
	} else {
		v1 := solve(memo ,n - 1)
		(*memo)[n-1] = v1
		return v1 + solve(memo ,n - 2)
	}
}

// dp
func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int,n + 1)
	dp[1],dp[2] = 1,2
	for i := 3;i <= n;i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}

// based on dp method: dp[i] = dp[i - 1] + dp[i - 2] -> Fibonacci
func climbStairs2(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	a, b, c := 1, 2, 3
	for i := 3;i <= n;i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

// use formula to calculate Fibonacci: +0.1 to keep result's float->int accuracy
func climbStairs3(n int) int {
	return int((math.Pow(1.6180339887498949, float64(n + 1)) - math.Pow(-0.6180339887498949, float64(n + 1))) / 2.2360679774997898 + 0.1)
}