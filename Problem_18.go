/*
4Sum
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target?
Find all unique quadruplets in the array which gives the sum of target.
Note: The solution set must not contain duplicate quadruplets.

For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	lenth := len(nums)
	for idx1 := 0; idx1 < lenth-3; idx1++ {
		for idx2 := idx1 + 1; idx2 < lenth-2; idx2++ {
			front, back := idx2+1, lenth-1
			v1, v2, v3, v4 := nums[idx1], nums[idx2], nums[front], nums[back]
			tar := v1 + v2 - target
			for current := nums[front] + nums[back] + tar; front < back; current = nums[front] + nums[back] + tar {
				v3, v4 = nums[front], nums[back]
				if current < 0 {
					front++
				} else if current > 0 {
					back--
				} else {
					res = append(res, []int{v1, v2, v3, v4})
					for front < back && nums[front] == v3 {
						front++
					}
					for front < back && nums[back] == v4 {
						back--
					}
				}
			}
			for idx2 < lenth-2 && nums[idx2+1] == nums[idx2] {
				idx2++
			}
		}
		for idx1 < lenth-3 && nums[idx1+1] == nums[idx1] {
			idx1++
		}
	}
	return res
}
