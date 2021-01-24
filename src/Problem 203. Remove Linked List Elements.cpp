// 203. Remove Linked List Elements

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode newhead(0, head);
        ListNode *p = &newhead;
        while(p != nullptr && p->next != nullptr) {
            if (p->next->val == val) p->next = p->next->next;
            else p = p->next;
        }
        return newhead.next;
    }
};