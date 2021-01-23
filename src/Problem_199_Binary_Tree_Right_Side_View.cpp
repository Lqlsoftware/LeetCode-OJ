/*
199. Binary Tree Right Side View
Given a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Example:
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/
#include <vector>
#include <queue>
#include <utility>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    vector<int> rightSideView(TreeNode* root) {
        vector<int> ret;
        if (root == nullptr) return ret;
        queue<pair<TreeNode*, int>> q;
        q.push(pair(root, 0));
        while(!q.empty()) {
            TreeNode* p = q.front().first;
            int dep = q.front().second;
            q.pop();
            if (dep == ret.size()) ret.push_back(p->val);
            if (p->right != nullptr) q.push(pair(p->right, dep + 1));
            if (p->left != nullptr)  q.push(pair(p->left, dep + 1));
        }
        return ret;
    }
};

class Solution2 {
private:
    void dfs(TreeNode* root, vector<int> &ret, int height) {
        if(!root) return;
        if(height == ret.size()) ret.push_back(root->val);
        dfs(root->right, ret, height+1);
        dfs(root->left , ret, height+1);
    }
public:
    vector<int> rightSideView(TreeNode* root) {
        vector<int> ret;
        dfs(root, ret, 0);
        return ret;
    }
};