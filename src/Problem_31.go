/*
Next Permutation

Implement next permutation,
which rearranges numbers into the lexicographically next greater permutation of numbers.
If such arrangement is not possible,
it must rearrange it as the lowest possible order (ie, sorted in ascending order).
The replacement must be in-place, do not allocate extra memory.

Here are some examples.
Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

package main

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i:=len(nums) - 2;i >= 0;i-- {
		if nums[i] < nums[i + 1] {
			sort.Ints(nums[i + 1:])
			for j:=i + 1;j<len(nums);j++ {
				if nums[j] > nums[i] {
					nums[j],nums[i] = nums[i],nums[j]
					return
				}
			}
		}
	}
	sort.Ints(nums)
}