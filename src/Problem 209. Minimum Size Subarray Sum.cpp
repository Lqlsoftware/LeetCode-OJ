//209. Minimum Size Subarray Sum

// O(NlogN)
class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int m = nums.size() + 1;
        // Sum of prefix
        int dp[m]; dp[0] = 0;
        for (int i = 1;i < m;i++) dp[i] = dp[i - 1] + nums[i - 1];
        
        // Binary search for the largest element less than dp[i][0]
        int ret = INT_MAX;
        for (int i = 1;i < m;i++) {
            if (dp[i] < s) continue;
            auto bound = upper_bound(dp, dp + m, dp[i] - s);
            if (bound != dp) ret = min(ret, i - int(bound - dp) + 1);
        }
        return ret == INT_MAX ? 0 : ret;
    }
};

// O(N)
class Solution2 {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int ret = INT_MAX, l = 0, r = 0, sum = 0;
        while (r < nums.size()) {
            sum += nums[r++];
            if (sum >= s) {
                while (sum >= s) sum -= nums[l++];
                ret = min(ret, r - l + 1);
            }
        }
        return ret == INT_MAX ? 0 : ret;
    }
};