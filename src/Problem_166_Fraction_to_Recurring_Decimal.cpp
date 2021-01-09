/*
166. Fraction to Recurring Decimal
Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.
If the fractional part is repeating, enclose the repeating part in parentheses.
If multiple answers are possible, return any of them.
It is guaranteed that the length of the answer string is less than 104 for all the given inputs.

Example 1:
Input: numerator = 1, denominator = 2
Output: "0.5"

Example 2:
Input: numerator = 2, denominator = 1
Output: "2"

Example 3:
Input: numerator = 2, denominator = 3
Output: "0.(6)"

Example 4:
Input: numerator = 4, denominator = 333
Output: "0.(012)"

Example 5:
Input: numerator = 1, denominator = 5
Output: "0.2"

Constraints:
-231 <= numerator, denominator <= 231 - 1
denominator != 0
*/
#include <string>
#include <unordered_map>
using namespace std;

// 0ms
class Solution {
public:
    string fractionToDecimal(int numerator, int denominator) {
        if (numerator == 0) return "0";
        bool neg = (numerator ^ denominator) < 0;
        long num = abs(numerator), den = abs(denominator);
        string ret = neg ? "-" + to_string(num / den) : to_string(num / den);
        
        // Is integer?
        num = num % den;
        if (num == 0) return ret;
        else ret += ".";

        // Get divided, record remainder on each pow of 10
        //   Until repeat or remainder == 0
        unordered_map<int, bool> m;
        long rem = num;
        while (m.find(rem) == m.end() && rem != 0) {
            m[rem] = true;
            rem = (rem * 10) % den;
        }
        
        // Build string
        while (num != rem) {
            num *= 10;
            ret += to_string(num / den);
            num %= den;
        }

        // Remainder dividable
        if (num == 0) return ret;
        
        // Repeat part
        ret += "(";
        do {
            num *= 10;
            ret += to_string(num / den);
            num %= den;
        } while (num != rem);
        return ret + ")";
    }
};