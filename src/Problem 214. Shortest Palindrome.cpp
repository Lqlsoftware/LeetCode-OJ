// 214. Shortest Palindrome

#include <string>

using namespace std;

/*
    ababacbaba        'a' 'b' 'a' 'b' ac 'b' 'a' 'b' 'a'
    |        |         |              |
    i        j         j              i

    j: n - 1 -> 0
    i: if s[i] == s[j] i++
    n - i: least length of unsymmetry sequence
    so the split point move to 0 ~ i, clearly we can narrow down the range i -> 0
*/
class Solution {
public:
    string shortestPalindrome(string s) {
        int n = s.size();
        int i = 0, j;
        for (j = n - 1; j >= 0; j--)
            if (s[i] == s[j]) i++;
        if (i == n) return s;
        string rev = s.substr(i);
        reverse(rev.begin(), rev.end());
        return rev + shortestPalindrome(s.substr(0, i)) + s.substr(i);
    }
};