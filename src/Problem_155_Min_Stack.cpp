/*
155. Min Stack
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

Example 1:
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2

Constraints:
Methods pop, top and getMin operations will always be called on non-empty stacks.
*/
#include <memory>

// Double stack to store each min state
class MinStack {
public:
    /** initialize your data structure here. */
    MinStack() {
        topptr = -1;
        Stack = (val*)malloc(16000 * sizeof(int));
    }
    
    void push(int x) {
        int min = x;
        if (topptr >= 0 && min > Stack[topptr].m) min = Stack[topptr].m;
        Stack[++topptr].v = x;
        Stack[topptr].m = min;
    }
    
    void pop() {
        topptr--;
    }
    
    int top() {
        return Stack[topptr].v;
    }
    
    int getMin() {
        return Stack[topptr].m;
    }
private:
    struct val {
        int v;
        int m;
        val(){}
    };
    val *Stack;
    int topptr;
};

// Double stack to push only x <= minStack[minptr]
class MinStack {
public:
    /** initialize your data structure here. */
    MinStack() {
        topptr = -1;
        minptr = -1;
    }
    
    void push(int x) {
        Stack[++topptr] = x;
        if (minptr == -1 || x <= minStack[minptr]) minStack[++minptr] = x;
    }
    
    void pop() {
        if (minptr >= 0 && Stack[topptr] == minStack[minptr]) minptr--;
        topptr--;
    }
    
    int top() {
        return Stack[topptr];
    }
    
    int getMin() {
        return minStack[minptr];
    }
private:
    int Stack[7500];
    int minStack[20];
    int topptr;
    int minptr;
};

// Construct a list that store the min value
class MinStack_minList {
public:
    /** initialize your data structure here. */
    MinStack_minList() {
        topptr = -1;
        min = new ListNode(0);
    }
    
    void push(int x) {
        Stack[++topptr] = x;
        ListNode *n = new ListNode(x);
        ListNode *p = min;
        while(p->next != nullptr && x > p->next->val) p = p->next;
        n->next = p->next;
        p->next = n;
    }
    
    void pop() {
        ListNode *p = min, *q;
        while(p->next != nullptr && Stack[topptr] != p->next->val) p = p->next;
        q = p->next;
        p->next = q->next;
        delete q;
        topptr--;
    }
    
    int top() {
        return Stack[topptr];
    }
    
    int getMin() {
        return min->next->val;
    }
private:
    struct ListNode {
        int val;
        ListNode *next;
        
        ListNode(int x) {
            val = x;
            next = nullptr;
        }
    };
    ListNode *min;
    int Stack[8000];
    int topptr;
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(x);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */