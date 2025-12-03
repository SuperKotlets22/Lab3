#ifndef COMPLETEBINARYTREE_H
#define COMPLETEBINARYTREE_H

#include <iostream>
#include <string>
#include <queue>
#include <fstream>

using namespace std;

struct TreeNode {
    int data;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int val) : data(val), left(nullptr), right(nullptr) {}
};

class CompleteBinaryTree {
private:
    TreeNode* root;
    int size;

    void deleteTree(TreeNode* node) {
        if (!node) return;
        deleteTree(node->left);
        deleteTree(node->right);
        delete node;
    }
    
    
    void serializeBinaryRecursive(TreeNode* node, ofstream& file) {
        
    }

public:
    CompleteBinaryTree() : root(nullptr), size(0) {}

    ~CompleteBinaryTree() {
        deleteTree(root);
    }

    void insert(int value) {
        TreeNode* newNode = new TreeNode(value);
        if (!root) {
            root = newNode;
            size++;
            return;
        }

        queue<TreeNode*> q;
        q.push(root);

        while (!q.empty()) {
            TreeNode* curr = q.front();
            q.pop();

            if (!curr->left) {
                curr->left = newNode;
                size++;
                return;
            } else {
                q.push(curr->left);
            }

            if (!curr->right) {
                curr->right = newNode;
                size++;
                return;
            } else {
                q.push(curr->right);
            }
        }
    }

    bool find(int value) const {
        if (!root) return false;
        queue<TreeNode*> q;
        q.push(root);
        while (!q.empty()) {
            TreeNode* curr = q.front();
            q.pop();
            if (curr->data == value) return true;
            if (curr->left) q.push(curr->left);
            if (curr->right) q.push(curr->right);
        }
        return false;
    }

    bool isComplete() const {
        if (!root) return true;
        queue<TreeNode*> q;
        q.push(root);
        bool flag = false; 
        while(!q.empty()){
            TreeNode* temp = q.front();
            q.pop();
            if(temp->left){
                if(flag) return false;
                q.push(temp->left);
            } else {
                flag = true;
            }
            if(temp->right){
                if(flag) return false;
                q.push(temp->right);
            } else {
                flag = true;
            }
        }
        return true;
    }
    
    void saveToText(const string& filename) {
        ofstream file(filename);
        if(!root) return;
        queue<TreeNode*> q;
        q.push(root);
        while(!q.empty()) {
            TreeNode* temp = q.front();
            q.pop();
            file << temp->data << endl;
            if(temp->left) q.push(temp->left);
            if(temp->right) q.push(temp->right);
        }
        file.close();
    }

    void loadFromText(const string& filename) {
        deleteTree(root);
        root = nullptr;
        size = 0;
        ifstream file(filename);
        int val;
        while(file >> val) {
            insert(val);
        }
        file.close();
    }
};

#endif