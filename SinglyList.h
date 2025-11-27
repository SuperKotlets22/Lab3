#pragma once
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

struct SNode {
    string val;
    SNode* next;
    SNode(string v) : val(v), next(nullptr) {}
};

class SinglyList {
private:
    SNode* head;
    SNode* tail;
    int size;

public:
    SinglyList() : head(nullptr), tail(nullptr), size(0) {}

    ~SinglyList() {
        while (head) {
            SNode* temp = head;
            head = head->next;
            delete temp;
        }
    }

    void pushBack(string val) {
        SNode* node = new SNode(val);
        if (!head) head = tail = node;
        else {
            tail->next = node;
            tail = node;
        }
        size++;
    }

    string getHead() {
        if (!head) return "";
        return head->val;
    }

    void saveToText(string filename) {
        ofstream file(filename);
        file << size << endl;
        SNode* curr = head;
        while (curr) {
            file << curr->val << endl;
            curr = curr->next;
        }
        file.close();
    }

    void loadFromText(string filename) {
        ifstream file(filename);
        int newSize;
        file >> newSize;
        string trash, val;
        getline(file, trash);
        
        // очистка не реализована для краткости, подразумевается пустой список
        for(int i=0; i<newSize; i++) {
            getline(file, val);
            pushBack(val);
        }
        file.close();
    }
    
    // бинарная сериализация аналогична массиву, идем по узлам
    void saveToBinary(string filename) {
        ofstream file(filename, ios::binary);
        file.write((char*)&size, sizeof(size));
        SNode* curr = head;
        while(curr) {
            size_t len = curr->val.size();
            file.write((char*)&len, sizeof(len));
            file.write(curr->val.c_str(), len);
            curr = curr->next;
        }
        file.close();
    }
};