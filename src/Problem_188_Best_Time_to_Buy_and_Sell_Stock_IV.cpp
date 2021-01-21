/*
188. Best Time to Buy and Sell Stock IV
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
Design an algorithm to find the maximum profit. You may complete at most k transactions.
Notice that you may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Example 1:
Input: k = 2, prices = [2,4,1]
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.

Example 2:
Input: k = 2, prices = [3,2,6,5,0,3]
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
 
Constraints:
0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000
*/
#include <vector>

using namespace std;

/*  dp[i][j]: the max profit in day j under max of i transactions.
 */
class Solution {
private:
    int dp[128][1000];
public:
    int maxProfit(int k, vector<int>& prices) {
        if (k == 0 || prices.size() < 2) return 0;
        for (int i = 1;i <= k;i++) {
            int buy = -prices[0];
            for (int j = 1;j < prices.size();j++) {
                /* 
                 * dp[i][j] = dp[i][j - 1];
                 * for (int l = 0;l < j;l++)
                 *    dp[i][j] = max(dp[i][j], max(dp[i][l], dp[i - 1][l] - prices[l] + prices[j]));
                 * We actually don't need to calculate all the diff between j and l
                 * Record the max of (dp[i - 1][l] - prices[l])
                 */
                dp[i][j] = max(dp[i][j - 1], buy + prices[j]);
                buy = max(dp[i - 1][j] - prices[j], buy);
            }
        }
        return dp[k][prices.size() - 1];
    }
};