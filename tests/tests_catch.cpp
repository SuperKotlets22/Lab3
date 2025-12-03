#define CATCH_CONFIG_MAIN
#include "catch.hpp" 

#include "DynamicArray.h"
#include "SinglyList.h"
#include "DoublyList.h"
#include "Stack.h"
#include "Queue.h"
#include "HashTable.h"
#include "CompleteBinaryTree.h"

using namespace std;

TEST_CASE("DynamicArray functionality", "[Array]") {
    DynamicArray arr;
    
    SECTION("Push and Get") {
        arr.pushBack("Val1");
        arr.pushFront("Val0"); 
        REQUIRE(arr.getSize() == 2);
        REQUIRE(arr.get(0) == "Val0");
        REQUIRE(arr.get(1) == "Val1");
    }
    
    SECTION("Exceptions") {
        REQUIRE_THROWS_AS(arr.get(5), out_of_range);
        REQUIRE_THROWS_AS(arr.removeAt(-1), out_of_range);
    }
    
    SECTION("Serialization") {
        arr.pushBack("SaveMe");
        arr.saveToText("catch_arr.txt");
        
        DynamicArray loaded;
        loaded.loadFromText("catch_arr.txt");
        REQUIRE(loaded.getSize() == 1);
        REQUIRE(loaded.get(0) == "SaveMe");
    }
}

TEST_CASE("SinglyList functionality", "[SinglyList]") {
    SinglyList list;
    list.pushBack("A");
    list.pushBack("B");
    
    REQUIRE(list.getSize() == 2);
    REQUIRE(list.getHead() == "A");
    
    list.popFront();
    REQUIRE(list.getHead() == "B");
}

TEST_CASE("DoublyList functionality", "[DoublyList]") {
    DoublyList list;
    list.pushFront("First");
    list.pushBack("Last");
    
    REQUIRE(list.getTail() == "Last");
    
    list.popBack();
    REQUIRE(list.getTail() == "First");
}

TEST_CASE("Stack functionality", "[Stack]") {
    Stack st;
    st.push("10");
    st.push("20");
    
    REQUIRE(st.pop() == "20");
    REQUIRE(st.peek() == "10");
    REQUIRE(st.pop() == "10");
    REQUIRE(st.pop() == "");
}

TEST_CASE("Queue functionality", "[Queue]") {
    Queue q;
    q.push("alpha");
    q.push("beta");
    
    REQUIRE(q.pop() == "alpha");
    
    q.saveToBinary("catch_queue.bin");
    Queue loaded;
    loaded.loadFromBinary("catch_queue.bin");
    REQUIRE(loaded.pop() == "beta");
}

TEST_CASE("HashTable functionality", "[HashTable]") {
    HashTable ht(5);
    ht.put("hello", "world");
    ht.put("foo", "bar");
    
    REQUIRE(ht.get("hello") == "world");
    REQUIRE(ht.get("foo") == "bar");
    
    ht.remove("hello");
    REQUIRE(ht.get("hello") == "");
}

TEST_CASE("CompleteBinaryTree functionality", "[Tree]") {
    CompleteBinaryTree tree;
    tree.insert(1);
    tree.insert(2);
    tree.insert(3);
    
    REQUIRE(tree.find(3) == true);
    REQUIRE(tree.find(4) == false);
    
    tree.saveToText("catch_tree.txt");
    CompleteBinaryTree loaded;
    loaded.loadFromText("catch_tree.txt");
    REQUIRE(loaded.find(2) == true);
}