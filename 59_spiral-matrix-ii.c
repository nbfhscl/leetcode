#include <stdlib.h>
#include <string.h>

/*
 * 中等(medium)
 * 数组(array)
 *
 * Constrains
 * 1. n? [1,20]
 *
 * Ideas and Complexities
 * outer circle -> inner circle, write one by one.
 */

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** generateMatrix(int n, int* returnSize, int** returnColumnSizes){
    *returnSize = n;
    *returnColumnSizes = malloc(sizeof(int)*n);
    memset(*returnColumnSizes, n, sizeof(int));
    /* for (int i=0; i<n; i++) { */
    /*     (*returnColumnSizes)[i] = n; */
    /* } */
    int** res = malloc(sizeof(int*)*n);
    for (int i=0; i<n; i++) {
        res[i] = malloc(sizeof(int)*n);
    }
    int n2 = n*n;
    for (int round = 0, k=0, i=-1, j=-1;; round++) {
        for (++i; ++j<n-round;) {
            res[i][j] = ++k;
        }
        if (k >= n2) break;
        for (--j; ++i<n-round;) {
            res[i][j] = ++k;
        }
        if (k >= n2) break;
        for (--i; --j>=0+round;) {
            res[i][j] = ++k;
        }
        if (k >= n2) break;
        for (++j; --i>0+round;) {
            res[i][j] = ++k;
        }
        if (k >= n2) break;
    }
    return res;
}



