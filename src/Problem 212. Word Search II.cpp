// 212. Word Search II

#include <vector>
#include <string>

using namespace std;

class Solution {
private:
    bool visit[16][16];
    int m;
    int n;
    bool dfs(vector<vector<char>>& board, int i, int j, string word, int idx) {
        if (board[i][j] != word[idx++]) return false;
        else if (idx == word.size())    return true;
        visit[i][j] = true;
        if (     i > 0     && visit[i - 1][j] == false && dfs(board, i - 1, j, word, idx)) return true;
        else if (i < m - 1 && visit[i + 1][j] == false && dfs(board, i + 1, j, word, idx)) return true;
        else if (j > 0     && visit[i][j - 1] == false && dfs(board, i, j - 1, word, idx)) return true;
        else if (j < n - 1 && visit[i][j + 1] == false && dfs(board, i, j + 1, word, idx)) return true;
        visit[i][j] = false;
        return false;
    }
    bool search(vector<vector<char>>& board, string word) {
        for (int i = 0;i < m;i++) {
            for (int j = 0;j < n;j++) {
                if (board[i][j] != word[0]) continue;
                memset(visit, 0, 256 * sizeof(bool));
                if (dfs(board, i, j, word, 0)) return true;
            }
        }
        return false;
    }
public:
    vector<string> findWords(vector<vector<char>>& board, vector<string>& words) {
        m = board.size(), n = board[0].size();
        vector<string> ret;
        for (string word : words)
            if (search(board, word)) ret.push_back(word);
        return ret;
    }
};