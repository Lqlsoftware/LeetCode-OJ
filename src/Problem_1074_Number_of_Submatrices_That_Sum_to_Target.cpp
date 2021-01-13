/*
1074. Number of Submatrices That Sum to Target
Given a matrix and a target, return the number of non-empty submatrices that sum to target.
A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 <= x <= x2 and y1 <= y <= y2.
Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some coordinate that is different: for example, if x1 != x1'.

Example 1:
Input: matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
Output: 4
Explanation: The four 1x1 submatrices that only contain 0.

Example 2:
Input: matrix = [[1,-1],[-1,1]], target = 0
Output: 5
Explanation: The two 1x2 submatrices, plus the two 2x1 submatrices, plus the 2x2 submatrix.

Example 3:
Input: matrix = [[904]], target = 0
Output: 0

Constraints:
1 <= matrix.length <= 100
1 <= matrix[0].length <= 100
-1000 <= matrix[i] <= 1000
-10^8 <= target <= 10^8
*/

// Special version of hashtable
#define HT_LENGTH 256
struct HashTable {
    int keys[HT_LENGTH];
    int vals[HT_LENGTH];
    
    HashTable() { this->clear();}
    
    void clear() {
        keys[0] = 0; vals[0] = 1;
        for (int i = 1;i < HT_LENGTH;i++) keys[i] = -1;
    }
    
    int find(int key) {
        uint32_t ikey = (uint32_t)key % HT_LENGTH;
        for (int i = 0;i < HT_LENGTH;i++) {
            if (keys[ikey] == -1)  return 0;
            if (keys[ikey] == key) return vals[ikey];
            ikey = ++ikey == HT_LENGTH ? 0 : ikey;
        }
        return false;
    }
    
    void inc(int key) {
        uint32_t ikey = (uint32_t)key % HT_LENGTH;
        for (int i = 0;i < HT_LENGTH;i++) {
            if (keys[ikey] == -1) {
                keys[ikey] = key;
                vals[ikey] = 1;
                return;
            }
            if (keys[ikey] == key) {
                vals[ikey]++;
                return;
            }
            ikey = ++ikey == HT_LENGTH ? 0 : ikey;
        }
    }
};

// 36ms solution using 2D prefix of sum
class Solution {
private:
    HashTable ht;
public:
    int numSubmatrixSumTarget(vector<vector<int>>& matrix, int target) {
        int rprefix[100][101];
        int m = matrix.size(), n = matrix[0].size();
        int ret = 0;
        for (int i = 0;i < m;i++) {
            rprefix[i][0] = 0;
            for (int j = 1;j <= n;j ++) {
                rprefix[i][j] = rprefix[i][j - 1] + matrix[i][j - 1];
            }
        }
        
        for (int i = 1;i <= n;i++) {
            for (int j = i;j <= n;j++) {
                ht.clear();
                int cprefix = 0;
                for (int row = 0;row < m;row++) {
                    cprefix += rprefix[row][j] - rprefix[row][i - 1];
                    ret += ht.find(cprefix - target);
                    ht.inc(cprefix);
                }
            }
        }
        return ret;
    }
};