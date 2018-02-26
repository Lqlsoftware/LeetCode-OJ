/*
Largest Rectangle in Histogram
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
find the area of largest rectangle in the histogram.

Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
The largest rectangle is shown in the shaded area, which has area = 10 unit.

For example,
Given heights = [2,1,5,6,2,3],
return 10.
*/

package main

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	max,end := 0,0
	stack,top := make([]int,len(heights)),-1
	for end < len(heights) {
		if top == -1 || heights[stack[top]] <= heights[end] {
			top++
			stack[top] = end
			end++
		} else {
			min,length := heights[stack[top]],0
			top--
			if top == -1 {
				length = end
			} else {
				length = end - stack[top] - 1
			}
			if x := min * length;x > max {
				max = x
			}
		}
	}
	return max
}