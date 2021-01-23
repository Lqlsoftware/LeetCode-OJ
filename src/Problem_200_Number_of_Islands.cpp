/*
200. Number of Islands
Given an m x n 2d grid map of '1's (land) and '0's (water), return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/
#include <vector>

using namespace std;

class Solution {
private:
    int m, n;
    void visit(vector<vector<char>>& grid, int i, int j) {
        if (grid[i][j] == '0') return;
        grid[i][j] = '0';
        if (i + 1 < m)  visit(grid, i + 1, j);
        if (j + 1 < n)  visit(grid, i, j + 1);
        if (j > 0)      visit(grid, i, j - 1);
        if (i > 0)      visit(grid, i - 1, j);
    }
public:
    int numIslands(vector<vector<char>>& grid) {
        m = grid.size();
        n = m > 0 ? grid[0].size() : 0;
        int ret = 0;
        for (int i = 0;i < m;i++) {
            for (int j = 0;j < n;j++) {
                if (grid[i][j] == '0') continue;
                visit(grid, i, j);
                ret++;
            }
        }
        return ret;
    }
};

class Solution2 {
private:
    int m, n;
    int g[300][300];
    void visit(vector<vector<char>>& grid, int i, int j) {
        if (g[i][j] == 1 || grid[i][j] == '0') return;
        g[i][j] = 1;
        if (i + 1 < m)  visit(grid, i + 1, j);
        if (j + 1 < n)  visit(grid, i, j + 1);
        if (j > 0)      visit(grid, i, j - 1);
        if (i > 0)      visit(grid, i - 1, j);
    }
public:
    int numIslands(vector<vector<char>>& grid) {
        m = grid.size();
        n = m > 0 ? grid[0].size() : 0;
        int ret = 0;
        for (int i = 0;i < m;i++) {
            for (int j = 0;j < n;j++) {
                if (g[i][j] == 1 || grid[i][j] == '0') continue;
                visit(grid, i, j);
                ret++;
            }
        }
        return ret;
    }
};