#include <stdlib.h>
#include <limits.h>

/*
 * 中等(medium)
 * 数组(array)
 * 双指针(double pointer)
 */

/*
 * Constrants
 * 1. How large can n be? 3 <= numsSize <= 1000
 * 2. How large can every number be? -1000 <= nums[i] <= 1000
 * 3. How large can target be? -10^4 <= target <= 10^4
 * 4. Can a number be choosed multiple times?
 *
 * Ideas and Complexities
 *
 * Test Cases
 * [-1, 0, 0, 1], 5 => 1
 * [-1, 0, 0, 1], -2 => -1
 * [-2, -1, 0, 1, 2, 4, 5], 0 => 0
 *
 */

int compar(const void *a, const void *b) {
    return *(int*)a - *(int*)b;
}

int threeSumClosest(int* nums, int numsSize, int target){
    qsort(nums, numsSize, sizeof(int), compar);
    int res;
    for (int i = 0; i < numsSize - 2; i++) {
        if (i > 0 && nums[i] == nums[i-1]) continue;
        for (int l = i+1, r = numsSize-1; l < r;) {
            int sum = nums[i] + nums[l] + nums[r];
            if ( abs(sum - target) < abs(res - target) ) {
                res = sum;
            }
            if ( sum == target ) {
                return target;
            } else if ( sum < target ) {
                l++;
            } else {
                r--;
            }
        }
    }
    return res;
}

