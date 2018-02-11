/*
Text Justification
Given an array of words and a length L,
format the text such that each line has exactly L characters and is fully (left and right) justified.
You should pack your words in a greedy approach; that is,
pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly L characters.
Extra spaces between words should be distributed as evenly as possible.
If the number of spaces on a line do not divide evenly between words,
the empty slots on the left will be assigned more spaces than the slots on the right.
For the last line of text, it should be left justified and no extra space is inserted between words.

For example,
words: ["This", "is", "an", "example", "of", "text", "justification."]
L: 16.
Return the formatted lines as:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
Note: Each word is guaranteed not to exceed L in length.

Corner Cases:
A line other than the last line might contain only one word. What should you do in this case?
In this case, that line should be left-justified.
*/

package main

func fullJustify(words []string, maxWidth int) []string {
	if maxWidth == 0 {
		return []string{""}
	}
	stack,top,left := make([]int,maxWidth / 2 + 1),0,maxWidth - len(words[0])
	var res []string
	for i := 1;i < len(words);i++ {
		v := words[i]
		if len(v) < left {
			top++
			stack[top] = i
			left -= len(v) + 1
		} else {
			result := words[stack[0]]
			if top != 0 {
				average, empty := left / top, " "
				left = left % top
				for j := 0; j < average; j++ {
					empty += " "
				}
				for j := 1; j <= top; j++ {
					if j <= left {
						result += empty + " " + words[stack[j]]
					} else {
						result += empty + words[stack[j]]
					}
				}
			} else {
				for j, length := 0, maxWidth - len(result); j < length; j++ {
					result += " "
				}
			}
			res,left, stack[0], top = append(res, result),maxWidth-len(words[i]), i, 0
		}
	}
	result := words[stack[0]]
	for j := 1;j <= top;j++ {
		result += " " + words[stack[j]]
	}
	for j, length := 0, maxWidth - len(result);j < length;j++ {
		result += " "
	}
	return append(res, result)
}