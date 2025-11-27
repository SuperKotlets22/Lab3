#include <iostream>
#include <chrono>
#include "../DynamicArray.h"

using namespace std;

int main() {
    DynamicArray arr;
    auto start = chrono::high_resolution_clock::now();
    
    // бенчмарк добавления
    for(int i=0; i<100000; i++) {
        arr.pushBack("bench");
    }
    
    auto end = chrono::high_resolution_clock::now();
    chrono::duration<double> diff = end - start;
    
    cout << "добавление 100000 элементов заняло: " << diff.count() << " сек." << endl;
    return 0;
}