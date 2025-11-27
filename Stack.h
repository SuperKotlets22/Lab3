#pragma once
#include "DynamicArray.h"

class Stack {
private:
    DynamicArray arr;
public:
    void push(string val) {
        arr.pushBack(val);
    }

    string pop() {
        if (arr.getSize() == 0) return "";
        
        // Берем верхний элемент
        string val = arr.get(arr.getSize() - 1);
        
        // ВАЖНО: Удаляем его из массива!
        arr.removeLast(); 
        
        return val;
    }

    void saveToText(string filename) {
        arr.saveToText(filename);
    }
    
    void loadFromText(string filename) {
        arr.loadFromText(filename);
    }
};