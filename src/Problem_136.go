/*
Single Number
Given a non-empty array of integers, every element appears twice except for one. Find that single one.
Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
Input: [2,2,1]
Output: 1

Example 2:
Input: [4,1,2,1,2]
Output: 4
*/

package main

// eg: a,a,b,b,c  res = a ^ a ^ b ^ b ^ c
// 						a ^ a = 0
//						b ^ b = 0
//					  = c
func singleNumber(nums []int) int {
	res := 0
	for _,v := range nums {
		res ^= v
	}
	return res
}