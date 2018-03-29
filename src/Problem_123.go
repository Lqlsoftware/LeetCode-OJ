/*
Best Time to Buy and Sell Stock III
Say you have an array for which the ith element is the price of a given stock on day i.
Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
*/

package main

// an elegant solution
func maxProfit(prices []int) int {
	buy1, buy2, sell1, sell2 := math.MinInt32, math.MinInt32, 0, 0
	for _, p := range prices {
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1 + p)
		buy2 = max(buy2, sell1 - p)
		sell2 = max(sell2, buy2 + p)
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// my 8ms solution
func maxProfit(prices []int) int {
	// find max of all
	transactionsOne,transactionsTwo,minLeft,maxLeft := 0,0,0,math.MinInt32
	start,end := 0, 0
	for i,v := range prices {
		profit := v - prices[minLeft]
		if profit < 0 {
			minLeft = i
		} else if profit > transactionsOne {
			transactionsOne = profit
			start,end = minLeft,i
		}
	}
	// find inside
	for i := start + 1;i < end;i++ {
		profit := maxLeft - prices[i]
		if profit < 0 {
			maxLeft = prices[i]
		} else if profit > transactionsTwo {
			transactionsTwo = profit
		}
	}
	// find outside
	if front := findProfit(prices,0,start);front > transactionsTwo {
		transactionsTwo = front
	}
	if back := findProfit(prices,end,len(prices));back > transactionsTwo {
		transactionsTwo = back
	}
	return transactionsOne + transactionsTwo
}

func findProfit(prices []int, start, end int) int {
	result,minLeft := 0,math.MaxInt32
	for i := start;i < end;i++ {
		profit := prices[i] - minLeft
		if profit < 0 {
			minLeft = prices[i]
		} else if profit > result {
			result = profit
		}
	}
	return result
}