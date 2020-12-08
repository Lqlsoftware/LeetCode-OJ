/*
Maximum Product Subarray
Given an integer array nums, find the contiguous subarray within an array
(containing at least one number) which has the largest product.

Example 1:
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

package main

// keep a local max & min and update result
func maxProduct(nums []int) int {
	currMin, currMax, res := nums[0], nums[0], nums[0]
	for i := 1;i < len(nums);i++ {
		v := nums[i]
		if currMin == 0 && currMax == 0 {
			currMin, currMax = v, v
		} else if nums[i] > 0 {
			currMin *= v
			currMax *= v
		} else {
			currMin, currMax = currMax * v, currMin * v
		}
		currMin = min(currMin, v)
		currMax = max(currMax, v)
		res = max(res, currMax)
	}
	return res
}

// my idiot solution: keep a local positive & negetive and update result
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pos, neg, res := 0, 0, math.MinInt32
	for _,v := range nums {
		if v == 0 {
			res = max(res, pos)
			pos, neg = 0, 0
		} else if v > 0 {
			if pos == 0 {
				pos = 1
			}
			pos *= v
			neg *= v
		} else {
			res = max(res, pos)
			// update neg & reset pos
			pos, neg = max(0, neg * v), min(v, pos * v)
		}
	}
	return max(res, pos)
}


func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
