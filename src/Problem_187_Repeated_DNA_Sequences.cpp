/*
187. Repeated DNA Sequences
All DNA is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T', 
for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.
Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

Example 1:
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]

Example 2:
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]

Constraints:
0 <= s.length <= 105
s[i] is 'A', 'C', 'G', or 'T'.
*/
#include <vector>
#include <set>
#include <unordered_map>

using namespace std;
class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        unordered_map<string, bool> m;
        set<string> rset;
        string str;
        for (int i = 0;i <= int(s.size()) - 10;i++) {
            str = s.substr(i, 10);
            if (m.find(str) == m.end()) m[str] = false;
            else rset.insert(str);
        }
        vector<string> ret(rset.begin(), rset.end());
        return ret;
    }
};

// 4ms solution
class Solution2 {
	inline bool testAndSetBit(unsigned int *intArr, unsigned int offset) {
		unsigned int byteOff = offset >> 5;
		unsigned int bitOff = offset & ((1 << 5) - 1);

		unsigned int bit = intArr[byteOff] & (1 << bitOff);
		if (bit) {
			return true;
		} else {
			intArr[byteOff] |= 1 << bitOff;
			return false;
		}
	}
public:
	vector<string> findRepeatedDnaSequences(string s) {
		vector<string> result;
		if (s.length() < 11) {
			return result;
		}

		static const int intArrSize = 1 << (2 * 10 - 5);
		unsigned int dataArr[intArrSize] = {0};
		unsigned int recordArr[intArrSize] = {0};

		const char *str = s.c_str();
		int size = s.length();
		static const unsigned int mask = ((1 << 20) - 1);

		unsigned int data = 0;
		for (int i = 0; i < 10; ++i) {
			unsigned int twobits = (str[i] >> 1) & 0x3;
			twobits <<= ((9 - i) << 1);
			data |= twobits;
		}
		testAndSetBit(dataArr, data);

		char dna[11] = { 0 };
		for (int i = 10; i < size; ++i) {
			unsigned int twobits = (str[i] >> 1) & 0x3;

			data = ((data << 2) | twobits) & mask;
			bool hasSet = testAndSetBit(dataArr, data);
			if (!hasSet) {
				continue;
			}

			hasSet = testAndSetBit(recordArr, data);
			if (hasSet) {
				continue;
			}
			memcpy(dna, str + i - 9, 10);
			result.push_back(dna);
		}

		return result;
	}
};