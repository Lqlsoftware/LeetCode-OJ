#include <queue>
#include <vector>
using namespace std;

// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> neighbors;
    Node() {
        val = 0;
        neighbors = vector<Node*>();
    }
    Node(int _val) {
        val = _val;
        neighbors = vector<Node*>();
    }
    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};

class Solution {
public:
    Node* cloneGraph(Node* node) {
        if (node == nullptr) return nullptr;
        visit[node->val] = new Node(node->val);
        
        queue<Node*> q;
        q.push(node);

        while(!q.empty()) {
            Node *old_node = q.front(); q.pop();
            visit[old_node->val]->neighbors.reserve(old_node->neighbors.size());
            for (auto neighbor : old_node->neighbors) {
                if (visit[neighbor->val] == nullptr) {
                    visit[neighbor->val] = new Node(neighbor->val);
                    q.push(neighbor);
                }
                visit[old_node->val]->neighbors.push_back(visit[neighbor->val]);
            }
        }
        return visit[node->val];
    }
private:
    Node* visit[101];
};