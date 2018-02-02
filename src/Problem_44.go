/*
Wildcard Matching
Implement wildcard pattern matching with support for '?' and '*'.
'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "*") → true
isMatch("aa", "a*") → true
isMatch("ab", "?*") → true
isMatch("aab", "c*a*b") → false
*/

package main

func isMatch(s string, p string) bool {
	i,j,k,l := 0,0,-1,0
	for i < len(s) {
		if j < len(p) && (p[j] == '?' || p[j] == s[i]) {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			k = j
			l = i
			j++
		} else if k != -1 {
			j = k + 1
			i = l + 1
			l = i
		} else {
			return false
		}
	}
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}

// 下面代码TLE了 尽量不要用递归
func _isMatch(s string, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		} else {
			return false
		}
	}
	if len(s) > 0 && (p[0] == '?' || uint8(p[0]) == s[0]) {
		return isMatch(s[1:],p[1:])
	} else if p[0] == '*' {
		// 下一个匹配的字符
		for i := 1;i < len(p);i++ {
			if p[i] != '*' {
				// 空串
				if isMatch(s,p[i:]) {
					return true
				}
				for j := 0;j < len(s);j++ {
					if (s[j] == p[i] || p[i] == '?') && isMatch(s[j:],p[i:]) {
						return true
					}
				}
				return isMatch(s[len(s):],p[i:])
			}
		}
		return true
	}
	return false
}
