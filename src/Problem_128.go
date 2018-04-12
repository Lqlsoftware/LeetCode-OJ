/*
Longest Consecutive Sequence
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

For example,
Given [100, 4, 200, 1, 3, 2],
The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.

Your algorithm should run in O(n) complexity.
*/

package main

// use hashmap bucket
var exist = struct{}{}
func longestConsecutive(nums []int) int {
	bucket := make(map[int]struct{})
	// get numbers into bucket
	for _,v := range nums {
		bucket[v] = exist
	}
	res := 0
	// For every numbers in bucket, visit its neighbor(num - 1, num + 1).
	// Remove every number once you visit it.
	// So the total time complexity is O(n) as we visit all numbers through bucket.
	for v := range bucket {
		delete(bucket, v)
		count := visit(bucket, v - 1, -1) + visit(bucket, v + 1, 1) + 1
		if count > res {
			res = count
		}
	}
	return res
}

func visit(bucket map[int]struct{}, current, toward int) int {
	if _,ok := bucket[current];!ok {
		return 0
	}
	delete(bucket, current)
	return visit(bucket, current + toward, toward) + 1
}