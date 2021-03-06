// 556. Next Greater Element III

#include <iostream>

using namespace std;

class Solution {
private:
    int nums[10];
    bool nge(int digits) {
        if (digits <= 0) return false;
        if (nge(digits - 1)) return true;
        int curr = nums[digits];
        if (curr >= nums[digits - 1]) return false;
        int *find = upper_bound(nums, nums + digits, curr);
        // SWAP find & nums[digits]
        nums[digits] = *find;
        *find = curr;
        sort(nums, nums + digits, greater<int>());
        return true;
    }
public:
    int nextGreaterElement(int n) {
        int digits = -1;
        for (int curr = n;curr > 0;curr /= 10) nums[++digits] = curr % 10;
        if (!nge(digits)) return -1;
        
        uint64_t ret = 0;
        while (digits >= 0) {
            ret *= 10;
            ret += nums[digits];
            digits--;
        }
        
        if (ret > INT_MAX) return -1;
        return (int)ret;
    }
};