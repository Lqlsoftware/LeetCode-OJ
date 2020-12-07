#include <unordered_map>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// O(N)
class Solution {
public:
    bool hasCycle(ListNode *head) {
        while (head != nullptr) {
            map[head] = true;
            if (map.find(head->next) != map.end()) return true;
            head = head->next;
        }
        return false;
    }
private:
    unordered_map<ListNode*, bool> map;
};

// O(N) fast visit twice of the list, if there is cycle it can be sure that fast will equal to slow
class Solution1 {
public:
    bool hasCycle(ListNode *head) {
        if(head == nullptr) return false;
        ListNode *slow = head;
        ListNode *fast = head->next;
        
        while(slow != fast) {
            if(fast == nullptr || fast->next == nullptr)
                return false;
            fast = fast->next->next;
            slow = slow->next;
        }
        return true;
    }
};