/*
Simplify Path
Given an absolute path for a file (Unix-style), simplify it.

For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"
click to show corner cases.

Corner Cases:
Did you consider the case where path = "/../"?
In this case, you should return "/".
Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
In this case, you should ignore redundant slashes and return "/home/foo".
*/

package main

func simplifyPath(path string) string {
	path += "/"
	stack, top, start := make([]string,len(path) / 2), -1, 1
	for i := 1;i < len(path);i++ {
		if path[i] == '/' {
			x := path[start:i]
			if x == ".." {
				if top != -1 {
					top--
				}
			} else if x != "." && x != "" {
				top++
				stack[top] = x
			}
			start = i + 1
		}
	}
	result := "/"
	for i := 0;i <= top;i++ {
		result += stack[i]
		if i != top {
			result += "/"
		}
	}
	return result
}

func simplifyPath1(path string) string {
	str,result := strings.Split(path, "/"),[]string{}
	for _,v := range str {
		switch v {
		case ".","":
			continue
		case "..":
			if len(result) != 0 {
				result = result[:len(result) - 1]
			}
		default:
			result = append(result, v)
		}
	}
	return "/" + strings.Join(result,"/")
}