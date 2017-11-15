/*
Container With Most Water
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container,
such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	s := 0
	for left < right {
		var t int
		if height[left] < height[right] {
			t = height[left] * (right - left)
			left++
		} else {
			t = height[right] * (right - left)
			right--
		}
		if t > s {
			s = t
		}
	}
	return s
}
