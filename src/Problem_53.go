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

// divide and conquer approach 将数组一分为二分治处理
func maxSubArray1(nums []int) int {
	_,_,res,_ := dac(nums)
	return res
}

func dac(nums []int) (Start,End,Max,Sum int){
	if len(nums) == 1 {
		return nums[0],nums[0],nums[0],nums[0]
	} else {
		start1,end1,max1,sum1 := dac(nums[:len(nums) / 2])
		start2,end2,max2,sum2 := dac(nums[len(nums) / 2:])
		return max(start1,sum1 + start2),max(end2,end1 + sum2),max(max(max1,max2),end1 + start2),sum1 + sum2
	}
}