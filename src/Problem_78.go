/*
Subsets
Given a set of distinct integers, nums, return all possible subsets (the power set).
Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,3], a solution is:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

package main

// recursive C(num) = n * C(num[1:]) + C(num[1:])
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	// 得到不包含第一个数字的子集
	result := subsets(nums[1:])
	length := len(result)
	for i := 0;i < length;i++ {
		// 把得到的子集复制入结果集
		temp := make([]int,len(result[i]))
		copy(temp,result[i])
		result = append(result, temp)
		// 向得到的子集中添加第一个数字构成新子集
		result[i] = append(result[i],nums[0])
	}
	return result
}

/*
	Number of subsets for {1 , 2 , 3 } = 2^3.
	case    possible outcomes for the set of subsets
	  1   ->          Take or dont take = 2
	  2   ->          Take or dont take = 2
	  3   ->          Take or dont take = 2

	Take = 1
	Dont take = 0

	0) 0 0 0  -> Dont take 3 , Dont take 2 , Dont take 1 = { }
	1) 0 0 1  -> Dont take 3 , Dont take 2 ,   take 1    = {1 }
	2) 0 1 0  -> Dont take 3 ,    take 2   , Dont take 1 = { 2 }
	3) 0 1 1  -> Dont take 3 ,    take 2   ,   take 1    = { 1 , 2 }
	4) 1 0 0  ->    take 3   , Dont take 2 , Dont take 1 = { 3 }
	5) 1 0 1  ->    take 3   , Dont take 2 ,   take 1    = { 1 , 3 }
	6) 1 1 0  ->    take 3   ,    take 2   , Dont take 1 = { 2 , 3 }
	7) 1 1 1  ->    take 3   ,    take 2   ,   take 1    = { 1 , 2 , 3 }

	Time complexity : O(n*2^n) , for every input element loop traverses the whole solution set length i.e. 2^n
 */
func subsets(nums []int) [][]int {
    total := 1 << uint(len(nums))
	result := make([][]int,total)
	for i := range nums {
		for j := 0;j < total;j++ {
			if j >> uint(i) & 1 == 1 {
				result[j] = append(result[j],nums[i])
			}
		}
	}
	return result
}