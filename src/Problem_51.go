/*
N-Queens
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle.
Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

For example,
There exist two distinct solutions to the 4-queens puzzle:
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
*/

package main

func solveNQueens(n int) [][]string {
	var result [][]string
	Queen(&result, make([]int,n), 0)
	return result
}

func Queen(result *[][]string, current []int, cur int) {
	n := len(current)
	if cur == len(current) {
		var temp []string
		for r := 0; r < n; r++ {
			row := make([]byte, n)
			for c := 0; c < n; c++ {
				if current[r] == c {
					row[c] = 'Q'
				} else {
					row[c] = '.'
				}
			}
			temp = append(temp, string(row))
		}
		*result = append(*result,temp)
		return
	}
	for i,j := 0,0;i < n;i++ {
		current[cur] = i
		for j = 0;j < cur;j++ {
			if current[j] == i || abs(i - current[j]) + j == cur {
				break
			}
		}
		if j == cur {
			Queen(result, current, cur + 1)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}