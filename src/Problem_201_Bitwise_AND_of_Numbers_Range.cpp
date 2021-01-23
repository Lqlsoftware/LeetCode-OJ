/*
201. Bitwise AND of Numbers Range
Given a range [m, n] where 0 <= m <= n <= 2147483647,
return the bitwise AND of all numbers in this range, inclusive.

Example 1:
Input: [5,7]
Output: 4

Example 2:
Input: [0,1]
Output: 0
*/
#include <cstdint>
/*
    The bitwise AND of numbers between m and n will not affect the common binary prefix(CBP) of them(because they are sharing the same CBP), so the idea is to find the CBP of m and n.

                            5   b'01 01
                            7   b'01 11
    Common binary prefix    4   b'01(00)
    We use an unsigned-int mask filled with 1 (0xffffffff) to AND with m and n to find the common binary prefix.
 */
class Solution {
public:
    int rangeBitwiseAnd(int m, int n) {
        uint32_t mask = 0xffffffff;
        while ((m & mask) != (n & mask)) mask <<= 1;
        return m & mask;
    }
};

/*
    Using the highest binary 1 of (m ^ n) (the highest difference bit) to build a mask
    0100 0000 -> 0110 0000 -> 0111 1000 -> 0111 1111
      reverse -> 1000 0000 -> mask
    m & mask -> Common binary prefix
 */
class Solution2 {
public:
    int rangeBitwiseAnd(int m, int n) {
        unsigned int x = m ^ n;
        x |= x >> 1, x |= x >> 2, x |= x >> 4, x |= x >> 8, x |= x >> 16;
        return m & ~x;  
    }
};