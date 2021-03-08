// 216. Combination Sum III

#include <vector>

using namespace std;

class Solution {
private:
    vector<vector<int>> res;
    void dfs(vector<int>& curr, int k, int n) {
        if (k == 0 && n == 0) {
            vector<int> RET(curr.begin(), curr.end());
            res.push_back(RET);
            return;
        }
        if (n == 0) return;
        k--;
        vector<int> RET(curr.begin(), curr.end());
        for (int i = (curr.empty() ? 1 : curr.back() + 1);i < 10;i++) {
            if (i > n) break;
            RET.push_back(i);
            dfs(RET, k, n - i);
            RET.pop_back();
        }
    }
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<int> curr;
        dfs(curr, k, n);
        return res;
    }
};