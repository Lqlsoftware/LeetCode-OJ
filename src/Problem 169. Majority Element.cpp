// 169. Majority Element

#include <vector>

using namespace std;

// Boyer-Moore Voting Algorithm
//   Count of a majority element must > n / 2, 
//     which means it is accepted to remove m other elements while remove m majority elements.
//     (n - m) / 2 > (n - 2m) / 2, so the majority element is not changed.
// 
//   Count = 0: we delete 2m element which may include m majority elements or not.
//     Both of the situations enlarge the ratio of majority element in array.
//     So the last candidate will be majority element.
class Solution {
public:
    int majorityElement(vector<int>& nums) {
        int candidate, count = 0;
        for (auto num : nums) {
            if (count == 0) candidate = num;
            if (num == candidate) count ++;
            else count--;
        }
        return candidate;
    }
};