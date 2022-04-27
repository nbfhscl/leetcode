#include <stdio.h>
#include <assert.h>

/*
 * review binary search special case need to handle
 *
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
    int l = 0, r = numsSize-1, m;
    for (; l<r;) {
        m = l + ((r-l)>>1);
        if (nums[m] == target) {
            return m;
        } else if (nums[m] < target) {
            l = m + 1;
        } else {
            r = m;
        }
    }
    if (l == numsSize-1 && nums[numsSize-1]<target) l++;
    return l;
}

int main() {
    /* int nums[5] = {0,1,3,4,5}, target = 2; */
    int nums[4] = {1, 3, 5, 6}, target = 7;
    printf("%d\n",searchInsert(nums, sizeof(nums)/sizeof(int), target));
    assert(searchInsert(nums, sizeof(nums)/sizeof(int), target) == 4);
}

