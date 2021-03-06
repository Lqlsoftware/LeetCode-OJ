// 556. Next Greater Element III

#include <iostream>

using namespace std;

// 找到第一个非递增数字，与前序列的upper_bound交换，再对前序列进行排序
class Solution {
public:
    int nextGreaterElement(int n) {
        int nums[10];
        int digits = -1, i = 0;
        for (;n > 0;n /= 10) nums[++digits] = n % 10;
        
        // find peak
        for (i = 1;i <= digits;i++) if (nums[i - 1] > nums[i]) break;
        if (i == digits + 1) return -1;

        // make new peak
        int *find = upper_bound(nums, nums + i, nums[i]);
        int temp = nums[i];
        nums[i] = *find;
        *find = temp;
        sort(nums, nums + i, greater<int>());
    
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