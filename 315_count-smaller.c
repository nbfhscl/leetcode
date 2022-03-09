#include <stdlib.h>
#include <string.h>
#include <stdio.h>

/*
 * Constrains
 * 1. numsSize? [1, 100000]
 * 2. nums[i]? [-10000, 10000]
 * 3. nums[i] == nums[j], i != j. true
 * Ideas and Complexities
 * 1. iterate. Time O(n^2), Memory O(n).
 * 2. insert sort. Time O(n^2), Memory O(n).
 * 3. binary search and insert. Time O(n(logn+n)) = O(n^2), Memory O(n).
 * 4. mergesort. mergesort is often used to count reverse order pairs. Time
 *    O(nlogn), Memory O(n).
 */

// 2. 超时
/* int* countSmaller(int* nums, int numsSize, int* returnsize){ */
/*     int* res = malloc(sizeof(int)*numsSize); */
/*     *returnsize = numsSize; */
/*     int a[numsSize], j=0; */
/*     a[j] = nums[numsSize-1], res[numsSize-1] = 0; */
/*     for (int i=numsSize-2; i>=0; i--) { */
/*         int v = nums[i], k; */
/*         for (k=++j; k>0 && a[k-1]>=v; k--) { */
/*             a[k] = a[k-1]; */
/*         } */
/*         a[k] = v; */
/*         res[i] = k; */
/*     } */
/*     return res; */
/* } */

// 3. 超时
/*int find_the_biggest_less_than_target(int* a, int size, int target) {*/
/*    int l=0, r=size-1, m, v;*/
/*    for (;l<r;) {*/
/*        m = l + ((r-l+1)>>1);*/
/*        v = a[m];*/
/*        if (v == target) {*/
/*            r = m - 1;*/
/*        } else if (v < target) {*/
/*            l = m;*/
/*        } else r = m - 1;*/
/*    }*/
/*    if (a[l] < target) return l;*/
/*    else return -1;*/
/*}*/
/*int* countSmaller(int* nums, int numsSize, int* returnsize){*/
/*    int* res = malloc(sizeof(int)*numsSize);*/
/*    *returnsize = numsSize;*/
/*    int a[numsSize];*/
/*    a[0] = nums[numsSize-1], res[numsSize-1] = 0;*/
/*    for (int i=numsSize-2; i>=0; i--) {*/
/*        int j = numsSize-i-1;*/
/*        int k = find_the_biggest_less_than_target(a, j, nums[i]);*/
/*        for (int jj=j; jj>k+1; jj--) a[jj] = a[jj-1];*/
/*        a[k+1] = nums[i];*/ 
/*        res[i] = k+1;*/
/*    }*/
/*    return res;*/
/*}*/

void merge_sort_count(int* a, int l, int r, int* res, int* index) {
    if (l >= r) return;
    int m = l + ((r-l)>>1);
    merge_sort_count(a, l, m, res, index);
    merge_sort_count(a, m+1, r, res, index);
    int tmp[r-l+1], i1=l, i2=m+1, j=0;
    for (;i1<=m && i2<=r;) {
        if (a[index[i1]] > a[index[i2]]) {
            tmp[j++] = index[i1];
            res[index[i1]] += r-i2+1;
            i1++;
        } else {
            tmp[j++] = index[i2++];
        }
    }
    for (;i1<=m;) {
        tmp[j++] = index[i1++];
    }
    for (;i2<=r;) {
        tmp[j++] = index[i2++];
    }
    for (int i=0; i<j; i++) {
        index[l+i] = tmp[i];
    }
}

int* countSmaller(int* nums, int numsSize, int* returnsize){
    int* res = malloc(sizeof(int)*numsSize); *returnsize = numsSize;
    memset(res, 0, sizeof(int)*numsSize);
    int* index = malloc(sizeof(int)*numsSize);
    for (int i=0; i<numsSize; i++) index[i] = i;
    merge_sort_count(nums, 0, numsSize-1, res, index);
    free(index);
    return res;
}

int main() {
    int nums[6] = {5,4,3,2,1,0};
    int returnSize;
    int* res = countSmaller(nums, sizeof(nums)/sizeof(int), &returnSize);
    printf("[");
    for (int i=0; i<returnSize; i++) {
        if (i>0) printf(",");
        printf("%d", res[i]);
    }
    printf("]\n");
    free(res);
}

