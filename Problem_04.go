/*
Median of Two Sorted Arrays
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5
*/

package main

import "golang.org/x/tools/container/intsets"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		nums1, nums2 = nums2, nums1
		len1, len2 = len2, len1
	}
	var c1, c2, L1, L2, R1, R2, low, high int
	low, high = 0, 2*len1
	for low <= high {
		c1 = (low + high) / 2
		c2 = len1 + len2 - c1
		if c1 == 0 {
			L1 = intsets.MinInt
		} else {
			L1 = nums1[(c1-1)/2]
		}
		if c1 == 2*len1 {
			R1 = intsets.MaxInt
		} else {
			R1 = nums1[c1/2]
		}
		if c2 == 0 {
			L2 = intsets.MinInt
		} else {
			L2 = nums2[(c2-1)/2]
		}
		if c2 == 2*len2 {
			R2 = intsets.MaxInt
		} else {
			R2 = nums2[c2/2]
		}

		if L1 > R2 {
			high = c1 - 1
		} else if L2 > R1 {
			low = c1 + 1
		} else {
			break
		}
	}
	if L1 < L2 {
		L1, L2 = L2, L1
	}
	if R1 > R2 {
		R1, R2 = R2, R1
	}
	return float64(L1+R1) / 2
}

/*
func find(nums1 []int, nums2 []int, range1 []int, range2 []int, left int, right int) float64 {
	target := (range2[0] + range2[1]) / 2
	var i int
	for i = range1[0]; i < range1[1]; i++ {
		if nums2[target] < nums1[i] {
			left += (range2[1] - range2[0]) / 2
			right += (range2[1] - range2[0]) / 2
			if left > right + 1 {
				return find(nums1,nums2,[]int{0,0},[]int{range2[0],target},left,right)
			} else if left < right - 1 {
				return find(nums1,nums2,[]int{range1[0],range1[1]},[]int{target,range2[0]},left,right)
			} else if left == right {
				return float64(nums2[target])
			} else {
				if left == right - 1 { return float64(nums2[target] + nums2[target - 1])}
			}
		} else if nums1[i] <= nums2[target] && nums1[i+1] > nums2[target] {
			left += (range2[1] - range2[0]) / 2 + i + 1
			right += (range2[1] - range2[0]) / 2 + range1[1] - range1[0] - i - 1
			break
		}
	}
	if nums2[target] > nums1[i] {
		left += range1[1] - range1[0] + (range2[1] - range2[0]) / 2
		right += (range2[1] - range2[0]) / 2
	}
	if left > right + 1 {
		range1[0],range1[1] = range1[0] + i,range1[1]
		return find(nums1,nums2,[]int{})
	}
}
*/
