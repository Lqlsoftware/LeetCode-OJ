// 1492. The kth Factor of n

#include <iostream>

using namespace std;

class Solution {
public:
    int kthFactor(int n, int k) {
        int sqrt_n = floor(sqrt(n));
        for (int i = 1;i <= sqrt_n;i++) {
            if (n % i == 0) k--;
            if (k == 0) return i;
        }
        if (sqrt_n * sqrt_n == n) k++;
        for (int i = sqrt_n;i > 0;i--) {
            if (n % i == 0) k--;
            if (k == 0) return (n / i);
        }
        return -1;
    }
};