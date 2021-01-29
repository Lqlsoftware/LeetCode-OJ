// 210. Course Schedule II

#include <vector>

using namespace std;

class Solution {
private:
    int degree[2001];
public:
    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
        vector<vector<int>> adj(numCourses);
        vector<int> ret(numCourses);
        for (int i = 0;i < prerequisites.size();i++) {
            degree[prerequisites[i][0]]++;
            adj[prerequisites[i][1]].push_back(prerequisites[i][0]);
        }
        // Use ret as a queue;
        int curr = 0, tail = 0;
        // Out degree == 0 means it's an end node
        for (int i = 0;i < numCourses;i++) if (degree[i] == 0) ret[tail++] = i;
        while(curr < tail)
            for (int j : adj[ret[curr++]])
                if (--degree[j] == 0)
                    ret[tail++] = j;
        if (tail < numCourses) ret.clear();
        return ret;
    }
};