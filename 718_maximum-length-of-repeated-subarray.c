#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <math.h>

/*
 * Medium(中等)
 * brute force(暴力)
 * dynamic programming(动态规划)
 * siliding window(滑动窗口)
 * binary search(二分查找)
 * hashmap(哈希表)
 *
 * Constrains
 * 1. numsSize? [1,10000]
 * 2. nums[i]? [0,100]
 *
 * Ideas and Complexities
 * 1. brute force solution
 *      iterate nums1[i] and nums2[j], if nums[i]==nums[j], find the max length
 *      of common subsequence start with nums1[i] and nums2[j]
 *    Time O(n^3), Memory O(1)
 * 2. dynamic programming
 *      if nums1[i]==nums2[j], dp[i,j] = dp[i+1,j+1] + 1;
 *    Time O(n^2), Memory O(n^2)
 * 3. sliding window
 *      one iteration is enough to find common subsequence of nums1 and nums2 
 *      with same index. using sliding window to exhausting the index alignment.
 * 4. binary search and hash
 *      if common subsequence of length k exists, length j<K must exists too.
 *      binary search from length 1 to length n.
 *      calculate hash of all subsequences of nums1 and nums2 with length i, if
 *      hash matches, common subsequence is found.
 */

long qPow(long x, long n, long mod) {
    long ret = 1;
    while (n) {
        if (n & 1) {
            ret = ret * x % mod;
        }
        x = x * x % mod;
        n >>= 1;
    }
    return ret;
}
typedef struct MAP {
    int index;
    int key;
} Map;
int get_index(int key, int size) {
    return abs(key)%size + 1;
}
int store(Map* map, int size, int key) {
    int i = get_index(key, size);
    for (;map[i].index;) i = get_index(i, size);
    map[i].index = i;
    map[i].key = key;
    return i;
}
int find(Map* map, int size, int key) {
    int i = get_index(key, size);
    for (;map[i].index && map[i].key != key;) i = get_index(i, size);
    if (map[i].index && map[i].key == key) return i;
    return 0;
}
int exist(int* a, int sizeA, int* b, int sizeB, int len) {
    // 0 <= nums[i] <= 100, base > 100;
    int base = 113, mod = 1000000009;
    int sizeMapA = 2*(sizeA-len+1);
    Map* mapA = malloc(sizeof(Map) * (sizeMapA+1));
    memset(mapA, 0, sizeof(Map) * (sizeMapA+1));
    long hash=0;
    for (int i=0; i<len; i++) {
        hash = (hash*base + a[i]) % mod;
    }
    store(mapA, sizeMapA, hash);
    long mult = qPow(base, len-1, mod);
    for (int i=len; i<sizeA; i++) {
        hash = ((hash - a[i-len]*mult%mod + mod)%mod*base + a[i])%mod;
        store(mapA, sizeMapA, hash);
    }
    hash = 0;
    for (int i=0; i<len; i++) {
        hash = (hash*base + b[i]) % mod;
    }
    if (find(mapA, sizeMapA, hash)) {
        free(mapA);
        return 1;
    }
    for (int i=len; i<sizeB; i++) {
        hash = ((hash - b[i-len]*mult%mod + mod)%mod*base + b[i])%mod;
        if (find(mapA, sizeMapA, hash)) {
            free(mapA);
            return 1;
        }
    }
    free(mapA);
    return 0;
}
int findLength_bs(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int l=1,r=nums1Size < nums2Size ? nums1Size : nums2Size, m;
    for (;l<r;) {
        m = l + ((r-l+1)>>1);
        if (exist(nums1, nums1Size, nums2, nums2Size, m)) {
            l = m;
        } else {
            r = m - 1;
        }
    }
    if (exist(nums1, nums1Size, nums2, nums2Size, l)) return l;
    else return 0;
}

int findLength_sw(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int ret = 0;
    for (int k=0; k<nums1Size; k++) {
        int len = nums1Size-k < nums2Size ? nums1Size-k : nums2Size;
        if (len < ret) {
            break;
        }
        int num=0;
        for (int i=k,j=0; j<len; i++,j++) {
            if (nums1[i] == nums2[j]) {
                num++;
            } else {
                if (num > ret) {
                    ret = num;
                }
                num = 0;
            }
        }
        if (num > ret) {
            ret = num;
        }
    }
    for (int k=1; k<nums2Size; k++) {
        int len = nums2Size-k < nums1Size ? nums2Size-k : nums1Size;
        if (len < ret) {
            break;
        }
        int num=0;
        for (int i=0,j=k; i<len; i++,j++) {
            if (nums1[i] == nums2[j]) {
                num++;
            } else {
                if (num > ret) {
                    ret = num;
                }
                num = 0;
            }
        }
        if (num > ret) {
            ret = num;
        }
    }
    return ret;
}

int findLength_dp(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int ret=0;
    int dp[nums1Size+1][nums2Size+1];
    memset(dp, 0, sizeof(dp));
    for (int i=nums1Size-1; i>=0; i--) {
        for (int j=nums2Size-1; j>=0; j--) {
            if (nums1[i] == nums2[j]) {
                dp[i][j] = dp[i+1][j+1] + 1;
                if (dp[i][j] > ret) {
                    ret = dp[i][j];
                }
            }
        }
    }
    return ret;
}

int findLength_bf(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int ret = 0;
    for (int i=0; i<nums1Size; i++) {
        for (int j=0; j<nums2Size; j++) {
            if (nums1[i] == nums2[j]) {
                int k = 1;
                for (;i+k<nums1Size && j+k<nums2Size; k++) {
                    if (nums1[i+k] != nums2[j+k]) {
                        break;
                    }
                }
                if (k > ret) {
                    ret = k;
                }
            }
        }
    }
    return ret;
}

int findLength(int* nums1, int nums1Size, int* nums2, int nums2Size){
    return findLength_bs(nums1, nums1Size, nums2, nums2Size);
}

int main() {
    /* int nums1[5] = {1,2,3,2,1}; */
    /* int nums2[5] = {3,2,1,4,7}; */
    /* int nums1[5] = {0,0,0,0,0}; */
    /* int nums2[5] = {0,0,0,0,0}; */
    int nums1[] = {0,0,1,0,0,1,0,0,1,1,1,1,0,0,1,1,1,0,1,0,0,0,0,1,0,1,1,0,1,1,1,1,0,1,1,1,1,0,1,0,1,0,0,1,1,1,1,1,0,0,1,1,1,0,0,1,0,1,1,0,1,1,1,1,1,1,0,1,0,0,1,1,1,1,0,1,0,1,1,1,0,1,1,0,1,0,1,1,1,1,1,0,0,0,0,0,1,0,1,0};
    int nums2[] = {0,0,0,0,1,1,1,0,1,0,1,0,0,0,0,0,0,1,0,1,0,0,0,1,1,1,1,1,1,1,0,1,0,0,1,0,1,0,1,1,0,1,1,1,1,1,0,0,0,0,0,1,1,0,1,0,0,0,1,1,0,0,1,1,1,0,0,1,1,0,0,1,1,1,0,0,0,0,0,0,0,1,0,0,1,0,0,0,1,0,0,0,0,1,1,1,1,1,0,0};
    int ret = findLength(nums1, sizeof(nums1)/sizeof(int), nums2, sizeof(nums2)/sizeof(int));
    printf("%d\n", ret);
}
