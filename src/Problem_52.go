/*
N-Queens II
Follow up for N-Queens problem.

Now, instead outputting board configurations, return the total number of distinct solutions.
*/

package main

func totalNQueens(n int) int {
	result := 0
	Queen(&result, make([]bool,n), make([]bool,n << 1), make([]bool,n << 1), 0)
	return result
}

func Queen(result *int, col, left, right []bool, cur int) {
	if cur == len(col) {
		*result++
		return
	}
	for i := 0;i < len(col);i++ {
		idxL,idxR := i + cur,i - cur + len(col)
		if col[i] == false && left[idxL] == false && right[idxR] == false {
			col[i], left[idxL], right[idxR] = true, true, true
			Queen(result, col, left, right, cur + 1)
			col[i], left[idxL], right[idxR] = false, false, false
		}
	}
}

// use bitmask
func totalNQueens1(n int) int {
	result := 0
	Queen1(&result, 1 << uint(n) - 1, 0, 0, 0)
	return result
}

func Queen1(result *int, full, col, left, right int) {
	if col == full {
		*result++
		return
	}
	next := ^(col | left | right)
	for next & full != 0 {
		bit := next & -next
		next -= bit
		Queen1(result, full, col | bit, (left | bit) >> 1, (right | bit) << 1)
	}
}