// 898. Bitwise ORs of Subarrays

#include <bits/stdc++.h>
#include <vector>
#include <unordered_set>
using namespace std;

// Enum 3 and get 2 by min stack
class Solution {
public:
    int subarrayBitwiseORs(vector<int>& arr) {
        unordered_set<int> ret;
        for (int i = 0;i < arr.size();i++) {
            ret.insert(arr[i]);
            for (int j = i - 1;j >= 0;j--) {
                if ((arr[i] & arr[j]) == arr[i]) break;
                arr[j] |= arr[i];
                ret.insert(arr[j]);
            }
        }
        return ret.size();
    }
};