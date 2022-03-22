#include <stdlib.h>
#include <stdio.h>

/* 
 * medium(中等)
 * dynamic programming(动态规划)
 * binary search(二分查找)
 *
 * Constrains
 * 1. numsSize? [1, 2500]
 * 2. nums[i]? [-10000, 10000]
 *
 * Ideas and Complexities
 * 1. a[i], lengthOfLIS end with nums[i]
 *    a[i] = max(a[j])+1, 0<=j<i && nums[j]<nums[i]
 *    a[0] = 1
 *    Time O(n^2), Memory O(n)
 * 2. a[i]，长度为i的递增子序列中第i位可能的最小值
 *    The minimum value of the ith of the increasing subsequence with length i
 */

int find_the_minimal_greater_than_or_equal_to_target(int* a, int size, int target) {
    int l=0, r=size-1, m;
    for (;l<r;) {
        m = l + ((r-l)>>1);
        if (a[m] == target) {
            r = m;
        } else if (a[m] < target) {
            l = m + 1;
        } else {
            r = m;
        }
    }
    if (a[l] >= target) return l;
    else return -1;
}

int lengthOfLIS_nlogn(int* nums, int numsSize){
    int res = 0;
    int a[numsSize];
    a[0] = nums[0];
    for (int i = 1; i < numsSize; i++) {
        if (nums[i] > a[res]) {
            a[++res] = nums[i];
        } else {
            int index = find_the_minimal_greater_than_or_equal_to_target(a, res+1, nums[i]);
            if (index != -1 && a[index] > nums[i]) {
                a[index] = nums[i];
            }
        }
    }
    return res+1;
}

int lengthOfLIS_n2(int* nums, int numsSize){
    int res = 1;
    int* a = malloc(sizeof(int)*numsSize);
    a[0] = 1;
    for (int i = 1; i < numsSize; i++) {
        int max = 0;
        for (int j = 0; j < i; j++) {
            if (nums[j] < nums[i] && a[j] > max) {
                max = a[j];
            }
        }
        a[i] = max + 1;
        if (a[i] > res) {
            res = a[i];
        }
    }
    free(a);
    return res;
}

int lengthOfLIS(int* nums, int numsSize){
    /* return lengthOfLIS_n2(nums, numsSize); */
    return lengthOfLIS_nlogn(nums, numsSize);
}

int main() {
    int nums[8] = {10,9,2,5,3,7,101,18};
    int res = lengthOfLIS(nums, sizeof(nums)/sizeof(int));
    printf("%d\n", res);
}

