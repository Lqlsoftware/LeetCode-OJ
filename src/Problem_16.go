/*
3Sum Closest
Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target.
Return the sum of the three integers.
You may assume that each input would have exactly one solution.

For example, given array S = {-1 2 1 -4}, and target = 1.
The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/

package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	length, min, sum := len(nums)-1, math.MaxInt32, 0
	for i, v := range nums {
		front, back := i+1, length
		for front < back {
			v1, v2 := nums[front], nums[back]
			cal := v1 + v2 + v
			calcu := cal - target
			if calcu < 0 {
				calcu = -calcu
			}
			if calcu < min {
				min = calcu
				sum = v1 + v2 + v
			}
			if cal < target {
				front++
			} else if cal == target {
				return cal
			} else {
				back--
			}
		}
	}
	return sum
}
