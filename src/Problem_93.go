/*
Restore IP Addresses
Given a string containing only digits, restore it by returning all possible valid IP address combinations.

For example:
Given "25525511135",
return ["255.255.11.135", "255.255.111.35"]. (Order does not matter)
*/

package main

// DFS
func restoreIpAddresses(s string) []string {
	result := &[]string{}
	solve(result, s, "", 0, 0)
	return *result
}

func solve(result *[]string, s, current string, cur, count int) {
	if count == 4 && cur == len(s) {
		current = current[:len(current) - 1]
		*result = append(*result, current)
		return
	} else if cur + 12 - count * 3 < len(s) || cur + 4 - count > len(s) {
		return
	}
	for sum,i := 0,cur;i < cur + 3;i++ {
		sum = sum * 10 + int(s[i] - '0')
		if sum < 256 {
			current += string(s[i])
			solve(result, s, current + ".", i + 1, count + 1)
		}
		if sum == 0 || i == len(s) - 1 {
			return
		}
	}
}