/*
Merge Sorted Array
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
The number of elements initialized in nums1 and nums2 are m and n respectively.
*/

package main

// 从后往前将大的数置于最后
func merge(nums1 []int, m int, nums2 []int, n int)  {
	i,j,target := m - 1,n - 1,m + n - 1
	for j >= 0 {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[target] = nums1[i]
			i--
		} else {
			nums1[target] = nums2[j]
			j--
		}
		target--
	}
}