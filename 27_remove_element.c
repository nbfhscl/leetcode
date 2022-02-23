
/*
 * 简单(easy)
 * 数组(array)
 * 双指针(double pointer)
 *
 * Constrains
 * 1. n?
 * 2. nums[n]?
 * 3. val?
 *
 * Ideas and Complexities
 * 1. range and copy. O(n), O(1)
 *
 * Test Cases
 *
 */

int removeElement(int* nums, int numsSize, int val){
    int j = 0;
    for (int i=0; i<numsSize; i++) {
        if (nums[i] != val) {
            nums[j++] = nums[i];
        }
    }
    return j;
}

