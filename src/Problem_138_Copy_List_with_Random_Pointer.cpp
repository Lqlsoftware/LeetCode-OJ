#include <unordered_map>
#include <vector>
using namespace std;

// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = nullptr;
        random = nullptr;
    }
};

/*
    Simple sulotion
        Using unordered map to project orignal Node* to new Node*
        require 2N list access and N map finding
        2N + N ~ O(N)
        
        map:
            Node1 -> NewNode1
            Node2 -> NewNode2
            ...

        NewNode1.random = Node1.random
        NewNode1.random <- map(NewNode1.random)

*/
class Solution {
public:
    Node* copyRandomList(Node* head) {
        // nullptr
        if (head == nullptr) return nullptr;

        // new head
        Node *new_head = new Node(head->val);
        new_head->random = head->random;
        map[head] = new_head;

        // Build new list (Copy random pointer & insert projection into map)
        Node *p = new_head, *q = head->next;
        while (q != nullptr) {
            Node *next = new Node(q->val);
            next->random = q->random;
            map[q] = next;
            p->next = next;
            p = p->next;
            q = q->next;
        }
        
        // Using map to build random connections of new list 
        p = new_head;
        while (p != nullptr) {
            if (p->random != nullptr) p->random = map[p->random];
            p = p->next;
        }
        return new_head;
    }
private:
    unordered_map<Node*, Node*> map;
};

/*
    Advanced sulotion
        Not simple but efficent
        Connect new & old list, make new[i].random = old[i], old[i].next = new[i]
        So that random pointer can be easily build that new.random = new.random.random.next
        5N ~ O(N)
        (And we must recover the orignal list)
        NewNode1 .next   ->                          NewNode2  .next   ->   NewNode3
                 .random -> Node1.next -> NewNode1             .random -> Node2.next -> NewNode2
*/
class Solution1 {
public:
    Node* copyRandomList(Node* head) {
        if (head == nullptr) return nullptr;

        vector<Node*> v;
        Node *p, *new_head, *q, *temp;
        // Backup sequence
        for (p = head;p != nullptr;p = p->next) {
            v.push_back(p->next);
        } 

        // Build new list
        new_head = new Node(head->val);
        p = new_head;
        for (int i = 0;i < v.size() - 1;i++) {
            p->next = new Node(v[i]->val);
            p = p->next;
        }

        // Connect new and old list
        p = new_head;
        q = head;
        while (p != nullptr) {
            temp = q->next;
            q->next = p;
            p->random = q;
            q = temp;
            p = p->next;
        }
        
        // Build random pointers
        for (p = new_head;p != nullptr;p = p->next)
            if (p->random->random != nullptr)
                p->random = p->random->random->next;
            else
                p->random = nullptr;

        // Recover from backup
        p = head;
        for (int i = 0;i < v.size();i++) {
            p->next = v[i];
            p = p->next;
        }
        return new_head;
    }
};