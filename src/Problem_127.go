/*
Word Ladder
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:
Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
For example,

Given:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log","cog"]
As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.

Note:
Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
*/

package main

// 16ms solution 利用map的hash来减少字符串比较的运算时间
// 用全局变量减少栈空间的申请次数
var exist = struct{}{}
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// wordSet用来标记是否访问过,forward是目标访问集合,backward是结果所在集合
	wordSet, forward, backward := make(map[string]struct{},len(wordList)), make(map[string]struct{},1), make(map[string]struct{},1)
	for _,v := range wordList {
		wordSet[v] = exist
	}
	if _,ok := wordSet[endWord];!ok {
		return 0
	}
	delete(wordSet,beginWord)
	forward[beginWord],backward[endWord] = exist,exist
	count := 2
	// forward 大小一直小于 backward
	for len(forward) > 0 {
		// 总是对个数少的集合进行访问
		if len(backward) < len(forward) {
			forward, backward = backward, forward
		}
		next := make(map[string]struct{},0)
		for v := range forward {
			word := []byte(v)
			for i := range word {
				char := word[i]
				// 对每个字母从A-Z进行更改
				for word[i] = 'a'; word[i] <= 'z'; word[i]++ {
					if word[i] == char {
						continue
					}
					s := string(word)
					if _, ok := backward[s]; ok {
						return count
					}
					if _, ok := wordSet[s]; ok {
						next[s] = exist
						delete(wordSet, s)
					}
				}
				word[i] = char
			}
		}
		forward = next
		count++
	}
	return 0
}