/*
Valid Number
Validate if a given string is numeric.

Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
Note: It is intended for the problem statement to be ambiguous.
You should gather all requirements up front before implementing one.
*/

package main

// The worst problem i have ever met in this oj. Regex is good.
func isNumber(s string) bool {
	matched,_ := regexp.MatchString("^\\s*[+-]?(\\d+(\\.\\d*)?|\\.\\d+)(e[+-]?\\d+)?\\s*$", s)
	return matched
}