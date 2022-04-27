#include <stdlib.h>
#include <stdio.h>
#include "io.h"
#include "tree.h"

/**
 * Constrains
 * 1. number of nodes in the tree? [0,2000]
 * Ideas and Complexities
 * 1. maximum height of the tree could be 2000
 * 2. maximum number of nodes in one layer could be larger than 2000-1023=977 but less than 1024
 */

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** levelOrder(struct TreeNode* root, int* returnSize, int** returnColumnSizes){
    if (root == NULL) {
        *returnSize = 0;
        return NULL;
    }
    int cap = 1024;
    struct TreeNode* fifo[cap]; int layer[cap];
    fifo[0] = root; layer[0] = 0;
    // empty node not count to node num
    *returnSize = 2000;
    int** res = malloc(sizeof(int*)*(*returnSize));
    *returnColumnSizes = malloc(sizeof(int)*(*returnSize));
    int thisLay=0, lastLayer=-1, layerSize=1, layerIndex=0;
    struct TreeNode* t;
    int i=0,j=1;
    for (;i != j;) {
        t = fifo[i];
        thisLay = layer[i];
        i = (i+1) % cap;
        if (thisLay != lastLayer) {
            (*returnColumnSizes)[thisLay] = layerSize;
            res[thisLay] = malloc(sizeof(int)*layerSize);
            layerSize = 0;
            layerIndex = 0;
            lastLayer = thisLay;
        }
        res[thisLay][layerIndex++] = t->val;
        if (t->left != NULL) {
            fifo[j] = t->left;
            layer[j] = thisLay + 1;
            layerSize++;
            j = (j+1) % cap;
        }
        if (t->right != NULL) {
            fifo[j] = t->right;
            layer[j] = thisLay + 1;
            layerSize++;
            j = (j+1) % cap;
        }
    }
    *returnSize = thisLay+1;
    return res;
}

int** levelOrder_2(struct TreeNode* root, int* returnSize, int** returnColumnSizes){
    if (root == NULL) {
        *returnSize = 0;
        return NULL;
    }
    *returnSize = 2000;
    int** ret = malloc(sizeof(int*)*(*returnSize));
    *returnColumnSizes = malloc(sizeof(int)*(*returnSize));
    int cap = 1000;
    struct TreeNode* queue[cap]; queue[0] = root;
    int i = 0, j = 1, k = 0;
    struct TreeNode* t;
    for (int tj;i!=j;k++) {
        tj = j;
        (*returnColumnSizes)[k] = tj-i;
        ret[k] = malloc(sizeof(int)*(tj-i));
        int index = 0;
        for (;i!=tj;) {
            t = queue[i];
            i = (i+1)%cap;
            ret[k][index++] = t->val;
            if (t->left != NULL) {
                queue[j] = t->left;
                j = (j+1)%cap;
            }
            if (t->right != NULL) {
                queue[j] = t->right;
                j = (j+1)%cap;
            }
        }
    }
    *returnSize = k;
    return ret;
}


int main() {
    int arr[7] = {3,9,20,0,0,15,7};
    struct TreeNode* root = arrTree(arr, sizeof(arr)/sizeof(int), 0);
    int returnSize; int* returnColumnSizes;
    int** ret = levelOrder_2(root, &returnSize, &returnColumnSizes);
    printRet(ret, returnSize, returnColumnSizes);
    freeTree(root);
}
