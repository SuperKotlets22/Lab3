#pragma once
#include <iostream>
#include <string>
#include <fstream>

using namespace std;

class DynamicArray {
private:
    string* data;
    int size;
    int capacity;

    void resize() {
        capacity *= 2;
        string* newData = new string[capacity];
        for (int i = 0; i < size; i++) newData[i] = data[i];
        delete[] data;
        data = newData;
    }

public:
    // конструктор
    DynamicArray(int cap = 10) {
        capacity = cap;
        size = 0;
        data = new string[capacity];
    }

    // деструктор
    ~DynamicArray() {
        delete[] data;
    }

    void pushBack(string val) {
        if (size == capacity) resize();
        data[size++] = val;
    }

    string get(int index) {
        if (index < 0 || index >= size) throw out_of_range("индекс за границами");
        return data[index];
    }

    int getSize() { return size; }

    // Новый метод для удаления последнего элемента (нужен для Стека)
    void removeLast() {
        if (size > 0) {
            size--;
        }
    }

    // сохранение в текст
    void saveToText(string filename) {
        ofstream file(filename);
        file << size << endl;
        for (int i = 0; i < size; i++) file << data[i] << endl;
        file.close();
    }

    // загрузка из текста
    void loadFromText(string filename) {
        ifstream file(filename);
        if (!file.is_open()) return;
        int newSize;
        file >> newSize;
        string temp;
        getline(file, temp); // пропуск перевода строки
        
        // очистка старого (просто сбрасываем размер)
        size = 0;
        for (int i = 0; i < newSize; i++) {
            getline(file, temp);
            pushBack(temp);
        }
        file.close();
    }

    // бинарное сохранение
    void saveToBinary(string filename) {
        ofstream file(filename, ios::binary);
        file.write((char*)&size, sizeof(size));
        for (int i = 0; i < size; i++) {
            size_t len = data[i].size();
            file.write((char*)&len, sizeof(len));
            file.write(data[i].c_str(), len);
        }
        file.close();
    }

    // бинарная загрузка
    void loadFromBinary(string filename) {
        ifstream file(filename, ios::binary);
        if (!file.is_open()) return;
        int newSize;
        file.read((char*)&newSize, sizeof(newSize));
        
        size = 0; // сброс текущего
        for (int i = 0; i < newSize; i++) {
            size_t len;
            file.read((char*)&len, sizeof(len));
            char* buf = new char[len + 1];
            file.read(buf, len);
            buf[len] = '\0';
            pushBack(string(buf));
            delete[] buf;
        }
        file.close();
    }
};