/*
Trapping Rain Water
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it is able to trap after raining.

For example,
Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
*/

package main

// use stack
func trap(height []int) int {
	stack,top := make([]int,100),-1
	water := 0
	for idx,v := range height {
		if top >= 0 && v >= height[stack[top]] {
			last := height[stack[top]]
			top--
			for top != -1 {
				x := stack[top]
				if height[x] <= v {
					water += (height[x] - last) * (idx - x - 1)
					top--
				} else {
					water += (v - last) * (idx - x - 1)
					break
				}
				last = height[x]
			}
		}
		top++
		stack[top] = idx
	}
	return water
}

// 记录两侧最高值
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left,right := make([]int,len(height)),make([]int,len(height))
	water := 0
	left[0] = height[0]
	right[len(height) - 1] = height[len(height) - 1]
	for idx := 1;idx < len(height);idx++ {
		left[idx] = max(left[idx - 1],height[idx])
	}
	for idx := len(height) - 2;idx >= 0;idx-- {
		right[idx] = max(right[idx + 1],height[idx])
	}
	for idx := 1;idx < len(height) - 1;idx++ {
		water += min(left[idx],right[idx]) - height[idx]
	}
	return water
}

func max(a,b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a <= b {
		return a
	}
	return b
}

// two pointer
func trap3(height []int) int {
	left,right := 0,len(height) - 1
	level,water := 0,0
	for left < right {
		var low int
		if height[left] < height[right] {
			low = height[left]
			left++
		} else {
			low = height[right]
			right--
		}
		level = max(level,low)
		water += level - low
	}
	return water
}