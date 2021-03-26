// 931. Minimum Falling Path Sum

#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int minFallingPathSum(vector<vector<int> >& matrix) {
        int n = matrix.size();
        vector<int> dp(n, 0);
        for (int i = n - 1;i >= 0;i--) {
            int old_val = dp[0];
            for (int j = 0;j < n;j++) {
                int temp = min(old_val, dp[j]);
                old_val = dp[j];
                if (j + 1 < n) temp = min(temp, dp[j + 1]);
                dp[j] = temp + matrix[i][j];
            }
        }
        int ret = 0;
        for (int i = 0;i < n;i++) if (dp[i] < dp[ret]) ret = i;
        return dp[ret];
    }
};

class Solution2 {
public:
    int minFallingPathSum(vector<vector<int> >& matrix) {
        int n = matrix.size();
        int dp[n][n];

        for (int i = 0;i < n;i++) dp[n - 1][i] = matrix[n - 1][i];

        for (int i = n - 2;i >= 0;i--) {
            for (int j = 0;j < n;j++) {
                dp[i][j] = dp[i + 1][j];
                if (j - 1 >= 0) dp[i][j] = min(dp[i][j], dp[i + 1][j - 1]);
                if (j + 1 < n) dp[i][j] = min(dp[i][j], dp[i + 1][j + 1]);
                dp[i][j] += matrix[i][j];
            }
        }
        int ret = 0;
        for (int i = 0;i < n;i++) if (dp[0][i] < dp[0][ret]) ret = i;
        return dp[0][ret];
    }
};