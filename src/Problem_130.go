/*
Surrounded Regions
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

For example,
X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
*/

package main

var toward = [][]int{
	{0 , 1},
	{0 ,-1},
	{1 , 0},
	{-1, 0},
}

// dfs search for 'O' on the edge of matrix and mark as notfix.
func solve(board [][]byte) {
	m := len(board) - 1
	if m < 0 {
		return
	}
	n := len(board[0]) - 1
	for i := 0;i <= m;i++ {
		if board[i][0] == 'O' {
			dfsEdge(board, i, 0)
		}
		if board[i][n] == 'O' {
			dfsEdge(board, i, n)
		}
	}
	for i := 1;i < n;i++ {
		if board[0][i] == 'O' {
			dfsEdge(board, 0, i)
		}
		if board[m][i] == 'O' {
			dfsEdge(board, m, i)
		}
	}
	for i := range board {
		for j,v := range board[i] {
			if v == 'N' {
				board[i][j] = 'O'
			} else if v == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfsEdge(board [][]byte, i, j int) {
	board[i][j] = 'N'
	for k := range toward {
		x, y := i + toward[k][0],j + toward[k][1]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] == 'O' {
			dfsEdge(board, x, y)
		}
	}
}

// normal dfs to judge that if we can flop it into 'X'
func solve1(board [][]byte) {
	visit := make([][]bool,len(board))
	for i := range visit {
		visit[i] = make([]bool,len(board[0]))
	}
	for i := 1;i < len(board);i++ {
		for j := 1;j < len(board[0]);j++ {
			if board[i][j] == 'X' || visit[i][j] {
				continue
			}
			if dfs(board, visit, i, j) {
				fix(board, i, j)
			}
		}
	}
}

func dfs(board [][]byte, visit [][]bool, i, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	} else if board[i][j] == 'X' || visit[i][j] {
		return true
	}
	visit[i][j] = true
	res := true
	for k := range toward {
		res = dfs(board, visit, i + toward[k][0], j + toward[k][1]) && res
	}
	return res
}

func fix(board [][]byte, i, j int) {
	board[i][j] = 'X'
	for k := range toward {
		x,y := i + toward[k][0],j + toward[k][1]
		if board[x][y] == 'O' {
			fix(board, x, y)
		}
	}
}