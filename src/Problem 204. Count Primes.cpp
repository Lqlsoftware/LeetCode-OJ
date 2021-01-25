// 204. Count Primes

/* The Sieve of Eratosthenes
    Mark the non-prime numbers
    j = i * i (2 * i ~ (i - 1) * i is already marked)
    j += i (for each number that can be devided by i)
*/
class Solution {
public:
    int countPrimes(int n) {
        if (n == 0) return 0;
        bool primes[n];
        memset(primes, 0, n * sizeof(bool));
        int ret = 0;
        for (int i = 2; i * i < n; i++) {
            if (primes[i]) continue;
            for (int j = i * i;j < n;j += i) {
                primes[j] = true;
            } 
        }
        for (int i = 2; i < n; i++) if (!primes[i]) ret++;
        return ret;
    }
};