/*
Roman to Integer
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.
*/

package main

// new method 去掉了乘法操作
func romanToInt(s string) int {
	roman := []rune(s)
	dict := make(map[rune]int, 7)
	dict['M'] = 1000
	dict['D'] = 500
	dict['C'] = 100
	dict['L'] = 50
	dict['X'] = 10
	dict['V'] = 5
	dict['I'] = 1

	res := 0
	for i, v := range roman {
		next := 0
		if i < len(roman)-1 {
			next = dict[roman[i+1]]
		}
		if dict[v] < next {
			res -= dict[v]
		} else {
			res += dict[v]
		}
	}
	return res
}

// old method
func romanToInt(s string) int {
	roman := []rune(s)
	dict := make(map[rune]int, 7)
	dict['M'] = 1000
	dict['D'] = 500
	dict['C'] = 100
	dict['L'] = 50
	dict['X'] = 10
	dict['V'] = 5
	dict['I'] = 1
	dict['0'] = 0

	target := '0'
	count := 0
	res := 0
	for i := range roman {
		if roman[i] != target {
			if dict[target] < dict[roman[i]] {
				res -= count * dict[target]
			} else {
				res += count * dict[target]
			}
			count = 1
			target = roman[i]
		} else {
			count++
		}
	}
	res += count * dict[target]

	return res
}
