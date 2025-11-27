#pragma once
#include <iostream>
#include <string>
#include <vector>
#include <fstream>

using namespace std;

struct HashEntry {
    string key;
    string value;
    bool isOccupied = false;
    bool isDeleted = false;
};

class HashTable {
private:
    vector<HashEntry> table;
    int capacity;
    int size;

    size_t hashFunc(string key) {
        size_t hash = 0;
        for (char c : key) hash = hash * 31 + c;
        return hash % capacity;
    }

public:
    HashTable(int cap = 20) : capacity(cap), size(0) {
        table.resize(capacity);
    }

    void put(string key, string value) {
        size_t idx = hashFunc(key);
        while (table[idx].isOccupied) {
            if (table[idx].key == key && !table[idx].isDeleted) {
                table[idx].value = value;
                return;
            }
            idx = (idx + 1) % capacity;
        }
        table[idx] = {key, value, true, false};
        size++;
    }

    string get(string key) {
        size_t idx = hashFunc(key);
        size_t start = idx;
        while (table[idx].isOccupied || table[idx].isDeleted) {
            if (table[idx].isOccupied && table[idx].key == key) {
                return table[idx].value;
            }
            idx = (idx + 1) % capacity;
            if (idx == start) break;
        }
        return "";
    }
    
    // сохраняем пары ключ-значение
    void saveToText(string filename) {
        ofstream file(filename);
        file << size << endl;
        for(const auto& entry : table) {
            if(entry.isOccupied && !entry.isDeleted) {
                file << entry.key << endl << entry.value << endl;
            }
        }
        file.close();
    }

    void loadFromText(string filename) {
        ifstream file(filename);
        int count;
        file >> count;
        string trash, k, v;
        getline(file, trash);
        
        // пересоздаем таблицу
        table.assign(capacity, HashEntry());
        size = 0;
        
        for(int i=0; i<count; i++) {
            getline(file, k);
            getline(file, v);
            put(k, v);
        }
        file.close();
    }
};