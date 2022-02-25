#include <stdlib.h>
#include <string.h>
#include <stdio.h>

/*
 * 中等(medium)
 * 数组(array)
 * 回溯(back tracking)
 * 剪枝(branch cut)
 *
 * Constrains
 * 1. candidatesSize? [1,30]
 * 2. candidates[i]? [1,200]
 * 3. target? [1,500]
 * 4. candidates[i] != candidates[j], i != j
 * 5. res[i] == res[j], i != j is allowed
 * 6. returnSize < 150
 *
 * Test Cases
 * 1. [2,3,6,7], 7 => [[2,2,3],[7]]
 *
 * Ideas and Complexities
 * back tracking
 */

void dfs(int* a, int aSize, int target, int cur, int* path, int pathIndex, int** res, int* resIndex, int* returnColumnSizes) {
    if (target == 0) {
        int* tmp = malloc(sizeof(int)*pathIndex);
        memcpy(tmp, path, sizeof(int)*pathIndex);
        res[(*resIndex)] = tmp;
        returnColumnSizes[(*resIndex)] = pathIndex;
        (*resIndex)++;
        return;
    }
    for (int i=cur;i<aSize;i++) {
        if (target<0 && a[i]>=0) {
            continue;
        }
        //如果有负数的话怎么剪枝叶？
        path[pathIndex] = a[i];
        dfs(a, aSize, target-a[i], i, path, pathIndex+1, res, resIndex, returnColumnSizes);
    }
}

int compar(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** combinationSum(int* candidates, int candidatesSize, int target, int* returnSize, int** returnColumnSizes){
    qsort(candidates, candidatesSize, sizeof(int), compar);
    *returnColumnSizes = malloc(sizeof(int) * 150);
    int** res = malloc(sizeof(int*) * 150);
    int resIndex=0;
    int* comb = malloc(sizeof(int) * 500);
    dfs(candidates, candidatesSize, target, 0, comb, 0, res, &resIndex, *returnColumnSizes);
    free(comb);
    *returnSize = resIndex;
    return res;
}

int main() {
    int nums[4] = {2,3,6,7}, len = 4, target = 7;
    int returnSize, *returnColumnSizes;
    int** res = combinationSum(nums, len, target, &returnSize, &returnColumnSizes);
    printf("[");
    for (int i=0;i<returnSize;i++) {
        if (i>0) printf(", ");
        printf("[");
        for (int j=0;j<returnColumnSizes[i];j++) {
            if (j>0) printf(",");
            printf("%d", res[i][j]);
        }
        free(res[i]);
        printf("]");
    }
    printf("]\n");
    free(res);
    free(returnColumnSizes);
    return 0;
}
