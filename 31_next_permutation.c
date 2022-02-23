#include <stdio.h>

/*
 * 中等(medium)
 * 数组(array)
 * 排列(permutation)
 *
 * Constrains
 * 1. numsSize?
 * 2. nums[i]?
 *
 * Test Cases
 * 1. [1,2,3,4] => [1,2,4,3]
 * 2. [1,3,2,4] => [1,3,4,2]
 * 3. [2,3,4,1] => [2,4,1,3]
 * 4. [2,5,4,3,1] => [3,1,2,4,5]
 * 5. [1,2,1] => [2,1,1]
 *
 * Ideas and Complexities
 * 1. find the first descend value from right to left in the array as value1.
 *    find the least value larger than value1 as value2 from value1 to most
 *    right value of array. swap value1 and value2. reverse values right to
 *    value1.
 *
 */

void reverse_array(int* nums, int l, int r) {
    int tmp;
    for (;l<r;l++,r--) {
        tmp = nums[l];
        nums[l] = nums[r];
        nums[r] = tmp;
    }
}

void nextPermutation(int* nums, int numsSize){
    int i;
    for (i=numsSize-2; i>=0; i--) {
        if (nums[i] < nums[i+1]) break;
    }
    if (i<0) {
        reverse_array(nums, 0, numsSize-1);
        return;
    }
    int j;
    for (j=i+2;j<numsSize;j++) {
        if (nums[j] <= nums[i]) break;
    }
    j--;
    int tmp;
    tmp = nums[i];
    nums[i] = nums[j];
    nums[j] = tmp;
    reverse_array(nums, i+1, numsSize-1);
}

int main() {
    int nums[5] = {2,5,4,3,1};
    nextPermutation(nums, 5);
    printf("[");
    for (int i=0; i<5; i++) {
        if (i>0) printf(",");
        printf("%d", nums[i]);
    }
    printf("]");
}
