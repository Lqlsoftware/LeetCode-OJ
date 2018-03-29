/*
Best Time to Buy and Sell Stock II
Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit.

You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times).
However, you may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
*/

package main

func maxProfit(prices []int) int {
	result,length := 0,len(prices) - 1
	for i := 0;i < length;i++ {
		if diff := prices[i + 1] - prices[i];diff > 0 {
			result += diff
		}
	}
	return result
}