/*
Pascal's Triangle
Given numRows, generate the first numRows of Pascal's triangle.

For example, given numRows = 5,
Return
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

package main

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := range result {
		result[i] = make([]int, i + 1)
		result[i][0],result[i][i] = 1,1
		for j := 1;j < i;j++ {
			result[i][j] = result[i - 1][j - 1] + result[i - 1][j]
		}
	}
	return result
}