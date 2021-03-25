// 667. Beautiful Arrangement II

#include <vector>

using namespace std;

/*
    {Max Min Max-1 Min+1 ... Max-k/2 Min+k/2} (Max-k/2-1) {Min+k/2+1 Min+k/2+2 ... Min+k/2+n-k}
    diff:             k - 1                                                1

    etc: 8 1 7 2 6 3 5 4 - 7 6 5 4 3 2 1
*/
class Solution {
public:
    vector<int> constructArray(int n, int k) {
        int s = 1; int e = n;
        vector<int> ret(n);
        k = k - 1;
        for (int i = 0;i < k - 1;i += 2) ret[i] = s++;
        for (int i = 1;i < k - 1;i += 2) ret[i] = e--;
        if (k % 2 == 0) for (int i = k - 1;i < n;i++) ret[i] = s++;
        else            for (int i = k - 1;i < n;i++) ret[i] = e--;
        return ret;
    }
};