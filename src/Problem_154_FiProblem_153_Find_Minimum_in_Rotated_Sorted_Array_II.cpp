/*
154. Find Minimum in Rotated Sorted Array II
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

The array may contain duplicates.

Example 1:
Input: [1,3,5]
Output: 1

Example 2:
Input: [2,2,2,0,1]
Output: 0

Note:
This is a follow up problem to Find Minimum in Rotated Sorted Array.
Would allow duplicates affect the run-time complexity? How and why?
*/

#include <vector>
using namespace std;

// Divide and conquer
class Solution {
public:
    int findMin(vector<int>& nums) {
        int left = 0, right = nums.size() - 1;
        while(left < right) {
            int mid = (left + right) / 2;
            // there is small value in left ~ mid
            if (nums[mid] < nums[right]) right = mid;
            // there is small value in mid + 1 ~ right
            else if (nums[mid] > nums[right]) left = mid + 1;
            // mid == right, abandon right
            else right--;
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
            // left == right == mid, trim it
            else if (nums[left] == nums[right] && nums[left] == nums[mid]) {
                left ++;
                right --;
            }
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