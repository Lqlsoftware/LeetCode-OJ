/*
Plus One
Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.
You may assume the integer do not contain any leading zero, except the number 0 itself.
The digits are stored such that the most significant digit is at the head of the list.
*/

package main

func plusOne(digits []int) []int {
	length := len(digits) - 1
	digits[length]++
	for i := length;i > 0 && digits[i] > 9;i-- {
		digits[i - 1] += digits[i] / 10
		digits[i] %= 10
	}
	if digits[0] == 10 {
		return append([]int{1,0},digits[1:]...)
	}
	return digits
}