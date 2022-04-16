#include <stdbool.h>
#include <stdio.h>


/*
 * 中等(medium)
 * 数组(array)
 * 贪心(greedy)
 *
 * Constrains
 * 1. numsSize? [1,30000]
 * 2. nums[i]? [0,100000]
 *
 * Test Cases
 * 1. [2,3,1,1,4] => true. 2
 * 2. [3,2,1,0,4] => false.
 */

#define MAX(a,b) ((a) > (b) ? (a) : (b))

bool canJump(int* nums, int numsSize){
    int max=0, end=0, i;
    for (i=0; i<numsSize; i++) {
        max = MAX(max, i+nums[i]);
        if (i==end) {
            end = max;
            if (max <= i) break;
        }
    }
    if (i < numsSize -1) return false;
    return true;
}

bool canJump_2(int* nums, int numsSize){
    int end = 0;
    for (int i=0; i<numsSize && i<=end; i++) {
        end = MAX(end, i+nums[i]);
        if (end >= numsSize-1) {
            return true;
        }
    }
    return false;
}

int main() {
    int nums[5] = {2,3,1,1,4};
    /* int nums[5] = {3,2,1,0,4}; */
    bool bl = canJump_2(nums, sizeof(nums)/sizeof(int));
    printf("%d\n", bl);
}

