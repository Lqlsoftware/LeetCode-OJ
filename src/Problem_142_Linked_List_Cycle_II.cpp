#include <unordered_map>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};

/*
    O(N) time O(1) space

    Meet in the Node(i) when fast take 2n step and slow take n step
    And fast already arrival at the Node once before.
    So the length of the cycle is n step.

                            fast
                            slow
    Node -> Node -> Node -> Node -> Node
                     ^_______________/

    slow begin at head pointer, and they both take 1 step at each time.

    slow                    fast
    Node -> Node -> Node -> Node -> Node
                     ^_______________/

    slow take n step to the Node(i), fast take n step(one cycle) to the Node(i).
    So they will meet at the Node that cycle started.

                    fast
                    slow
    Node -> Node -> Node -> Node -> Node
                     ^_______________/
*/
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        ListNode *slow = head, *fast = head;
        
        do {
            if (fast == nullptr || fast->next == nullptr) return nullptr;
            fast = fast->next->next;
            slow = slow->next;
        } while(slow != fast);
        
        slow = head;
        while(slow != fast) {
            slow = slow->next;
            fast = fast->next;
        }
        
        return slow;
    }
};

// O(N) time O(N) space
class Solution1 {
public:
    ListNode *detectCycle(ListNode *head) {
        while (head != nullptr) {
            map[head] = true;
            if (map.find(head->next) != map.end()) return head->next;
            head = head->next;
        }
        return nullptr;
    }
private:
    unordered_map<ListNode*, bool> map;
};