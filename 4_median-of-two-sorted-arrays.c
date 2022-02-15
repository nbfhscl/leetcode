#include <stdlib.h>
#include <stdio.h>

/*
 * 困难(hard)
 * 数组(array)
 * 归并排序(merge sort)
 * 二分查找(binary search)
 */

/*
 * Constrains 
 * 1. How large can m and n be? 0 <= m <= 1000, 0 <= n <= 1000
 * 2. Can values in array repeat.
 * 3. Be sure to return double.
 *
 * Ideas and Complexities
 * 1. This seems to be a selection problem. In case that the two array is already
 * sorted, Mergesort can be applied. O(m+n), O(1)
 * 2. Binary search is often used in this kind of problem. O(log(m+n)), O(1)
 *
 * the median eletement locate at for example:
 * 6 -> 2 3
 * 7 -> 3 3
 * 8 -> 3 4
 *
 * Test Cases
 * [], [] => 0.0
 * [], [1] => 1.0
 * [1], [] => 1.0
 * [], [1,1] => 1.0
 * [], [1,2] => 1.50
 * [1,2], [] => 1.50
 * [1,2,3], [] => 2.0
 * [], [1,2,3] => 2.0
 * [1], [1] => 1.0
 * [1], [2] => 1.5
 * [1,2], [1,2] => 1.50
 * [1,2,3], [4,5,6] => 3.50
 * [1,2,3], [4,5] => 3.0
 *
 * Summary
 */

double get_median(int* nums1, int nums1Size, int* nums2, int nums2Size){
    if (nums1Size == 0 && nums2Size == 0) {
        return 0;
    }

    int mid = (nums1Size + nums2Size) / 2;
    int arr[2];

    int p1=0, p2=0, next, k=mid;
    for (;;) {
        if ( k == 0 ) {
            return nums1Size == 1 ? nums1[0] : nums2[0];
        }
        if ( p1 > nums1Size - 1 ) {
            arr[0] = nums2[p2+k-1];
            p2 += k;
            break;
        } else if ( p2 > nums2Size - 1 ) {
            arr[0] = nums1[p1+k-1];
            p1 += k;
            break;
        } else if ( k == 1 ) {
            arr[0] = nums1[p1] < nums2[p2] ? nums1[p1++] : nums2[p2++];
            break;
        }

        next = k/2;
        if ( p1+next-1 > nums1Size-1 ) next = nums1Size - p1;
        if ( p2+next-1 > nums2Size-1 ) next = nums2Size - p2;
        k -= next;

        if ( nums1[p1+next-1] < nums2[p2+next-1] ) {
            p1 += next;
        } else {
            p2 += next;
        }
    }
    if ( p1 > nums1Size - 1 ) {
        arr[1] = nums2[p2];
    } else if ( p2 > nums2Size - 1 ) {
        arr[1] = nums1[p1];
    } else {
        arr[1] = nums1[p1] < nums2[p2] ? nums1[p1++] : nums2[p2++];
    }

    int odd = (nums1Size + nums2Size) % 2;
    return odd ? arr[1] : (arr[0]+arr[1])/2.0;
}


int get_kth_element(int* nums1, int nums1Size, int* nums2, int nums2Size, int k){
    int p1=0, p2=0, next;
    for (;;) {
        if ( p1 > nums1Size - 1 ) return nums2[p2+k-1];
        if ( p2 > nums2Size - 1 ) return nums1[p1+k-1];
        if ( k == 1 ) {
            return nums1[p1] < nums2[p2] ? nums1[p1] : nums2[p2];
        }

        next = k/2;
        if ( p1+next-1 > nums1Size-1 ) next = nums1Size - p1;
        if ( p2+next-1 > nums2Size-1 ) next = nums2Size - p2;
        k -= next;

        if ( nums1[p1+next-1] < nums2[p2+next-1] ) {
            p1 += next;
        } else {
            p2 += next;
        }
    }
    return 0;
}

double binary_search(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int mid = (nums1Size + nums2Size) / 2;
    int odd = (nums1Size + nums2Size) % 2;

    return odd ? get_kth_element(nums1, nums1Size, nums2, nums2Size, mid+1)
        : ( get_kth_element(nums1, nums1Size, nums2, nums2Size, mid+1) 
                + get_kth_element(nums1, nums1Size, nums2, nums2Size, mid)
                ) / 2.0;
}

double binary_search2(int* nums1, int nums1Size, int* nums2, int nums2Size){

    return 2.0;
}

double merge2(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int mid = (nums1Size + nums2Size) / 2;
    int odd = (nums1Size + nums2Size) % 2;
    int arr[2], p1 = 0, p2 = 0, tmp;
    while ( p1+p2 <= mid ) {
        arr[0] = tmp;
        if (p1 >= nums1Size) {
            tmp = nums2[p2++];
        } else if (p2 >= nums2Size) {
            tmp = nums1[p1++];
        } else {
            tmp = nums1[p1] < nums2[p2] ? nums1[p1++] : nums2[p2++];
        }
        arr[1] = tmp;
    }
    return odd ? arr[1] : (arr[0]+arr[1])/2.0;
}

double merge(int* nums1, int nums1Size, int* nums2, int nums2Size){
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

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    return merge(nums1, nums1Size, nums2, nums2Size);
}

int main(int argc, char *argv[]) {
    /* int nums1[4] = {1,3,4,9}; */
    /* int nums2[9] = {1,2,3,4,5,6,7,8,9}; */
    int nums1[2] = {1, 2};
    int nums2[2] = {3, 4};
    printf("%lf\n", get_median(nums1, 2, nums2, 2));
}
