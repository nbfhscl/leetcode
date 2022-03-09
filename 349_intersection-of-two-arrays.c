#include <stdlib.h>
#include <stdio.h>

/**
 * Constrains
 * 1. nums1Size? [1,100]
 * 2. nums2Size? [1,100]
 * 3. nums1[i]? [0,1000]
 * 4. nums2[i]? [0,1000]
 * 5. nums[i] == nums[j], i != j.? true
 *
 * Ideas and Complexities
 * 1. if nums1Max < nums2Min return null;
 *    if nums1Min > nums2Max return null;
 *    else min = MAX(nums1Min, nums2Min), max = MIN(nums1Max, nums2Max);
 *    find value in [min, max] and put into result if not already in.
 *    not a good idea compared to these below.
 * 2. use a hashmap to store one array, iterate the other. Time O(m+n), Memory
 *    O(min(m,n)).
 * 3. sort, use double pointer. Time O(mlogm+nlogn), Memory O(logm+logn).
 *    if a1[p1] < a2[p2] p1++
 *    else if a1[p1] == a2[p2] res[i] = a1[p1]
 *    else p2++;
 *
 */

void quicksort(int* a, int l, int r) {
    if (l >= r) return;
    int ll=l, rr=r, v = a[r];
    for (;ll<rr;) {
        for (;ll<rr && a[ll]<=v;) ll++;
        a[rr] = a[ll];
        for (;ll<rr && a[rr]>=v;) rr--;
        a[ll] = a[rr];
    }
    a[ll] = v;
    quicksort(a, l, ll-1);
    quicksort(a, ll+1, r);
}

int compar(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* intersection(int* nums1, int nums1Size, int* nums2, int nums2Size, int* returnSize){
    int* res = malloc(sizeof(int)*(nums1Size > nums2Size ? nums1Size : nums2Size));
    quicksort(nums1, 0, nums1Size-1);
    quicksort(nums2, 0, nums2Size-1);
    /* qsort(nums1, nums1Size, sizeof(int), compar); */
    /* qsort(nums2, nums2Size, sizeof(int), compar); */
    int i1=0, i2=0, i=0;
    for (; i1<nums1Size && i2<nums2Size;) {
        if (nums1[i1] == nums2[i2]) {
            if (i==0 || res[i-1] != nums1[i1]) {
                res[i++] = nums1[i1];
            }
            i1++;i2++;
        } else if (nums1[i1] < nums2[i2]) {
            i1++;
        } else i2++;
    }
    *returnSize = i;
    return res;
}

int main() {
    int nums1[3] = {4,9,5};
    int nums2[5] = {9,4,9,8,4};
    int returnSize;
    int* res = intersection(nums1, sizeof(nums1)/sizeof(int), nums2, sizeof(nums2)/sizeof(int), &returnSize);
    printf("[");
    for (int i=0; i<returnSize; i++) {
        if (i>0) printf(",");
        printf("%d", res[i]);
    }
    printf("]\n");
}
