// 213. House Robber II

#include <vector>

using namespace std;
// Max (Rob (0 ~ n - 1), Rob (1 ~ n))
// So that the first and the last house will not be both robbed in optimal solution
class Solution {
private:
    int robI(vector<int>& nums, int start, int end) {
        int prev = 0, last = 0, temp;
        for (int i = start;i < end;i++) {
            temp = last;
            last = max(last, prev + nums[i]);
            prev = temp;
        }
        return last;
    }
public:
    int rob(vector<int>& nums) {
        int m = nums.size();
        if      (m == 0) return 0;
        else if (m == 1) return nums[0];
        return max(robI(nums, 0, m - 1), robI(nums, 1, m));
    }
};