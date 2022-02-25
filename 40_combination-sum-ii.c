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
 * 1. candidatesSize? [1,100]
 * 2. candidates[i]? [1,50]
 * 3. target? [1,30]
 * 4. res[i] != res[j], i != j
 * 5. candidates[i] can equal to candidates[j], i != j
 */

int compar(const void *a, const void *b) {
    return *(int*)a - *(int*)b;
}

void dfs(int *a, int aSize, int cur, int *path, int pathIndex, int **res, int
        *resIndex, int *returnColumnSizes, int target) {
    if (target == 0) {
        int *tmp = malloc(sizeof(int)*pathIndex);
        memcpy(tmp, path, sizeof(int)*pathIndex);
        res[*resIndex] = tmp;
        returnColumnSizes[*resIndex] = pathIndex;
        (*resIndex)++;
        return;
    }
    for (int i=cur; i<aSize; i++) {
        if (i>cur && a[i] == a[i-1]) continue;
        if (target<0 && a[i]>=0) return;
        path[pathIndex] = a[i];
        dfs(a, aSize, i+1, path, pathIndex+1, res, resIndex, returnColumnSizes,
                target-a[i]);
    }
}

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** combinationSum2(int* candidates, int candidatesSize, int target, int* returnSize, int** returnColumnSizes){
    qsort(candidates, candidatesSize, sizeof(int), compar);
    *returnColumnSizes = malloc(sizeof(int)*500);
    int **res = malloc(sizeof(int*)*500);
    int *path = malloc(sizeof(int)*100);
    *returnSize = 0;
    dfs(candidates, candidatesSize, 0, path, 0, res, returnSize, *returnColumnSizes,
            target);
    free(path);
    return res;
}

int main() {
    int nums[7] = {10,1,2,7,6,1,5}, len = 7, target = 8;
    int returnSize, *returnColumnSizes;
    int** res = combinationSum2(nums, len, target, &returnSize, &returnColumnSizes);
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

