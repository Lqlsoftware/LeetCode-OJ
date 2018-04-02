/*
Valid Palindrome
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

For example,
"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.

Note:
Have you consider that the string might be empty? This is a good question to ask during an interview.
For the purpose of this problem, we define empty string as valid palindrome.
*/

package main

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	i,j := 0,len(s) - 1
	for i <= j {
		if a,ok := check(&s,i);ok {
			i++
		} else if b,ok := check(&s,j);ok {
			j--
		} else {
			if a != b {
				return false
			}
			i++
			j--
		}
	}
	return true
}

// check and to lower case
func check(s *string, idx int) (uint8,bool) {
	r := (*s)[idx]
	if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
		return r, false
	} else if r >= 'A' && r <= 'Z' {
		return r + 32, false
	}
	return 0, true
}