/*
Gray Code
The gray code is a binary numeral system where two successive values differ in only one bit.
Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code.
A gray code sequence must begin with 0.
For example, given n = 2, return [0,1,3,2]. Its gray code sequence is:

00 - 0
01 - 1
11 - 3
10 - 2
Note:
For a given n, a gray code sequence is not uniquely defined.
For example, [0,2,3,1] is also a valid gray code sequence according to the above definition.
For now, the judge is able to judge based on one instance of gray code sequence. Sorry about that.
*/

package main

// 二进制转换格雷码  Gi = Bi ^ Bi+1
func grayCode(n int) []int {
	m := 1 << uint(n)
	result := make([]int, m)
	for i := range result {
		result[i] = i ^ (i >> 1)
	}
	return result
}

// 递归解决(反置+最高位)
func grayCode1(n int) []int {
	if n == 0 {
		return []int{0}
	} else {
		x := grayCode(n - 1)
		y, base, idx := make([]int,len(x)), 1 << uint(n - 1), len(x) - 1
		for i := range y {
			y[i] = x[idx] + base
			idx--
		}
		return append(x,y...)
	}
}