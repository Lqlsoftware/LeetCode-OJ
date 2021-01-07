/*
164. Maximum Gap
Given an unsorted array, find the maximum difference between the successive elements in its sorted form.
Return 0 if the array contains less than 2 elements.

Example 1:
Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.

Example 2:
Input: [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.

Note:
You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
Try to solve it in linear time/space.
*/
#include <vector>
using namespace std;

/* Make sure that there will be a empty block (n nums into n + 1 blocks)
 */
class Solution {
public:
    int maximumGap(vector<int>& nums) {
        if (nums.size() < 2) return 0;
        // Min and Max elements
        int min = 0, max = 0;
        for (int i = 0;i < nums.size();i++) {
            if (nums[i] < nums[min]) min = i;
            else if (nums[i] > nums[max]) max = i;
        }
        
        // Split into nums.size() - 1 blocks
        int block_size = (nums[max] - nums[min] + nums.size()) / (nums.size() + 1);
        if (block_size == 0) return 0;
        int block_num  = nums.size() + 2;
        int* block_max = (int *)malloc(block_num * sizeof(int));
        int* block_min = (int *)malloc(block_num * sizeof(int));
        
        // Init blocks with INT_MIN and INT_MAX
        for (int i = 0;i < block_num;i++) {
            block_max[i] = INT_MIN;
            block_min[i] = INT_MAX;
        }
        
        // Put nums into blocks
        for (int i = 0;i < nums.size();i++) {
            int b = (nums[i] - nums[min]) / block_size;
            if (nums[i] < block_min[b]) block_min[b] = nums[i];
            if (nums[i] > block_max[b]) block_max[b] = nums[i];
        }
        
        // Find the max difference
        int max_diff = 0;
        int prev = nums[min], diff;
        for (int i = 0;i < block_num;i++) {
            if (block_max[i] == INT_MIN) continue;
            diff = block_min[i] - prev;
            prev = block_max[i];
            if (diff > max_diff) max_diff = diff;
        }
        return max_diff;
    }
};