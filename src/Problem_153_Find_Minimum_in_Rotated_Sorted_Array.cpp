/*
153. Find Minimum in Rotated Sorted Array
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums, return the minimum element of this array.

Example 1:
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.

Example 2:
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

Example 3:
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times. 
 
Constraints:
n == nums.length
1 <= n <= 5000
-5000 <= nums[i] <= 5000
All the integers of nums are unique.
nums is sorted and rotated between 1 and n times.
*/
#include <vector>
using namespace std;

// Divide and conquer
class Solution {
public:
    int findMin(vector<int>& nums) {
        int left = 0, right = nums.size() - 1;
        int ref = nums[right];
        while(left < right) {
            int mid = (left + right) / 2;
            if (nums[mid] > ref) left = mid + 1;
            else right = mid;
        }
        return nums[left];
    }
};

class Solution1 {
public:
    int findMin(vector<int>& nums) {
        int left = 0, right = nums.size() - 1, mid;
        while(left < right) {
            mid = (left + right) / 2;
            // Pivot in seq: left ~ mid 
            if (nums[left] > nums[mid]) right = mid;
            // Pivot in seq: mid + 1 ~ right
            else if (nums[mid + 1] > nums[right]) left = mid + 1;
            // Pivot on left or mid + 1
            else return min(nums[left], nums[mid + 1]);
        }
        return nums[left];
    }
private:
    int min(int a, int b) {
        return a > b ? b : a;
    }
};