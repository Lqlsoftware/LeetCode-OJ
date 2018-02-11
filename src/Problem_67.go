/*
Add Binary
Given two binary strings, return their sum (also a binary string).

For example,
a = "11"
b = "1"
Return "100".
*/

package main

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}
	res,c := make([]byte, len(a)),byte(0)
	for i,j := len(a) - 1,len(b) - 1;i >= 0;i-- {
		if j >= 0 {
			res[i] = a[i] + b[j] - '0' + c
			j--
		} else {
			res[i] = a[i] + c
		}
		if res[i] > '1' {
			c = 1
			res[i] -= 2
		} else {
			c = 0
		}
	}
	if c == 0 {
		return string(res)
	}
	return string(append([]byte{'1'},res...))
}