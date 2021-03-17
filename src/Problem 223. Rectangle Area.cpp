// 223. Rectangle Area

#include <iostream>
using namespace std;

#define min(a,b) (a < b ? a : b)

class Solution {
public:
    int computeArea(int A, int B, int C, int D, int E, int F, int G, int H) { 
        int w1 = C - A, h1 = D - B;
        int w2 = G - E, h2 = H - F;
        int s = w1 * h1 + w2 * h2;
        
        /* diff of the center (to reduce float ops, all plus by 2) */
        int diff_x = abs((A + C) - (E + G));
        int diff_y = abs((B + D) - (F + H));

        /* No overlap */
        if (diff_x >= w1 + w2 || diff_y >= h1 + h2) return s;
        
        /* Overlap area, calculate by the diff of the center */
        int w, h;
        if (diff_x <= abs(w1 - w2)) w = min(w1, w2);
        else w = (w1 + w2 - diff_x) / 2;
        
        if (diff_y <= abs(h1 - h2)) h = min(h1, h2);
        else h = (h1 + h2 - diff_y) / 2;
        
        return s - w * h;
    }
};