

/*
 * 中等(medium)
 * 数组(array)
 * 动态规划(dynamic programming)
 * 分治(divide and conquer)
 * 线段树(segment tree)
 *
 * Constrains
 * 1. numsSize? [1,100000]
 * 2. nums[i]? [-10000,10000]
 * 3. sum(nums[i]+...+nums[i+n])?
 *
 * Ideas and Complexities
 * max_sub_arr_sum_end_with[i] = max(0, max_sub_arr_sum_end_with[i-1]) + nums[i] i>0
 * Time O(n), Memory O(1)
 */

/* #define MAX(a,b) \ */
/* ({ \ */
/*     __auto_type _a = (a); \ */
/*     __auto_type _b = (b); \ */
/*     _a > _b ? _a : _b; \ */
/* }) */

#define MAX(a,b) ((a) > (b) ? (a) : (b))

int maxSubArray(int* nums, int numsSize){
    int max=nums[0], maxLast=nums[0];
    for (int i=1;i<numsSize;i++) {
        maxLast = MAX(maxLast, 0) + nums[i];
        max = MAX(max, maxLast);
    }
    return max;
}

