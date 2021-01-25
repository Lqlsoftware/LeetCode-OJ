// 205. Isomorphic Strings

/*
    Two array to store the projection from s -> t and t -> s
    `char <-> char` can be zip in a 256 Byte array
*/
class Solution {
private:
    char ch[128][2];
public:
    bool isIsomorphic(string s, string t) {
        if (s.size() != t.size()) return false;
        // memset(ch, 0, 256 * sizeof(char));
        for (int i = 0;i < s.size();i++) {
            if (ch[s[i]][0] == 0) ch[s[i]][0] = t[i];
            else if (ch[s[i]][0] != t[i]) return false;
            if (ch[t[i]][1] == 0) ch[t[i]][1] = s[i];
            else if (ch[t[i]][1] != s[i]) return false; 
        }
        return true;
    }
};