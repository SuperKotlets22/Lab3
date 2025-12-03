#define BOOST_TEST_MODULE Lab3_Tests
#include <boost/test/included/unit_test.hpp>

#include "DynamicArray.h"
#include "SinglyList.h"
#include "DoublyList.h"
#include "Stack.h"
#include "Queue.h"
#include "HashTable.h"
#include "CompleteBinaryTree.h"

using namespace std;

//  DYNAMIC ARRAY 
BOOST_AUTO_TEST_CASE(DynamicArray_Ops) {
    DynamicArray arr;
    arr.pushBack("A");
    arr.pushBack("B");
    
    BOOST_CHECK_EQUAL(arr.getSize(), 2);
    BOOST_CHECK_EQUAL(arr.get(1), "B");
    
    arr.set(0, "Z");
    BOOST_CHECK_EQUAL(arr.get(0), "Z");
    
    BOOST_CHECK_THROW(arr.get(10), out_of_range);
    
    // Сериализация
    arr.saveToText("boost_arr.txt");
    DynamicArray loaded;
    loaded.loadFromText("boost_arr.txt");
    BOOST_CHECK_EQUAL(loaded.getSize(), 2);
    BOOST_CHECK_EQUAL(loaded.get(0), "Z");
}

//  SINGLY LIST 
BOOST_AUTO_TEST_CASE(SinglyList_Ops) {
    SinglyList list;
    list.pushBack("1");
    list.pushBack("2");
    
    BOOST_CHECK_EQUAL(list.getSize(), 2);
    
    list.popFront();
    BOOST_CHECK_EQUAL(list.getHead(), "2");
    
    list.clear();
    BOOST_CHECK_EQUAL(list.getSize(), 0);
}

//  DOUBLY LIST 
BOOST_AUTO_TEST_CASE(DoublyList_Ops) {
    DoublyList list;
    list.pushFront("Head");
    list.pushBack("Tail");
    
    BOOST_CHECK_EQUAL(list.getTail(), "Tail");
    
    list.popBack();
    BOOST_CHECK_EQUAL(list.getTail(), "Head");
}

//  STACK 
BOOST_AUTO_TEST_CASE(Stack_Ops) {
    Stack st;
    st.push("data");
    BOOST_CHECK_EQUAL(st.peek(), "data");
    
    string val = st.pop();
    BOOST_CHECK_EQUAL(val, "data");
    BOOST_CHECK_EQUAL(st.getSize(), 0);
}

//  QUEUE 
BOOST_AUTO_TEST_CASE(Queue_Ops) {
    Queue q;
    q.push("one");
    q.push("two");
    
    BOOST_CHECK_EQUAL(q.pop(), "one");
    BOOST_CHECK_EQUAL(q.peek(), "two");
    
    q.saveToBinary("boost_queue.bin");
    Queue loaded;
    loaded.loadFromBinary("boost_queue.bin");
    BOOST_CHECK_EQUAL(loaded.peek(), "two");
}

//  HASH TABLE 
BOOST_AUTO_TEST_CASE(HashTable_Ops) {
    HashTable ht;
    ht.put("user", "admin");
    ht.put("pass", "1234");
    
    BOOST_CHECK_EQUAL(ht.get("user"), "admin");
    
    ht.remove("user");
    BOOST_CHECK_EQUAL(ht.get("user"), "");
}

//  TREE 
BOOST_AUTO_TEST_CASE(Tree_Ops) {
    CompleteBinaryTree tree;
    tree.insert(50);
    tree.insert(25);
    tree.insert(75);
    
    BOOST_CHECK(tree.find(75));
    BOOST_CHECK(!tree.find(100));
    
    tree.saveToText("boost_tree.txt");
    CompleteBinaryTree loaded;
    loaded.loadFromText("boost_tree.txt");
    BOOST_CHECK(loaded.find(25));
}