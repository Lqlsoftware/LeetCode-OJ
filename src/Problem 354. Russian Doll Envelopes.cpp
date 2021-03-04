// 354. Russian Doll Envelopes

#include <string>

using namespace std;

/*
    LIS Problem
*/

class Solution {
private:
    int dp[5001];
public:
    int maxEnvelopes(vector<vector<int>>& envelopes) {
        sort(envelopes.begin(), envelopes.end(), [](vector<int>& a, vector<int>& b) {
            return a[0] < b[0] || a[0] == b[0] && a[1] > b[1];
        });

        int maxE = 0;
        for (auto e : envelopes) {
            int *idx = lower_bound(dp, dp + maxE, e[1]);
            if (idx == dp + maxE) maxE++;
            *idx = e[1];
        }
        
        return maxE;
    }
};
