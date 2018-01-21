/*
String to Integer (atoi)
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases.
If you want a challenge, please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs).
You are responsible to gather all the input requirements up front.
*/

package main

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	r := []rune(str)
	x,i := r[0],0
	for x == ' ' {
		i++
		x = r[i]
	}
	symbol := 1
	switch x {
	case '-':
		{
			symbol = -1
			x = '0'
		}
	case '+':
		{
			symbol = 1
			x = '0'
		}
	case ' ':
		{
			symbol = 1
			x = '0'
		}
	default:
		symbol = 1
	}
	res,tmp := int(x) - '0',0
	if res > 9 || tmp < 0 {
		return 0
	}
	for i = i + 1;i<len(r);i++ {
		tmp = int(r[i]) - '0'
		if tmp == ' ' {
			continue
		} else if tmp > 9 || tmp < 0 {
			break
		}
		res = res * 10 + tmp
		if res * symbol > 2147483647 || res * symbol < -2147483648 {
			break
		}
	}
	res *= symbol
	if res > 2147483647 {
		return 2147483647
	} else if res < -2147483648 {
		return -2147483648
	} else {
		return res
	}
}
