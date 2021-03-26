// 856. Score of Parentheses

#include <iostream>
#include <string>
#include <stack>

using namespace std;

// Stack
class Solution {
public:
    int scoreOfParentheses(string S) {
        stack<int> stk;
        stk.push(0);
        for (auto ch : S) {
            if (ch == '(') {
                stk.push(0);
            } else {
                int temp = max(stk.top() * 2, 1);
                stk.pop();
                temp += stk.top();
                stk.pop();
                stk.push(temp);
            }
        }
        return stk.top();
    }
};

// Recurrsion
class Solution {
private:
    char *start, *end;
    int _scoreOfParentheses() {
        int ret = 0;
        while (start != end) {
            if (*start == '(') {
                start++;
                ret += _scoreOfParentheses();
            } else if (*start == ')') {
                return max(2 * ret, 1);
            }
            start++;
        }
        return ret;
    }
public:
    int scoreOfParentheses(string S) {
        start = const_cast<char *>(S.data()), end = const_cast<char *>(S.data() + S.length());
        return _scoreOfParentheses();
    }
};