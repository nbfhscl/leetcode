#include <stdio.h>
#include <stdlib.h>

/*
 * review binarysearch
 * 中等(medium)
 * 数组(array)
 * 二分查找(binary search)
 *
 * Constrains
 * 1. numsSize?
 * 2. nums[i]?
 * 3. target?
 * 4. nums[i+1] >= nums[i], i in [0,numsSize-2]
 *
 * Test Cases
 *
 * Ideas and Complexities
 * binary search
 * 1. if match, linearly find left boundary with one pointer and right with another.
 * 2. if match, to find the left boundary, keep the match as right boundary, on
 *    the other hand, to find the right boudary, keep the match as left boudary.
 *    be careful about the break condition (l<r instead of l<=r), and the mid
 *    calculation (mid = l + (r-l)/2 + 1 instead of mid = l + (r-l)/2 for right
 *    side finding)
 *
 */

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* searchRange(int* nums, int numsSize, int target, int* returnSize){
    *returnSize = 2;
    int *res = malloc(sizeof(int) * 2);
    res[0] = -1; res[1] = -1;
    if (numsSize <= 0) {
        return res;
    }
    int l, r, mid;
    // < rather than <= to stop the cycle
    for (l=0,r=numsSize-1; l < r;) {
        mid = l + (r-l)/2;
        if (target <= nums[mid]) {
            r = mid;
        } else {
            l = mid + 1;
        }
    }
    if (target == nums[l]) {
        res[0] = l;
    }
    for (l=0, r=numsSize-1; l < r;) {
        // +1 to ensure break condition will be reached, ex. [1,1]
        mid = l + (r-l)/2+1;
        if (target < nums[mid]) {
            r = mid - 1;
        } else {
            l = mid;
        }
    }
    if (target == nums[l]) {
        res[1] = l;
    }
    return res;
}

int main() {
    int returnSize, *res;
    /* int len = 7, nums[7] = {0,1,3,3,3,4,5}, target = 3; */
    /* int len = 7, nums[7] = {3,3,3,3,3,3,3}, target = 3; */
    /* int len = 7, nums[7] = {3,3,3,3,3,3,3}, target = 1; */
    int len = 6, nums[6] = {5,7,7,8,8,10}, target = 6;
    res = searchRange(nums, len, target, &returnSize);
    printf("[%d, %d]", res[0], res[1]);
    free(res);
}

