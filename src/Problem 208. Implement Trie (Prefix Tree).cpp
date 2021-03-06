//208. Implement Trie (Prefix Tree)

#include <string>

using namespace std;

class Trie {
private:
    struct TireNode {
        bool val;
        TireNode* next[26];

        TireNode() : val(false) {
            memset(next, 0, sizeof(TireNode*) * 26);
        }
        ~TireNode() {
            for (int i = 0;i < 26;i++)
                if (this->next[i] != nullptr) delete this->next[i];
        }
    };
    
    TireNode *root;
public:
    /** Initialize your data structure here. */
    Trie() {
        root = new TireNode();
    }
    
    ~Trie() {
        delete root;
    }
    
    /** Inserts a word into the trie. */
    void insert(string word) {
        TireNode *p = root;
        for (int i = 0;i < word.size();i++) {
            // New node
            if (p->next[word[i] - 'a'] == nullptr) {
                p->next[word[i] - 'a'] = new TireNode();
            }
            p = p->next[word[i] - 'a'];
        }
        p->val = true;
    }
    
    /** Returns if the word is in the trie. */
    bool search(string word) {
        TireNode *p = root;
        for (int i = 0;i < word.size();i++) {
            if (p->next[word[i] - 'a'] == nullptr) {
                return false;
            }
            p = p->next[word[i] - 'a'];
        }
        return p->val;
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    bool startsWith(string prefix) {
        TireNode *p = root;
        for (int i = 0;i < prefix.size();i++) {
            if (p->next[prefix[i] - 'a'] == nullptr) {
                return false;
            }
            p = p->next[prefix[i] - 'a'];
        }
        return true;
    }
};