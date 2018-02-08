/*
Permutation Sequence
The set [1,2,3,…,n] contains a total of n! unique permutations.
By listing and labeling all of the permutations in order,
We get the following sequence (ie, for n = 3):
"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.
Note: Given n will be between 1 and 9 inclusive.
*/

package main

// 每个位置选取的数字与k和n!有关
func getPermutation(n int, k int) string {
	res,used,N := make([]byte,n),make([]bool,n),getFactorial(n - 1)
	k--
	n--
	for length := n; n >= 0;n-- {
		count := k / N[n]
		k %= N[n]
		// 查找可用的数字
		for i,v := range used {
			if v == false {
				if count == 0 {
					res[length - n],used[i] = byte(i + '1'),true
					break
				}
				count--
			}
		}
	}
	return string(res)
}

func getFactorial(n int) []int {
	res := make([]int,n + 1)
	res[0] = 1
	for i := 1;i <= n;i++ {
		res[i] = res[i - 1] * i
	}
	return res
}