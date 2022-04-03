#include <stdlib.h>
#include <stdio.h>
#include <string.h>

/*
 * 中等(medium)
 * 数组(array)
 * 双指针(double pointer)
 * 回溯(backtracking)
 * 深度优先(dfs)
 */

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */

int comp(const void *a, const void *b) {
    return *(int*)a - *(int*)b;
}

typedef struct Node { int* v; struct Node *next; } node;

/*
 * n n sum
 * a array to search
 * l, r array index range
 * d depth
 * v value to reach
 * p path
 * t hold result
 */
void n_sum_dfs(int d, int *p, int n, int* a, int l, int r, int v, node *t) {
    if ( d < 2 ) {
        return;
    }
    if ( d == 2 ) {
        for (int ll=l, rr=r; ll < rr;) {
            if (v == a[ll] + a[rr]) {
                p[0] = a[rr];
                p[1] = a[ll];
                node *tt = malloc(sizeof(node));
                int* vv = malloc(sizeof(int)*n);
                memcpy(vv, p, sizeof(int) * n);
                tt->v = vv;
                tt->next = t->next;
                t->next = tt;
                t = tt;
                /* todo: check index before value */
                for (rr--; a[rr] == a[rr+1] && ll < rr; rr--);
                for (ll++; a[ll] == a[ll-1] && ll < rr; ll++);
            } else if (v < a[ll] + a[rr]) {
                /* todo: check index before value */
                for (rr--; a[rr] == a[rr+1] && ll < rr; rr--);
            } else {
                /* todo: check index before value */
                for (ll++; a[ll] == a[ll-1] && ll < rr; ll++);
            }
        }
        return;
    }
    for (int i=l; i <= r-d+1; i++) {
        if ( i>l && a[i] == a[i-1] ) continue;
        {
            long vv;
            int dd, j;
            for (vv=v, dd=d, j=i; dd>0; dd--, j++) {
                vv -= a[j];
                // sum is too large while num still positive.
                // plus a[j+1] will only make sum larger.
                if ( a[j] >= 0 && vv < 0 ) {
                    return;
                }
                // sum of minimum dd nums is too large
                if ( dd == 1 && vv < 0 ) {
                    return;
                }
            }
            for (vv=v, dd=d, j=r; dd>0; dd--, j--) {
                vv -= a[j];
                // sum is too small while num already negative.
                // plus a[j-1] will only make sum smaller.
                if ( a[j] <= 0 && vv > 0 ) {
                    return;
                }
                // sum of maximum dd nums is too small
                if ( dd == 1 && vv > 0 ) {
                    return;
                }
            }
        }
        p[d-1] = a[i];
        n_sum_dfs(d-1, p, n, a, i+1, r, v-a[i], t);
    }
}
int** nSum(int n, int* nums, int numsSize, int target, int* returnSize, int** returnColumnSizes){
    qsort(nums, numsSize, sizeof(int), comp);
    node *h, *z, *t;
    h = malloc(sizeof(*h)); z = malloc(sizeof(*z)); h->next = z; z->next = z; t = h;

    int* p = malloc(sizeof(int)*n);
    n_sum_dfs(n, p, n, nums, 0, numsSize-1, target, t);
    free(p);

    int i=0, size=0;
    for (t = h->next; t != z; t = t->next) size++;
    *returnSize = size;
    int** res = malloc(sizeof(int*) * size);
    *returnColumnSizes = malloc(sizeof(int) * size);
    if (size > 0) {
        for (i = 0, t = h->next; t != z; t = t->next, i++) {
            res[i] = malloc(sizeof(int) * n);
            memcpy(res[i], t->v, sizeof(int) * n);
            (*returnColumnSizes)[i] = n;
        }
    }
    z->next = h->next;
    free(h);
    h = z->next;
    for (t = h, h = h->next; t != z; t = h, h = h->next) {
        free(t->v);
        free(t);
    }
    free(z);
    return res;
}

int** fourSum(int* nums, int numsSize, int target, int* returnSize, int** returnColumnSizes){
    return nSum(4, nums, numsSize, target, returnSize, returnColumnSizes);
}

int main() {
    int returnSize;
    int* returnColumnSizes;
    /* int nums[6] = {0, 0, 1, 2, 3, 4}; */
    /* int nums[6] = {1, 0, -1, 0, -2, 2}; */
    /* int nums[5] = {2, 2,2,2,2}; */
    int nums[8] = {-3,-2,-1,0,0,1,2,3}, target = 0;
    /* int nums[4] = {1000000000,1000000000,1000000000,1000000000}, target = 0; */
    int** res = nSum(4, nums, 8, target, &returnSize, &returnColumnSizes);
    printf("[");
    for (int i=0; i < returnSize; i++) {
        if (i>0) printf(", ");
        printf("[");
        for (int j=0; j < returnColumnSizes[i]; j++) {
            if (j>0) printf(",");
            printf("%d", res[i][j]);
        }
        free(res[i]);
        printf("]");
    }
    printf("]\n");
    free(returnColumnSizes);
    free(res);
    return 0;
}
