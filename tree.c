#include "tree.h"

struct TreeNode* arrTree(int* array, int size, int i) {
    // array val 0 used as null node
    if (i>=size || array[i] == 0) {
        return NULL;
    }
    struct TreeNode* node = malloc(sizeof(struct TreeNode));
    node->val = array[i];
    node->left = arrTree(array, size, 2*i+1);
    node->right = arrTree(array, size, 2*i+2);
    return node;
}

void freeTree(struct TreeNode* node) {
    if (node == NULL) return;
    freeTree(node->left);
    freeTree(node->right);
    free(node);
}

