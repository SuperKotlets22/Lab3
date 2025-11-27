#pragma once
#include <iostream>
#include <queue>
#include <fstream>

using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int v) : val(v), left(nullptr), right(nullptr) {}
};

class CompleteBinaryTree {
private:
    TreeNode* root;
    int size;

    void deleteTree(TreeNode* node) {
        if(!node) return;
        deleteTree(node->left);
        deleteTree(node->right);
        delete node;
    }

public:
    CompleteBinaryTree() : root(nullptr), size(0) {}
    ~CompleteBinaryTree() { deleteTree(root); }

    // вставка в первое свободное место (bfs)
    void insert(int val) {
        TreeNode* newNode = new TreeNode(val);
        if (!root) {
            root = newNode;
            size++;
            return;
        }
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty()) {
            TreeNode* curr = q.front();
            q.pop();
            if(!curr->left) {
                curr->left = newNode;
                size++;
                return;
            } else q.push(curr->left);
            
            if(!curr->right) {
                curr->right = newNode;
                size++;
                return;
            } else q.push(curr->right);
        }
    }
    
    bool find(int val) {
        if (!root) return false;
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty()){
            TreeNode* curr = q.front();
            q.pop();
            if(curr->val == val) return true;
            if(curr->left) q.push(curr->left);
            if(curr->right) q.push(curr->right);
        }
        return false;
    }

    void saveToText(string filename) {
        ofstream file(filename);
        file << size << endl;
        if(!root) return;
        
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty()){
            TreeNode* curr = q.front();
            q.pop();
            file << curr->val << " "; // пишем в строку через пробел
            if(curr->left) q.push(curr->left);
            if(curr->right) q.push(curr->right);
        }
        file.close();
    }
    
    void loadFromText(string filename) {
        ifstream file(filename);
        int newSize;
        file >> newSize;
        
        deleteTree(root);
        root = nullptr;
        size = 0;
        
        int val;
        for(int i=0; i<newSize; i++) {
            file >> val;
            insert(val);
        }
        file.close();
    }
};