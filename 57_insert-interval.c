#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*
 * 中等(medium)
 * 数组(array)
 *
 * Constrains
 * 1. intervalsSize? [0,10000]
 * 2. intervals[i][j]? [0,100000]
 * 3. intervals[i][0] > intervals[i-1][1] >= intervals[i-1][0]
 *
 * Ideas and Complexities
 * 1. binary
 *  
 *
 */

#define MAX(a,b) ((a) > (b) ? (a) : (b))
#define MIN(a,b) ((a) < (b) ? (a) : (b))

int** insert_better(int** intervals, int intervalsSize, int* intervalsColSize, int* newInterval, int newIntervalSize, int* returnSize, int** returnColumnSizes){
    *returnSize = intervalsSize + 1;
    int** res = malloc(sizeof(int[2])*(*returnSize));
    *returnColumnSizes = malloc(sizeof(int)*(*returnSize));
    for (int i=0; i<*returnSize; i++) {
        (*returnColumnSizes)[i] = 2;
    }
    int i=0, k=0;
    /* for (;i<intervalsSize && intervals[i][1] < newInterval[0]; i++) { */
    /*     res[k++] = intervals[i]; */
    /* } */
    for (;i<intervalsSize && intervals[i][1] < newInterval[0]; i++) {
    }
    k = i;
    memcpy(res, intervals, k*sizeof(int[2]));
    for (;i<intervalsSize && intervals[i][0] <= newInterval[1]; i++) {
        newInterval[0] = MIN(newInterval[0], intervals[i][0]);
        newInterval[1] = MAX(newInterval[1], intervals[i][1]);
    }
    res[k++] = newInterval;
    if (i<intervalsSize) {
        memcpy(res+k, intervals+i, (intervalsSize-i)*sizeof(int[2]));
    }
    /* for (;i<intervalsSize; i++){ */
    /*     res[k++] = intervals[i]; */
    /* } */
    *returnSize = k;
    return res;
}

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** insert(int** intervals, int intervalsSize, int* intervalsColSize, int* newInterval, int newIntervalSize, int* returnSize, int** returnColumnSizes){
    *returnSize = intervalsSize + 1;
    int** res = malloc(sizeof(int[2])*(*returnSize));
    *returnColumnSizes = malloc(sizeof(int)*(*returnSize));
    for (int i=0; i<*returnSize; i++) {
        (*returnColumnSizes)[i] = 2;
    }
    /* if (intervalsSize == 0) { */
    /*     *returnSize = 1; */
    /*     res[0] = newInterval; */
    /*     return res; */
    /* } */
    for (int i=0; i<intervalsSize; i++) {
        // 区间是有序的，这个左边不重叠，则后面所有必然不重叠
        if (intervals[i][0] > newInterval[1]) break;
        // 这个右边不重叠，后面的可能重叠
        else if (intervals[i][1] < newInterval[0]) continue;
        // 有重叠，直接合并
        else {
            newInterval[0] = MIN(newInterval[0], intervals[i][0]);
            newInterval[1] = MAX(newInterval[1], intervals[i][1]);
        }
    }
    int k=0, flag=0;
    for (int j=0; j<intervalsSize; j++) {
        // 什么时候放入合并区间？找到合并区间右边那个区间时。
        if (!flag && intervals[j][0] > newInterval[1]) {
            res[k++] = newInterval;
            flag = 1;
        }
        // 当前区间不在合并区间内
        if (intervals[j][1] < newInterval[0] || intervals[j][0] > newInterval[1]) res[k++] = intervals[j];
    }
    // 合并区间始终没被放入
    if (!flag) res[k++] = newInterval;
    *returnSize = k;
    return res;
}

int main() {
    int* intervals[2];
    intervals[0] = (int[2]){3,5};
    intervals[1] = (int[2]){12,15};
    int intervalsSize = 2;
    int intervalsColSize[5] = {2,2,2,2,2};
    int newInterval[2] = {6,6};
    int returnSize, *returnColumnSizes;
    int** res = insert(intervals, intervalsSize, intervalsColSize, newInterval, 2, &returnSize, &returnColumnSizes);
    printf("[");
    for (int i=0; i<returnSize; i++) {
        if (i>0) printf(", ");
        printf("[");
        for (int j=0; j<returnColumnSizes[i]; j++) {
            if (j>0) printf(",");
            printf("%d", res[i][j]);
        }
        printf("]");
    }
    printf("]\n");
}
