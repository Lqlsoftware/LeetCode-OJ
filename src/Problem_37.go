/*
Sudoku Solver

Write a program to solve a Sudoku puzzle by filling the empty cells.
Empty cells are indicated by the character '.'.
You may assume that there will be only one unique solution.
*/

package main

// 位图检测+回溯法
func solveSudoku(board [][]byte)  {
	if len(board) != 9 || len(board[0]) != 9 {
		return
	}
	col := make([]uint16,9)
	row := make([]uint16,9)
	square := make([]uint16,9)
	empty := make([]uint16,9)
	// 初始化冲突位图
	for i,v := range board {
		for j,val := range v {
			if val == '.' {
				empty[i] |= 1 << uint(j)
				continue
			}
			idx := uint16(1 << (val - '0'))
			if col[j] & idx != uint16(0) || row[i] & idx != uint16(0) || square[i / 3 * 3 + j / 3] & idx != uint16(0) {
				return
			}
			col[j] |= idx
			row[i] |= idx
			square[i / 3 * 3 + j / 3] |= idx
		}
	}
	solve(board,empty,col,row,square,0)
}

// 回溯法递归解决数独问题
func solve(board [][]byte, empty, col, row, square []uint16, cur int) bool {
	for i := cur;i < 81;i ++ {
		r,c := i / 9,i % 9
		if empty[r] & (1 << uint(c)) == uint16(0) {
			continue
		}
		s := r / 3 * 3 + c / 3
		for v := 1;v <= 9;v++ {
			idx := uint16(1 << uint(v))
			if col[c] & idx == uint16(0) && row[r] & idx == uint16(0) && square[s] & idx == uint16(0) {
				col[c] |= idx
				row[r] |= idx
				square[s] |= idx
				board[r][c] = byte('0' + v)
				if solve(board, empty, col, row, square, i+1) {
					return true
				}
				col[c] ^= idx
				row[r] ^= idx
				square[s] ^= idx
			}
		}
		return false
	}
	return true
}

func sudokuPrint(x [][]byte) {
	for i,v := range x {
		if i % 3 == 0 {
			fmt.Println(" -----------------")
		}
		for j, val := range v {
			if j % 3 == 0 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(string(val))
		}
		fmt.Println("|")
	}
	fmt.Println(" -----------------")
}