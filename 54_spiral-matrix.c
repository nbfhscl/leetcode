#include <stdlib.h>

/* 中等(medium)
 * 数组(array)
 * 矩阵(matrix)
 *
 * Ideas and Complexities
 * 1. right -> down -> left -> up, print outer round, then inner round. Time
 *    O(mn), Memory O(1)
 * 2. print first line, rotate 90 degrees counterclockwise, print first line
 */


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* spiralOrder(int** matrix, int matrixSize, int* matrixColSize, int* returnSize){
    *returnSize = matrixSize * matrixColSize[0];
    int *res = malloc(sizeof(int)*(*returnSize));
    int i=-1,j=-1,k=-1;
    for (int round=0;;round++) {
        for (++i;++j<matrixColSize[i]-round;) {
            res[++k] = matrix[i][j];
        }
        if (k>=*returnSize-1) break;
        for (--j;++i<matrixSize-round;) {
            res[++k] = matrix[i][j];
        }
        if (k>=*returnSize-1) break;
        for (--i;--j>=0+round;) {
            res[++k] = matrix[i][j];
        }
        if (k>=*returnSize-1) break;
        for (++j;--i>=1+round;) {
            res[++k] = matrix[i][j];
        }
        if (k>=*returnSize-1) break;
    }
    return res;
}

int main() {
    int *matrix[3];
    matrix[0] = (int[4]){1,2,3,4};
    matrix[1] = (int[4]){5,6,7,8};
    matrix[2] = (int[4]){9,10,11,12};
    int matrixColSize[] = {4,4,4};
    int returnSize;
    spiralOrder(matrix, 3, matrixColSize, &returnSize);
}
