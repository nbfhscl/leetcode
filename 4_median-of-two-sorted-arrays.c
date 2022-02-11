#include <stdlib.h>

/*
 * 困难(hard)
 * 数组(array)
 * 二分查找(binary search)
 */

/*
 * Constrains 
 * 1. How large can m and n be?
 * 2. Can values in array repeat.
 *
 * Ideas and Complexities
 * This seems to be a selection problem. In case that the two array is already
 * sorted, Mergesort can be applied.
 *
 * the median eletement locate at for example:
 * 6 -> 2 3
 * 7 -> 3 3
 * 8 -> 3 4
 *
 * Test Cases
 *
 * Summary
 */

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int med = (nums1Size + nums2Size) / 2;
    int *arr = (int*)malloc((med + 1) * sizeof(int));
    int p1=0, p2=0, p=0;
    while(p <= med) {
        if (p1 >= nums1Size) {
            arr[p++] = nums2[p2++];
        } else if (p2 >= nums2Size) {
            arr[p++] = nums1[p1++];
        } else {
            arr[p++] = nums1[p1] < nums2[p2] ? nums1[p1++] : nums2[p2++];
        }
    }
    return (nums1Size + nums2Size) % 2 == 0 ?
                (arr[med-1] + arr[med]) / 2.0 :
                arr[med];
}

