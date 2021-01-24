// 202. Happy Number

/*  2, 3, 5, 6, 8, 9 -> 4 -> endless
    7 -> 1
*/
class Solution {
private:
    static inline int squareOfDigits(int n) {
        int ret = 0;
        while (n > 0) {
            ret += (n % 10) * (n % 10);
            n /= 10;
        }
        return ret;
    }
public:
    bool isHappy(int n) {
        while(n >= 10) n = squareOfDigits(n);
        return n == 1 || n == 7;
    }
};