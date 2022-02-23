#include <stdbool.h>

/*
 * 中等(Medium)
 * 数组(array)
 * 二分查找(binary search)
 *
 * Constrains
 * 1. nums[i] can equal to nums[j]
 *
 * Test Cases
 * 1. [1,1], 0 => false
 * 2. [1,1,1], 0 => false
 * 3. [1,1,1,0,0,0], 0 => true
 * 4. [1,1,1,0,0,1], 0 => true
 *
 * Ideas and Complexities
 *
 */


bool search(int* nums, int numsSize, int target){
    for (int l=0,r=numsSize-1,mid; l<=r;) {
        mid = l + (r-l)/2;
        if (nums[mid] == target) {
            return true;
        }
        if (nums[mid] > nums[l]) {
            if (target >= nums[l] && target < nums[mid]) {
                r = mid-1;
            } else {
                l = mid+1;
            }
        } else if (nums[mid] < nums[l]) {
            if (target > nums[mid] && target <= nums[r]) {
                l = mid+1;
            } else {
                r = mid-1;
            }
        } else {
            while (nums[++l] == nums[mid]);
        }
    }
    return false;
}

