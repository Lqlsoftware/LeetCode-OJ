/*
Combinations
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

For example,
If n = 4 and k = 2, a solution is:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

package main

// all my solutions beat 100%
// Backtracking
func combine(n int, k int) [][]int {
	var result [][]int
	solve(&result,[]int{},n,k,1)
	return result
}

func solve(result *[][]int, current []int, n, k, start int) {
	if len(current) == k {
		temp := make([]int,k)
		copy(temp,current)
		*result = append(*result, temp)
		return
	}
	for i := start;i <= n;i++ {
		solve(result, append(current,i),n,k,i + 1)
	}
}

// generate array (res[i] < res[i + 1])
func combine1(n int, k int) [][]int {
	var result [][]int
	res,i := make([]int,k),0
	for i >= 0 {
		if res[i]++;res[i] > n {
			i--
		} else if i == k - 1 {
			temp := make([]int,k)
			copy(temp,res)
			result = append(result,temp)
		} else {
			i++
			res[i] = res[i - 1]
		}
	}
	return result
}

// recursive C(n, k) = n * C(n - 1, k - 1) + C(n - 1, k)
func combine2(n int, k int) [][]int {
	if k == 0 {
		return [][]int{{}}
	} else if n == k {
		res := make([]int,k)
		for i := 1;i <= n;i++ {
			res[i - 1] = i
		}
		return [][]int{res}
	}
	result := combine2(n - 1,k - 1)
	for i := range result {
		result[i] = append(result[i],n)
	}
	result = append(result, combine2(n - 1, k)...)
	return result
}