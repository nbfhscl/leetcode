

/*
 * review dp and divide and conquer
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

int maxSubArray_dp(int* nums, int numsSize){
    int max=nums[0], maxLast=nums[0];
    for (int i=1;i<numsSize;i++) {
        maxLast = MAX(maxLast, 0) + nums[i];
        max = MAX(max, maxLast);
    }
    return max;
}

/**
 * lSum: maximum sum of sub array start with left element
 * rSum: maximum sum of sub array end with right element
 * mSum: maximum sum of sub array
 * tSum: total sum of array elements
 */
typedef struct {
    int l_sum;
    int r_sum;
    int m_sum;
    int t_sum;
} Sums;
Sums maxSubArray_divide_and_conquer(int* a, int l, int r) {
    // 注意到这里m_sum并不是0
    if (l==r) return (Sums){a[l], a[l], a[l], a[l]};
    int mid = l + ((r-l)>>1);
    Sums left_sub = maxSubArray_divide_and_conquer(a, l, mid);
    Sums right_sub = maxSubArray_divide_and_conquer(a, mid+1, r);
    return (Sums){
        MAX(left_sub.l_sum, left_sub.t_sum+right_sub.l_sum),
        MAX(right_sub.r_sum, right_sub.t_sum+left_sub.r_sum),
        MAX(MAX(left_sub.m_sum, right_sub.m_sum), left_sub.r_sum+right_sub.l_sum),
        left_sub.t_sum+right_sub.t_sum
    };
}
int maxSubArray(int* nums, int numsSize){
    return maxSubArray_divide_and_conquer(nums, 0, numsSize-1).m_sum;
}
