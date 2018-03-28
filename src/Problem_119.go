/*
Pascal's Triangle II
Given an index k, return the kth row of the Pascal's triangle.

For example, given k = 3,
Return [1,3,3,1].

Note:
Could you optimize your algorithm to use only O(k) extra space?
*/

package main

// math solution
func getRow(rowIndex int) []int {
	result := make([]int,rowIndex + 1)
	result[0],result[rowIndex] = 1,1
	for current,front,back := 1,1,rowIndex;front <= rowIndex / 2;front,back = front + 1,back - 1 {
		current = current * back / front
		result[front],result[back - 1] = current,current
	}
	return result
}

// dp solution
func getRow1(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}
	result := make([]int, rowIndex + 1)
	result[0] = 1
	for i := 0;i <= rowIndex;i++ {
		result[i] = 1
		for j := i - 1;j > 0;j-- {
			result[j] += result[j - 1]
		}
	}
	return result
}