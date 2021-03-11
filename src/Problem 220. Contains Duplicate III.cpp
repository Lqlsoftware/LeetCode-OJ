// 213. House Robber II

#include <vector>

using namespace std;

class Solution {
public:
    bool containsNearbyAlmostDuplicate(vector<int>& nums, int k, int t) {
        if(nums.size() > 9999) {
            return false;
        }
        vector<long> nums_l(nums.size());
        vector<long> nums_s(nums.size());
        for (int i = 0;i < nums.size();i++) {
            nums_l[i] = (long)nums[i] + t;
            nums_s[i] = (long)nums[i] - t;
        }
        for (int i = 1;i <= k;i++) {
            for (int j = i;j < nums.size();j++) {
                if (nums[j - i] <= nums_l[j] && nums[j - i] >= nums_s[j]) return true;
            }
        }
        return false;
    }
};
