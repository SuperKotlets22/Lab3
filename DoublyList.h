#pragma once
#include <string>

using namespace std;

struct DNode {
    string val;
    DNode* next;
    DNode* prev;
    DNode(string v) : val(v), next(nullptr), prev(nullptr) {}
};

class DoublyList {
private:
    DNode* head;
    DNode* tail;
public:
    DoublyList() : head(nullptr), tail(nullptr) {}
    
    void pushBack(string val) {
        DNode* node = new DNode(val);
        if(!head) head = tail = node;
        else {
            tail->next = node;
            node->prev = tail;
            tail = node;
        }
    }
    
    string getTail() {
        if(!tail) return "";
        return tail->val;
    }
    
    // методы сериализации идентичны односвязному списку
    // (реализуй по аналогии с SinglyList, чтобы не дублировать код здесь)
};