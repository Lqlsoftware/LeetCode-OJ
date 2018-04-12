/*
Word Ladder II
Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:
Only one letter can be changed at a time
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.

For example,
Given:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log","cog"]
Return
  [
    ["hit","hot","dot","dog","cog"],
    ["hit","hot","lot","log","cog"]
  ]

Note:
Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
*/

package main

// 同Problem 127, 用map存储BFS结果 DFS生成
var exist = struct{}{}
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet, forward, backward := make(map[string]struct{},len(wordList)), make(map[string]struct{},1), make(map[string]struct{},1)
	// 从end向begin保存相邻的节点结果
	resmap := make(map[string][]string)
	result := make([][]string,0)
	for _,v := range wordList {
		wordSet[v] = exist
		resmap[v] = []string{}
	}
	if _,ok := wordSet[endWord];!ok {
		return result
	}
	// 确保最短
	done := false
	// 标记是否翻转 放入map中是否翻转
	flip := false
	delete(wordSet,endWord)
	forward[beginWord],backward[endWord] = exist,exist
	for len(forward) > 0 {
		for v := range forward {
			delete(wordSet,v)
		}
		if len(forward) > len(backward) {
			forward,backward = backward,forward
			flip = !flip
		}
		next := make(map[string]struct{},1)
		for v := range forward {
			word := []byte(v)
			for i := range word {
				char := word[i]
				for word[i] = 'a'; word[i] <= 'z'; word[i]++ {
					if word[i] == char {
						continue
					}
					s := string(word)
					if _, ok := backward[s]; ok {
						if flip {
							resmap[v] = append(resmap[v], s)
						} else {
							resmap[s] = append(resmap[s], v)
						}
						done = true
					}
					if _, ok := wordSet[s]; ok {
						if flip {
							resmap[v] = append(resmap[v], s)
						} else {
							resmap[s] = append(resmap[s], v)
						}
						next[s] = exist
					}
				}
				word[i] = char
			}
		}
		forward = next
		if done {
			break
		}
	}
	// dfs从map中生成所有可行解
	return generate(resmap, endWord, beginWord)
}

func generate(wordSet map[string][]string, current, endWord string) [][]string {
	if current == endWord {
		return [][]string{{endWord}}
	}
	var res [][]string
	for _,v := range wordSet[current] {
		res = append(res,generate(wordSet, v, endWord)...)
	}
	for i := range res {
		res[i] = append(res[i],current)
	}
	return res
}