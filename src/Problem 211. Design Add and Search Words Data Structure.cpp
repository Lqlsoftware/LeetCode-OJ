//211. Design Add and Search Words Data Structure

#include <string>

using namespace std;

class WordDictionary {
private:
    struct TireNode {
        bool  val;
        TireNode* next[26];
        TireNode() : val(false) {
            memset(next, 0, 26 * sizeof(TireNode*));
        }
        ~TireNode() {
            for (int i = 0;i < 26;i++)
                if (this->next[i] != nullptr) delete this->next[i];
        }
    };

    TireNode *root;
    
    bool dfsSearch(TireNode *curr, int idx, string word) {
        for (int i = idx;i < word.size();i++) {
            if (word[i] == '.') {
                for (int j = 0;j < 26;j++)
                    if (curr->next[j] != nullptr && dfsSearch(curr->next[j], i + 1, word)) return true;
                return false;
            }
            if (curr->next[word[i] - 'a'] != nullptr)
                curr = curr->next[word[i] - 'a'];
            else return false;
        }
        return curr->val;
    }
    
public:
    /** Initialize your data structure here. */
    WordDictionary() {
        root = new TireNode();
    }
    
    ~WordDictionary() {
        delete root;
    }
    
    void addWord(string word) {
        TireNode *p = root;
        for (int i = 0;i < word.size();i++) {
            if (p->next[word[i] - 'a'] == nullptr) p->next[word[i] - 'a'] = new TireNode();
            p = p->next[word[i] - 'a'];
        }
        p->val = true;
    }
    
    bool search(string word) {
        return dfsSearch(root, 0, word);
    }
};