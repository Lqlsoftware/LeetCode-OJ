/*
Valid Sudoku

Determine if a Sudoku is valid, according to: Sudoku Puzzles - The Rules.
The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Note:
A valid Sudoku board (partially filled) is not necessarily solvable.
Only the filled cells need to be validated.
*/

package main

func isValidSudokuBit(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	col := make([]uint16,9)
	row := make([]uint16,9)
	square := make([]uint16,9)
	for i,v := range board {
		for j,val := range v {
			if val == '.' {
				continue
			}
			idx := uint16(1 << (val - '0'))
			if col[j] & idx != uint16(0) || row[i] & idx != uint16(0) || square[j / 3 * 3 + i / 3] & idx != uint16(0) {
				return false
			}
			col[j] |= idx
			row[i] |= idx
			square[j / 3 * 3 + i / 3] |= idx
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	// row test
	for _,col := range board {
		dict := make(map[byte]int,0)
		for _,row := range col {
			if row == '.' {
				continue
			} else if _,ok := dict[row];ok == true {
				return false
			}
			dict[row] = 1
		}
	}
	// col test
	for i := 0;i < 9;i++ {
		dict := make(map[byte]int,0)
		for j := 0;j < 9;j++ {
			if board[j][i] == '.' {
				continue
			} else if _,ok := dict[board[j][i]];ok == true {
				return false
			}
			dict[board[j][i]] = 1
		}
	}
	// square test
	for i := 0;i < 9;i++ {
		dict := make(map[byte]int,0)
		startX,startY := (i / 3) * 3,(i % 3) * 3
		for j := 0;j < 3;j++ {
			for k := 0;k < 3;k++ {
				x := board[startX + j][startY + k]
				if x == '.' {
					continue
				} else if _,ok := dict[x];ok == true {
					return false
				}
				dict[x] = 1
			}
		}
	}
	return true
}