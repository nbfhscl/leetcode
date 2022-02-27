#include <stdlib.h>

/**
 * Constrains
 * 1. How large can n be? 0 <= n <= 3000
 * 2. How large can value in array be? -10^5 <= nums[i] <= 10^5
 * 3. Is [0,0,0] allowed?
 *
 * Ideas and Complexities
 * 0. First of all, sort the array. O(nlogn) O（logn)
 * 1. Triple cycle. O(n^3), O(1)
 * 2. Replace the second and the third cycle with Two-Num-Sum solution using
 *    hashmap. O(n^2), O(n)
 * 3. Because the array is sorted, we can use two pointer instead of hashmap.
 *    O(n^2), O(1)
 * best complexity total: O(n^2), O(logn)
 *
 * Test Cases
 * 1. [0, 0, 0] => [0, 0, 0]
 * 2. [0, 0, 0, 0, 1] => [0, 0, 0]
 * 3. [-1, 0, 1] => [-1, 0, 1]
 * 4. [-1, -1, 0, 0, 1, 1] => [-1, 0, 1]
 * 5. [-1, 0, 0, 0, 1] => [[-1, 0, 1], [0, 0, 0]]
 *
 */

void quicksort(int *a, int l, int r) {
    if (l < r) {
        int i = l, j = r, v = a[r];
        while (i < j) {
            while (i<j && a[i] <= v) i++;
            a[j] = a[i];
            while (i<j && a[j] >= v) j--;
            a[i] = a[j];
        }
        a[i] = v;
        quicksort(a, l, j-1);
        quicksort(a, j+1, r);
    }
}

int cmp(const void *a,const void *b){
    return *(int*)a - *(int*)b;
}

/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** threeSum(int* nums, int numsSize, int* returnSize, int** returnColumnSizes){
    /* quicksort(nums, 0, numsSize-1); */
    qsort(nums,numsSize,sizeof(int),cmp);
    struct Node { int v[3]; struct Node *next; } *h, *z, *t;
    h = malloc(sizeof(*h)), z = malloc(sizeof(*z)), z->next = z, h->next = z, t=h;
    (*returnSize) = 0;
    for (int i = 0; i < numsSize-2; i++) {
        if (nums[i] > 0) break;
        if (i > 0 && nums[i] == nums[i-1]) continue;
        for (int l = i+1, r = numsSize-1; l < r; l++) {
            if (l > i+1 && nums[l] == nums[l-1]) continue;
                /* todo: check index before value */
            while ( (nums[i]+nums[l]) > -nums[r] && l < r ) r--;
            if (l >= r) break;
            if ( (nums[i]+nums[l]) == -nums[r] ) {
                struct Node *temp = malloc(sizeof(struct Node));
                temp->v[0] = nums[i], temp->v[1] = nums[l], temp->v[2] = nums[r];
                temp->next = z, t->next = temp, t = temp;
                (*returnSize)++;
            }
        }
    }
    int** res = malloc(sizeof(int*) * (*returnSize));
    *returnColumnSizes = malloc(sizeof(int) * (*returnSize));
    t = h;
    for (int i = 0; i < *returnSize; i++) {
        t = t->next;
        res[i] = malloc(sizeof(int)*3);
        res[i][0] = t->v[0];
        res[i][1] = t->v[1];
        res[i][2] = t->v[2];
        (*returnColumnSizes)[i] = 3;
    }
    while ( h != z ) {
        t = h;
        h = h->next;
        free(t);
    }
    free(z);
    return res;
}

