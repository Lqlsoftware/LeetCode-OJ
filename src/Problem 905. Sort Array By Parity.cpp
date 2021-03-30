// 905. Sort Array By Parity

#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> sortArrayByParity(vector<int>& A) {
        int s = 0, e = int(A.size()) - 1;
        while (s < e) {
            while (s < e && A[e] % 2 == 1) e--;
            while (s < e && A[s] % 2 == 0) s++;
            if (s < e) swap(A[s], A[e]);
        }
        return A;
    }
};