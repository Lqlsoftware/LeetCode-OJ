/*
Word Search
Given a 2D board and a word, find if the word exists in the grid.
The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring.
The same letter cell may not be used more than once.

For example,
Given board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
word = "ABCCED", -> returns true,
word = "SEE", -> returns true,
word = "ABCB", -> returns false.
*/

package main

// DFS 可以通过对原数组取反(byte(uint8) ^ 255)代替visit(不建议使用修改输入的方式来减少内存空间占用)
func exist(board [][]byte, word string) bool {
	// 				   left    up   right  down
	toward := [][]int{{0,-1},{-1,0},{0,1},{1,0}}
	for y,v := range board {
		for x,val := range v {
			if val == word[0] && solve(board,toward,y,x,1,word) {
				return true
			}
		}
	}
	return false
}

func solve(board [][]byte, toward [][]int, y, x int, current int, word string) bool {
	if current == len(word) {
		return true
	}
	m,n := len(board),len(board[0])
	board[y][x] ^= 255
	for i := 0;i < 4;i++ {
		posY,posX := y + toward[i][0],x + toward[i][1]
		if posY < 0 || posY == m || posX < 0 || posX == n || board[posY][posX] != word[current] {
			continue
		}
		if solve(board,toward,posY,posX,current + 1,word) {
			board[y][x] ^= 255
			return true
		}
	}
	board[y][x] ^= 255
	return false
}