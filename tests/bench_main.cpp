#include <iostream>
#include <chrono>
#include <vector>
#include <string>

// Подключаем ваши структуры
#include "DynamicArray.h"
#include "SinglyList.h"
#include "DoublyList.h"
#include "Stack.h"
#include "Queue.h"
#include "HashTable.h"
#include "CompleteBinaryTree.h"

using namespace std;
using namespace std::chrono;

// Вспомогательная функция для замера времени
template<typename Func>
void run_benchmark(const string& name, Func func) {
    auto start = high_resolution_clock::now();
    func();
    auto end = high_resolution_clock::now();
    auto duration = duration_cast<milliseconds>(end - start).count();
    cout << "[BENCHMARK] " << name << ": " << duration << " ms" << endl;
}

const int N = 100000; // Количество операций

int main() {
    cout << "=== Starting Benchmarks (N = " << N << ") ===" << endl;

    //  DYNAMIC ARRAY 
    run_benchmark("DynamicArray: PushBack", []() {
        DynamicArray arr;
        for(int i = 0; i < N; ++i) {
            arr.pushBack("data");
        }
    });

    //  SINGLY LIST 
    run_benchmark("SinglyList: PushFront", []() {
        SinglyList list;
        for(int i = 0; i < N; ++i) {
            list.pushFront("data");
        }
    });

    //  DOUBLY LIST 
    run_benchmark("DoublyList: PushBack", []() {
        DoublyList list;
        for(int i = 0; i < N; ++i) {
            list.pushBack("data");
        }
    });

    //  STACK 
    run_benchmark("Stack: Push", []() {
        Stack st;
        for(int i = 0; i < N; ++i) {
            st.push("data");
        }
    });

    //  QUEUE 
    run_benchmark("Queue: Push", []() {
        Queue q;
        for(int i = 0; i < N; ++i) {
            q.push("data");
        }
    });

    //  HASH TABLE 
    run_benchmark("HashTable: Put", []() {
        HashTable ht(N * 2); // Заранее дадим место, чтобы не было слишком много рехеширования
        for(int i = 0; i < N; ++i) {
            ht.put("key" + to_string(i), "val");
        }
    });

    run_benchmark("HashTable: Get", []() {
        HashTable ht(N); 
        // Сначала заполним
        for(int i = 0; i < N; ++i) ht.put("key" + to_string(i), "val");
        // Теперь ищем
        for(int i = 0; i < N; ++i) {
            ht.get("key" + to_string(i));
        }
    });

    //  BINARY TREE 
    // Дерево медленнее, уменьшим N для него
    int TreeN = 10000; 
    run_benchmark("CompleteBinaryTree: Insert (" + to_string(TreeN) + " ops)", [TreeN]() {
        CompleteBinaryTree tree;
        for(int i = 0; i < TreeN; ++i) {
            tree.insert(i);
        }
    });

    return 0;
}