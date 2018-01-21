/*
Regular Expression Matching
Implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true
*/

package main

/*   use dp
func isMatch(s string, p string) bool {
	dp := make([][]bool,len(s),len(p))
	dp[len(s) - 1][len(p) - 1] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			...
		}
	}
}
*/

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	str, regex := []rune(s), []rune(p)
	i, j := len(str)-1, len(regex)-1
	if j < 0 {
		return false
	}
	if regex[j] == '*' {
		if i >= 0 && (regex[j-1] == str[i] || regex[j-1] == '.') {
			return isMatch(s[:i+1], p[:j-1]) || isMatch(s[:i], p[:j+1]) || isMatch(s[:i], p[:j-1])
		}
		return isMatch(s[:i+1], p[:j-1])
	} else if i >= 0 && (regex[j] == '.' || regex[j] == str[i]) {
		return isMatch(s[:i], p[:j])
	}
	return false
}

/*
func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	str,regex := []rune(s),[]rune(p)
	i,j := len(str) - 1,len(regex) - 1
	if i < 0 {
		return true
	}
	if j < 0 {
		return false
	}
	switch regex[j] {
	case '.': return isMatch(s[:i],p[:j])
	case '*': {
		for i > 0 && j > 0 {
			if isMatch(s[:i],p[:j]) {
				return true
			}
			i--
		}
		return false
	}
	default:
		if str[i] == regex[j] {
			return isMatch(s[:i],p[:j])
		} else {
			return false
		}
	}
	return false
}
*/
