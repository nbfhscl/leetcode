#include <stdlib.h>
#include <stdio.h>


/**
 * 中等(medium)
 * 数组(array)
 * 排序(sort)
 *
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */

#define SWAP(a, b) \
({ \
    __auto_type _t = (a); \
    (a) = (b); \
    (b) = _t; \
})

#define MAX(a,b) ((a) > (b) ? (a) : (b))

void quick_sort(int** a, int l, int r) {
    if (l>=r) return;
    if (r-l < 12) {
        for (int i=l+1; i<=r; i++) {
            int j=i, *k = a[j];
            for (;--j>=0 && a[j][0]>k[0];) {
                 a[j+1] = a[j];
            }
            a[++j] = k;
        }
        return;
    }
    int mid = l + ((r-l)>>1);
    if (a[l][0] > a[r][0]) {
        SWAP(a[l], a[r]);
    }
    if (a[mid][0] > a[r][0]) {
        SWAP(a[mid], a[r]);
    }
    if (a[l][0] > a[mid][0]) {
        SWAP(a[l], a[mid]);
    }
    SWAP(a[mid], a[r-1]);
    int ll=l+1,rr=r-1;
    int *k=a[rr];
    for (;ll<rr;) {
        for(;ll<rr && a[ll][0]<=k[0];) ll++;
        a[rr] = a[ll];
        for(;ll<rr && a[rr][0]>=k[0];) rr--;
        a[ll] = a[rr];
    }
    a[rr] = k;
    quick_sort(a, l, rr-1);
    quick_sort(a, rr+1, r);
}

int compare(const void* a, const void* b) {
    // a is pointer to array, array is int*, so a is int**
    return (*(int**)a)[0] - (*(int**)b)[0];
}

int** merge(int** intervals, int intervalsSize, int* intervalsColSize, int* returnSize, int** returnColumnSizes){
    qsort(intervals, intervalsSize, sizeof(*intervals), compare);
    /* quick_sort(intervals, 0, intervalsSize-1); */
    int j=0;
    for (int i=1; i<intervalsSize; i++) {
        if (intervals[i][0] > intervals[j][1]) {
            j++;
            intervals[j][0] = intervals[i][0];
            intervals[j][1] = intervals[i][1];
        } else {
            intervals[j][1] = MAX(intervals[j][1], intervals[i][1]);
        }
    }
    *returnSize = j+1;
    *returnColumnSizes = malloc(sizeof(int)*(*returnSize));
    for (int i=0; i<(*returnSize); i++) {
        (*returnColumnSizes)[i] = 2;
    }
    return intervals;
}

int main() {
    int *intervals[2];
    intervals[0] = (int[2]){39,43};
    intervals[1] = (int[2]){43,49};
    int intervalsSize=2, intervalsColSize[2] = {2,2};
    int returnSize, *returnColumnSizes;
    int **res = merge(intervals, intervalsSize, intervalsColSize, &returnSize, &returnColumnSizes);
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

