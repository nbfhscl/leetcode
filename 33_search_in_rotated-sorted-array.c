#include <stdio.h>

/*
 * 中等(medium)
 * 数组(array)
 * 二分查找(binary search)
 *
 * Constrains
 * 1. numsSize? [1,5000]
 * 2. nums[i]? [-10000,10000]
 * 3. nums[i] != nums[j]
 * 4. k? [0,numsSize-1]
 * 5. target? [-10000,10000]
 *
 * Test Cases
 * 1. [4,5,6,7,0,1,2], 0 => 4
 * 2. [4,5,6,7,0,1,2], 3 => -1
 * 3. [1], 0 => -1
 * 4. [0,1,2,3],2 => 2
 * 5. [0,1,2,3,-1], -1 => 4
 *
 * Ideas and Complexities
 * 1. O(n) solution to iterate through the array
 * 2. O(logn) solution to binary search the array
 *
 */

int search(int* nums, int numsSize, int target){
    int l=0, r=numsSize-1, mid;
    for (;l<=r;) {
        mid = l + (r-l)/2;
        if (nums[mid] == target) {
            return mid;
        }
        if (nums[mid] >= nums[l]) {
            if (target >= nums[l] && target < nums[mid]) {
                r = mid-1;
            } else {
                l = mid+1;
            }
        } else {
            if (target > nums[mid] && target <= nums[r]) {
                l = mid+1;
            } else {
                r = mid-1;
            }
        }
    }
    return -1;
}

int main() {
    int nums[7] = {4,5,6,7,0,1,2};
    printf("%d\n", search(nums, 7, 0));
}

