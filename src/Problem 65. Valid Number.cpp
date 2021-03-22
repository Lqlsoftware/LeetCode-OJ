// Problem 65. Valid Number
#include <string>

using namespace std;

class Solution {
    bool dot(string& s, int& idx) {
        if (idx < s.length() && s[idx] == '.') {
            idx++;
            return true;
        }
        return false;
    }
    bool digits(string& s, int& idx) {
        bool flag = false;
        while (idx < s.length() && s[idx] >= '0' && s[idx] <= '9') {
            idx++;
            flag = true;
        }
        return flag;
    }
    bool sign(string& s, int& idx) {
        if (idx < s.length() && (s[idx] == '+' || s[idx] == '-')) {
            idx++;
            return true;
        }
        return false;
    }
    bool decimal(string& s, int& idx) {
        bool flag = false;
        sign(s, idx);
        flag = digits(s, idx);
        
        if (!dot(s, idx)) return false;
        
        if (!flag) return digits(s, idx);
        
        digits(s, idx);
        return true;
    }
    bool integer(string& s, int& idx) {
        sign(s, idx);
        return digits(s, idx);
    }
    bool e(string& s, int& idx) {
        if (idx < s.length() && (s[idx] == 'e' || s[idx] == 'E')) {
            idx++;
            return true;
        }
        return false;
    }
public:
    bool isNumber(string s) {
        bool flag = false;
        int idx1 = 0, idx2 = 0;
        if (decimal(s, idx1)) {
            if (idx1 == s.length()) return true;
            if (e(s, idx1) && integer(s, idx1)) return idx1 == s.length();
        } else if (integer(s, idx2)) {
            if (idx2 == s.length()) return true;
            if (e(s, idx2) && integer(s, idx2)) return idx2 == s.length();
        }
        return false;
    }
};