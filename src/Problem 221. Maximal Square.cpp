// 221. Maximal Square

#include <vector>

#define min(a,b) (a < b ? a : b)
#define max(a,b) (a > b ? a : b)

using namespace std;


// dp[i][j] = min(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1
class Solution {
public:
    int maximalSquare(vector< vector<char> > & matrix) {
        int ret, m = matrix.size(), n = matrix[0].size();
        vector<int> dp(n, 0);
        for (int i = 0;i < m;i++) {
            int min_val = 0, new_val;
            for (int j = 0;j < n;j++) {
                new_val = matrix[i][j] == '1' ? min(min_val, dp[j]) + 1 : 0;
                min_val = min(dp[j], new_val);
                ret     = max(ret, new_val);
                dp[j]   = new_val;
            }
        }
        return ret * ret;
    }
};