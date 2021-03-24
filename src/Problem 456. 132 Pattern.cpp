// 456. 132 Pattern

#include <bits/stdc++.h>
#include <vector>
#include <stack>
using namespace std;

// Enum 3 and get 2 by min stack
class Solution {
public:
    bool find132pattern(vector<int>& nums) {
        stack<int> min2;
        int N = nums.size(), last;
        vector<int> lmin(N, INT_MAX);
        for (int i = 1;i < N;i++) lmin[i] = min(lmin[i - 1], nums[i - 1]);

        min2.push(nums[N - 1]);
        for (int j = N - 2;j >= 1;j--) {
            last = INT_MIN;
            while (!min2.empty() && nums[j] > min2.top()) {
                last = min2.top();
                min2.pop();
            }
            if (last > lmin[j]) return true;
            min2.push(nums[j]);
        }
        return false;
    }
};

class Solution2 {
public:
    bool find132pattern(vector<int>& nums) {
        int N = nums.size();
        int numsi = nums[0];
        for (int j = 1;j < N - 1;j++) {
            for (int k = N - 1;k > j;k--) {
                if (numsi < nums[k] && nums[k] < nums[j]) return true;
            }
            if (nums[j] < numsi) numsi = nums[j];
        }
        return false;
    }
};