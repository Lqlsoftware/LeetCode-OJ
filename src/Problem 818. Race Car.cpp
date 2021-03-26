// 818. Race Car

#include <iostream>

using namespace std;

class Solution {
// speed 等比数列
// position = sum(speed[i]) + min(1 + racecar(sum(speed[i]) - target), 2 + racecar(target - sum(speed[i])))
private:
    int dp[10001];
public:
    int racecar(int target) {
        if (dp[target] > 0) return dp[target];

        // Exceed target
        int A = (int)(log2(target)) + 1;
        int pos = (1 << A) - 1;

        // Hit target
        if (pos == target) return dp[target] = A;

        // Reverse to get target
        int shortest = A + 1 + racecar(pos - target);

        // Try reversing before target
        int sub_target = target - (pos >> 1);
        for (int i = 0;i < A - 1;i++)
            shortest = min(shortest, A - 1 + 2 + i + racecar(sub_target + (1 << i) - 1));
        return dp[target] = shortest;
    }
};