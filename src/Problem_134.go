/*
Gas Station
There are N gas stations along a circular route, where the amount of gas at station i is gas[i].
You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to its next station (i+1).
You begin the journey with an empty tank at one of the gas stations.
Return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1.

Note:
If there exists a solution, it is guaranteed to be unique.
Both input arrays are non-empty and have the same length.
Each element in the input arrays is a non-negative integer.

Example 1:
Input:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.

Example 2:
Input:
gas  = [2,3,4]
cost = [3,4,3]
Output: -1
Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.
*/

package main

// 取出每个加油站并前进增加的油(gas - cost)
// 查找增加的油中最后一个和大于零的subarray的起始位置
// subarray和小于0说明从0到i(subarray结束的idx)都无法到达i + 1
// 所以从i+1开始继续寻找
// 如果存在解 肯定是最后一个和大于零的subarray的起始位置
func canCompleteCircuit(gas []int, cost []int) int {
	sum,max,res := 0,0,0
	for i := range gas {
		diff := gas[i] - cost[i]
		sum += diff
		max += diff
		if max < 0 {
			max = 0
			res = i + 1
		}
	}
	if sum < 0 {
		return -1
	}
	return res
}

// dfs
func canCompleteCircuit1(gas []int, cost []int) int {
	for start := range gas {
		if gas[start] < cost[start] {
			continue
		}
		if dfs(gas, cost, start, start, gas[start]) {
			return start
		}
	}
	return -1
}

func dfs(gas []int, cost []int, start, current, tank int) bool {
	next := (current + 1) % len(gas)
	tank -= cost[current]
	if tank < 0 {
		return false
	} else if next == start {
		return true
	}
	return dfs(gas, cost, start, next, tank + gas[next])
}