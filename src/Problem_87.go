/*
Scramble String
Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
To scramble the string, we may choose any non-leaf node and swap its two children.
For example, if we choose the node "gr" and swap its two children, it produces a scrambled string "rgeat".

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
We say that "rgeat" is a scrambled string of "great".
Similarly, if we continue to swap the children of nodes "eat" and "at", it produces a scrambled string "rgtae".

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
We say that "rgtae" is a scrambled string of "great".
Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.
*/

package main

/*
  	S1 [   x1    |         x2         ]

  	S2 [   y1    |         y2         ]
  	or
  	S2 [       y3        |     y4     ]

  	isScramble(x1,y1) && isScramble(x2,y2)
		|| isScramble(x1,y4) && isScramble(x2,y3)
*/
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return solve(s1,s2)
}

func solve(s1 string, s2 string) bool {
	// 剪枝
	if s1 == s2 {
		return true
	}
	char := make([]int,26)
	for i := range s1 {
		char[s1[i] - 'a']++
		char[s2[i] - 'a']--
	}
	for _,v := range char {
		if v != 0 {
			return false
		}
	}
	for l,i := len(s1),1;i < l;i++ {
		if a,b := s1[i:],s1[:i];solve(a,s2[i:]) && solve(b,s2[:i]) || solve(a,s2[:l - i]) && solve(b,s2[l - i:]) {
			return true
		}
	}
	return false
}

/*
  DP solution 通解
  S1 [   x1    |         x2         ]
     i         i + q                i + k - 1

  here we have two possibilities:

  S2 [   y1    |         y2         ]
     j         j + q                j + k - 1

  or

  S2 [       y1        |     y2     ]
     j                 j + k - q    j + k - 1
*/
func isScramble1(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	l := len(s1)
	dp := make([][][]bool,l)
	for i := range dp {
		dp[i] = make([][]bool,l)
		for j := range dp[i] {
			dp[i][j] = make([]bool,l + 1)
			dp[i][j][1] = s1[i] == s2[j]
		}
	}
	for k := 2;k <= l;k++ {
		for i := 0;i + k <= l;i++ {
			for j := 0;j + k <= l;j++ {
				for q := 1;q < k && !dp[i][j][k];q++ {
					dp[i][j][k] = dp[i][j][q] && dp[i + q][j + q][k - q] || dp[i][j + k - q][q] && dp[i + q][j][k - q]
				}
			}
		}
	}
	return dp[0][0][l]
}