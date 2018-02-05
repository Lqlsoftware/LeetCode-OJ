/*
Maximum Subarray
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.
click to show more practice.
More practice:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

package main

func maxSubArray(nums []int) int {
	preMax,curMax := nums[0],nums[0]
	for i := 1; i < len(nums); i++ {
		preMax = max(preMax + nums[i], nums[i])
		curMax = max(curMax, preMax)
	}
	return curMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}