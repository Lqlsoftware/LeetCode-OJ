/*
Multiply Strings
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2.

Note:
The length of both num1 and num2 is < 110.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

package main

func multiply(num1 string, num2 string) string {
	res := make([]byte,len(num1) + len(num2))
	for i := len(num1) - 1;i >= 0;i-- {
		for j := len(num2) - 1;j >= 0;j-- {
			v := (num1[i] - '0') * (num2[j] - '0')
			p1,p2 := i + j,i + j + 1
			lower := res[p2] + v % 10
			res[p1] += v / 10 + lower / 10
			res[p2] = lower % 10
		}
	}
	flag := 0
	for flag = 0;flag < len(res) && res[flag] == 0;flag++ {}
	if flag == len(res) {
		flag--
	}
	for idx := flag;idx < len(res);idx++ {
		res[idx] += '0'
	}
	return string(res[flag:])
}
