/*
Palindrome Number
Determine whether an integer is a palindrome. Do this without extra space.
*/

package main

func isPalindrome(x int) bool {
	if x < 0 || x % 10 == 0 && x != 0 {
		return false
	}
	j := 0
	for x > j {
		j = j * 10 + x % 10
		x /= 10
	}
	return x == j || x == j / 10
}
