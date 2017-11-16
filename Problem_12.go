/*
Integer to Roman
Given an integer, convert it to a roman numeral.

Input is guaranteed to be within the range from 1 to 3999.
*/

package main

func intToRoman(num int) string {
	res := make([]rune, 0)

	m := num / 1000
	for i := 0; i < m; i++ {
		res = append(res, 'M')
	}

	num %= 1000
	if num >= 500 {
		if num/100 == 9 {
			res = append(res, 'C')
			res = append(res, 'M')
			num -= 900
		} else {
			res = append(res, 'D')
			num -= 500
		}
	} else if num/100 == 4 {
		res = append(res, 'C')
		res = append(res, 'D')
		num -= 400
	}

	c := num / 100
	for i := 0; i < c; i++ {
		res = append(res, 'C')
	}

	num %= 100
	if num >= 50 {
		if num/10 == 9 {
			res = append(res, 'X')
			res = append(res, 'C')
			num -= 90
		} else {
			res = append(res, 'L')
			num -= 50
		}
	} else if num/10 == 4 {
		res = append(res, 'X')
		res = append(res, 'L')
		num -= 40
	}

	x := num / 10
	for i := 0; i < x; i++ {
		res = append(res, 'X')
	}

	num %= 10
	if num >= 5 {
		if num == 9 {
			res = append(res, 'I')
			res = append(res, 'X')
			num -= 9
		} else {
			res = append(res, 'V')
			num -= 5
		}
	} else if num == 4 {
		res = append(res, 'I')
		res = append(res, 'V')
		num -= 4
	}

	I := num
	for i := 0; i < I; i++ {
		res = append(res, 'I')
	}
	return string(res)
}
