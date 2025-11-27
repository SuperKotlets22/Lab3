#define BOOST_TEST_MODULE Lab3Tests
#include <boost/test/included/unit_test.hpp>
#include "../DynamicArray.h"
#include "../SinglyList.h"
#include "../Stack.h"
#include "../Queue.h"
#include "../HashTable.h"
#include "../CompleteBinaryTree.h"

using namespace std;

BOOST_AUTO_TEST_CASE(DynamicArray_Full) {
    DynamicArray arr;
    for(int i=0; i<10; ++i) arr.pushBack(to_string(i));
    BOOST_CHECK_EQUAL(arr.getSize(), 10);
    BOOST_CHECK_EQUAL(arr.get(0), "0");
    BOOST_CHECK_THROW(arr.get(-1), out_of_range);

    // Сериализация
    arr.saveToText("arr_boost.txt");
    DynamicArray loaded;
    loaded.loadFromText("arr_boost.txt");
    BOOST_CHECK_EQUAL(loaded.getSize(), 10);
}

BOOST_AUTO_TEST_CASE(SinglyList_Full) {
    SinglyList list;
    list.pushBack("val1");
    BOOST_CHECK_EQUAL(list.getHead(), "val1");
    
    list.saveToText("list_boost.txt");
    SinglyList loaded;
    loaded.loadFromText("list_boost.txt");
    BOOST_CHECK_EQUAL(loaded.getHead(), "val1");
}

BOOST_AUTO_TEST_CASE(Stack_Full) {
    Stack st;
    st.push("A");
    st.push("B");
    BOOST_CHECK_EQUAL(st.pop(), "B");
    BOOST_CHECK_EQUAL(st.pop(), "A");
    BOOST_CHECK_EQUAL(st.pop(), "");
}

BOOST_AUTO_TEST_CASE(Queue_Full) {
    Queue q(5);
    q.push("1");
    q.push("2");
    BOOST_CHECK_EQUAL(q.pop(), "1");
    BOOST_CHECK_EQUAL(q.pop(), "2");
    BOOST_CHECK_EQUAL(q.pop(), "");
}

BOOST_AUTO_TEST_CASE(HashTable_Full) {
    HashTable ht;
    ht.put("key", "value");
    BOOST_CHECK_EQUAL(ht.get("key"), "value");
    BOOST_CHECK_EQUAL(ht.get("no"), "");
}

BOOST_AUTO_TEST_CASE(Tree_Full) {
    CompleteBinaryTree tree;
    tree.insert(100);
    tree.insert(200);
    BOOST_CHECK(tree.find(100));
    BOOST_CHECK(!tree.find(500));
    
    tree.saveToText("tree_boost.txt");
    CompleteBinaryTree loaded;
    loaded.loadFromText("tree_boost.txt");
    BOOST_CHECK(loaded.find(200));
}