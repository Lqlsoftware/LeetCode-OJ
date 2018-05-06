/*
Candy
There are N children standing in a line. Each child is assigned a rating value.
You are giving candies to these children subjected to the following requirements:
Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
What is the minimum candies you must give?

Example 1:
Input: [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

Example 2:
Input: [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
             The third child gets 1 candy because it satisfies the above two conditions.
*/

package main

// 上升和下降时 糖果的分配都是以等差数列的形式进行分配
// 只需知道上升和下降的孩子数量即可
func candy(ratings []int) int {
	curr,last := 0,0
	raise,down := 0,0
	candies := 0
	for i := 1;i < len(ratings);i++ {
		curr = ratings[i] - ratings[i - 1]
		// valley or stage in upraise
		if last < 0 && curr >= 0 || last > 0 && curr == 0 {
			candies += count(down) + count(raise) + max(down, raise)
			raise = 0
			down = 0
		}
		if curr > 0 {
			raise++
		} else if curr < 0 {
			down++
		} else {
			candies++
		}
		last = curr
	}
	candies += count(down) + count(raise) + max(down, raise) + 1
	return candies
}

func count(n int) int {
	return n * (n + 1) / 2
}

// left to right and right to left
func candy1(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}
	for i := 1;i < len(candies);i++ {
		if ratings[i] > ratings[i - 1] {
			candies[i] = candies[i - 1] + 1
		}
	}
	res := candies[len(candies) - 1]
	for i := len(candies) - 2;i >= 0;i-- {
		if ratings[i] > ratings[i + 1] {
			candies[i] = max(candies[i], candies[i + 1] + 1)
		}
		res += candies[i]
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// dfs
func candy2(ratings []int) int {
	min := 0
	for i := range ratings {
		if ratings[i] < ratings[min] {
			min = i
		}
	}
	return dfs(ratings, min, 1, -1) + dfs(ratings, min, 1, 1) - 1
}

func dfs(ratings []int, curr, currCandy, toward int) int {
	if toward == -1 && curr == 0 || toward == 1 && curr == len(ratings) - 1 {
		return currCandy
	}
	start, end := 1, math.MaxInt32
	if ratings[curr + toward] > ratings[curr] {
		start = currCandy + 1
	} else if ratings[curr + toward] < ratings[curr] {
		end = currCandy
	}
	curr += toward
	for i := start;i < end;i++ {
		if res := dfs(ratings, curr, i, toward);res >= 0 {
			return res + currCandy
		}
	}
	return -1
}