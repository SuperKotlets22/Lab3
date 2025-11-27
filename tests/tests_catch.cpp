#define CATCH_CONFIG_MAIN
#include "catch.hpp"
#include "../DynamicArray.h"
#include "../SinglyList.h"
#include "../Stack.h"
#include "../Queue.h"
#include "../HashTable.h"
#include "../CompleteBinaryTree.h"

using namespace std;

TEST_CASE("DynamicArray Full", "[array]") {
    DynamicArray arr;
    arr.pushBack("A");
    REQUIRE(arr.getSize() == 1);
    REQUIRE(arr.get(0) == "A");
    REQUIRE_THROWS_AS(arr.get(10), out_of_range);

    arr.saveToText("arr_catch.txt");
    DynamicArray loaded;
    loaded.loadFromText("arr_catch.txt");
    REQUIRE(loaded.get(0) == "A");
}

TEST_CASE("SinglyList Full", "[list]") {
    SinglyList list;
    list.pushBack("head");
    REQUIRE(list.getHead() == "head");
}

TEST_CASE("Stack Full", "[stack]") {
    Stack st;
    st.push("1");
    REQUIRE(st.pop() == "1");
    REQUIRE(st.pop() == "");
}

TEST_CASE("Queue Full", "[queue]") {
    Queue q(5);
    q.push("Q");
    REQUIRE(q.pop() == "Q");
    REQUIRE(q.pop() == "");
}

TEST_CASE("HashTable Full", "[hash]") {
    HashTable ht;
    ht.put("k", "v");
    REQUIRE(ht.get("k") == "v");
    REQUIRE(ht.get("z") == "");
}

TEST_CASE("Tree Full", "[tree]") {
    CompleteBinaryTree tree;
    tree.insert(50);
    REQUIRE(tree.find(50) == true);
    REQUIRE(tree.find(10) == false);
    
    tree.saveToText("tree_catch.txt");
    CompleteBinaryTree loaded;
    loaded.loadFromText("tree_catch.txt");
    REQUIRE(loaded.find(50) == true);
}