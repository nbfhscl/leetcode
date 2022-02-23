#include <stdio.h>

/*
 * 简单(easy)
 * 数组(array)
 * 二分查找(binary search)
 *
 * Contrains
 *
 * Test Cases
 * 1. [0,1,3,4,5], 2 => 2
 * 2. [0,1,2,4,5], 3 => 3
 *
 */


int searchInsert(int* nums, int numsSize, int target){
    int l, r, mid;
    for (l=0,r=numsSize-1; l<=r;) {
        mid = l + (r-l)/2;
        if (target == nums[mid]) {
            return mid;
        } else if (target < nums[mid]) {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }
    return l;
}

int main() {
    /* int nums[5] = {0,1,3,4,5}, target = 2; */
    int nums[5] = {0,1,2,4,5}, target = 3;
    printf("%d\n", searchInsert(nums, 5, target));
}

