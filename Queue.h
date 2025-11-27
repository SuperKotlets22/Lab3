#pragma once
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

class Queue {
private:
    string* data;
    int capacity;
    int size;
    int head;
    int tail;

public:
    Queue(int cap = 10) : capacity(cap), size(0), head(0), tail(0) {
        data = new string[capacity];
    }
    
    ~Queue() { delete[] data; }

    void push(string val) {
        if (size == capacity) return; // или ресайз
        data[tail] = val;
        tail = (tail + 1) % capacity;
        size++;
    }

    string pop() {
        if (size == 0) return "";
        string val = data[head];
        head = (head + 1) % capacity;
        size--;
        return val;
    }
    
    // сериализация сохраняет только активные элементы по порядку
    void saveToText(string filename) {
        ofstream file(filename);
        file << size << endl;
        for(int i=0; i<size; i++) {
            file << data[(head + i) % capacity] << endl;
        }
        file.close();
    }
    
    void loadFromText(string filename) {
        ifstream file(filename);
        int newSize;
        file >> newSize;
        string trash, val;
        getline(file, trash);
        
        // сброс
        head = tail = size = 0;
        for(int i=0; i<newSize; i++) {
            getline(file, val);
            push(val);
        }
        file.close();
    }
};