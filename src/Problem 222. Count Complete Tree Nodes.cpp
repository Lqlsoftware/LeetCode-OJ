// 222. Count Complete Tree Nodes

#include <iostream>
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
private:
    int ldepth(TreeNode* root) {
        int d = 0;
        for (TreeNode* curr = root; curr != nullptr;curr = curr->left)  d++;
        return d;
    }
public:
    int countNodes(TreeNode* root) {
        int dl = ldepth(root), drl, count = 0;
        while(root != nullptr) {
            drl = ldepth(root->right);
            dl = dl - 1;
            count += (1 << drl);

            /* root->left is a full binary tree 
             *   check root->right next
             *   dl <- drl - 1 = dl - 1
             */
            if (drl == dl) root = root->right;

            /* root->right is a full binary tree 
             *   check root->left next
             *   dl <- dl - 1
             */
            else root = root->left;
        }
        return count;
    }
};