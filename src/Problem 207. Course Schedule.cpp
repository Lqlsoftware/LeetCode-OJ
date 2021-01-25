// 207. Course Schedule
// Deterime that a directed graph is a DAG or not

// DFS to find the cycle in Directed Graph
class Solution {
private:
    bool v[10000];
    bool ok[10000];
    unordered_map<int, vector<int>> node;
    bool visit(int curr) {
        if (ok[curr] == true) return true;
        if (v[curr] == true) return false;
        v[curr] = true;
        for (int vert : node[curr]) if (!visit(vert)) return false;
        v[curr] = false;
        return ok[curr] = true;
    }
public:
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        for (auto edge : prerequisites) node[edge[0]].push_back(edge[1]);
        for (int i = 0;i < numCourses;i++) if (!visit(i)) return false;
        return true;
    }
};

/* Find the node that out degree = 0,
    then follow the edge to sub the previous nodes' out degree by -1
*/
class Solution {
    int degree[10000];
public:
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        stack<int> s;
        for (int i = 0;i < prerequisites.size();i++) degree[prerequisites[i][0]]++;
        for (int i = 0;i < numCourses; i++) if (degree[i] == 0) s.push(i);
        int count = 0;
        while(!s.empty()) {
            int curr = s.top(); s.pop();
            count++;
            for (int i = 0;i < prerequisites.size();i++) {
                if (prerequisites[i][1] == curr) {
                    degree[prerequisites[i][0]]--;
                    if (degree[prerequisites[i][0]] == 0)
                        s.push(prerequisites[i][0]);
                }
            }
        }
        
        return count == numCourses;
    }
};