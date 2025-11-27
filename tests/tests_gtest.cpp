#include <gtest/gtest.h>
#include "../DynamicArray.h"
#include "../SinglyList.h"
#include "../DoublyList.h"
#include "../Stack.h"
#include "../Queue.h"
#include "../HashTable.h"
#include "../CompleteBinaryTree.h"

//  DYNAMIC ARRAY 
TEST(DynamicArrayTest, FullCoverage) {
    DynamicArray arr;
    // Тест расширения (resize)
    for(int i=0; i<20; ++i) arr.pushBack(to_string(i)); 
    EXPECT_EQ(arr.getSize(), 20);
    EXPECT_EQ(arr.get(0), "0");
    EXPECT_EQ(arr.get(19), "19");

    // Тест ошибок
    EXPECT_THROW(arr.get(-1), out_of_range);
    EXPECT_THROW(arr.get(100), out_of_range);

    // Тест сериализации (текст)
    arr.saveToText("arr.txt");
    DynamicArray loaded;
    loaded.loadFromText("arr.txt");
    EXPECT_EQ(loaded.getSize(), 20);
    EXPECT_EQ(loaded.get(5), "5");

    // Тест сериализации (бинар)
    arr.saveToBinary("arr.bin");
    DynamicArray loadedBin;
    loadedBin.loadFromBinary("arr.bin");
    EXPECT_EQ(loadedBin.getSize(), 20);
    EXPECT_EQ(loadedBin.get(5), "5");

    // Тест ошибки открытия файла
    DynamicArray bad;
    bad.loadFromText("non_existent.txt"); // Должен просто выйти
    EXPECT_EQ(bad.getSize(), 0);
    bad.loadFromBinary("non_existent.bin");
    EXPECT_EQ(bad.getSize(), 0);
}

//  SINGLY LIST 
TEST(SinglyListTest, FullCoverage) {
    SinglyList list;
    list.pushBack("val1");
    list.pushBack("val2");
    EXPECT_EQ(list.getHead(), "val1");

    // Сериализация
    list.saveToText("slist.txt");
    SinglyList loaded;
    loaded.loadFromText("slist.txt");
    EXPECT_EQ(loaded.getHead(), "val1");

    list.saveToBinary("slist.bin");
    SinglyList loadedBin;
    loadedBin.saveToBinary("slist.bin"); // просто вызов сохранения
}

//  DOUBLY LIST 
TEST(DoublyListTest, FullCoverage) {
    DoublyList list;
    list.pushBack("dval1");
    list.pushBack("dval2");
    EXPECT_EQ(list.getTail(), "dval2");
    
    // Проверка пустого
    DoublyList empty;
    EXPECT_EQ(empty.getTail(), "");
}

//  STACK 
TEST(StackTest, FullCoverage) {
    Stack st;
    st.push("s1");
    st.push("s2");
    EXPECT_EQ(st.pop(), "s2");
    EXPECT_EQ(st.pop(), "s1");
    EXPECT_EQ(st.pop(), ""); // пустой стек

    // Сериализация
    st.push("save_me");
    st.saveToText("stack.txt");
    Stack loaded;
    loaded.loadFromText("stack.txt");
    EXPECT_EQ(loaded.pop(), "save_me");
}

//  QUEUE 
TEST(QueueTest, FullCoverage) {
    Queue q(3); // Маленький размер для теста кольца
    q.push("1");
    q.push("2");
    q.push("3");
    q.push("4"); // переполнение (игнор или перезапись в зависимости от реализации)

    EXPECT_EQ(q.pop(), "1");
    
    // Сериализация
    q.saveToText("queue.txt");
    Queue loaded;
    loaded.loadFromText("queue.txt");
    EXPECT_NE(loaded.pop(), ""); // Проверяем что что-то загрузилось

    // Пустая очередь
    Queue empty;
    EXPECT_EQ(empty.pop(), "");
}

//  HASH TABLE 
TEST(HashTableTest, FullCoverage) {
    HashTable ht(5); // Мало места -> будут коллизии
    ht.put("a", "1");
    ht.put("f", "2"); // 'a' и 'f' могут дать коллизию при % 5
    ht.put("k", "3");

    EXPECT_EQ(ht.get("a"), "1");
    EXPECT_EQ(ht.get("f"), "2");
    EXPECT_EQ(ht.get("z"), ""); // Нет такого

    // Перезапись
    ht.put("a", "10");
    EXPECT_EQ(ht.get("a"), "10");

    // Сериализация
    ht.saveToText("hash.txt");
    HashTable loaded(5);
    loaded.loadFromText("hash.txt");
    EXPECT_EQ(loaded.get("a"), "10");
}

//  COMPLETE BINARY TREE 
TEST(TreeTest, FullCoverage) {
    CompleteBinaryTree tree;
    tree.insert(10);
    tree.insert(20);
    tree.insert(30);

    EXPECT_TRUE(tree.find(10));
    EXPECT_TRUE(tree.find(20));
    EXPECT_TRUE(tree.find(30));
    EXPECT_FALSE(tree.find(99));

    // Сериализация
    tree.saveToText("tree.txt");
    CompleteBinaryTree loaded;
    loaded.loadFromText("tree.txt");
    EXPECT_TRUE(loaded.find(20));

    // Пустое дерево
    CompleteBinaryTree empty;
    EXPECT_FALSE(empty.find(1));
    empty.saveToText("tree_empty.txt");
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}