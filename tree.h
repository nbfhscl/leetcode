#include <stddef.h>
#include <stdlib.h>

// Definition for a binary tree node.
struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

struct TreeNode* arrTree(int* array, int size, int i);
void freeTree(struct TreeNode* node);
