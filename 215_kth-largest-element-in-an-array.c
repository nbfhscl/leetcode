#include <stdio.h>
#include <assert.h>
#include <stdlib.h>

/**
 * review heap
 *
 * medium(中等)
 * quicksort(快排)
 * maximum_heap(最大堆)
 * minimum_heap(最小堆)
 * heap(堆排序)
 *
 * Constrains
 * 1. numsSize? k? [1, 10000]
 * 2. nums[i]? [-10000, 10000]
 *
 * Ideas and Complexities
 * 1. quicksort.
 *      sort only one part to achieve O(n) time complexity
 * 2. heapsort, minimum heap.
 *      iterate to construct a k elements minimum heap
 * 3. heapsort, maximum heap.
 *      construct a maximum heap in place, remove k-1 times the top ele
 *      
 */

#define SWAP(a,b) \
({\
    __auto_type _t = (a);\
    (a) = (b);\
    (b) = _t;\
})
int findKthQuicksort(int* nums, int l, int r, int k) {
    int mid = l + ((r-l)>>1);
    if (nums[l] < nums[r]) SWAP(nums[l], nums[r]);
    if (nums[l] < nums[mid]) SWAP(nums[l], nums[mid]);
    if (nums[r] < nums[mid]) SWAP(nums[r], nums[mid]);
    int ll=l, rr=r, v=nums[rr];
    for (;ll<rr;) {
        for (;ll<rr && nums[ll]>=v;) ll++;
        nums[rr] = nums[ll];
        for (;ll<rr && nums[rr]<=v;) rr--;
        nums[ll] = nums[rr];
    }
    nums[ll] = v;
    if (ll == k) return v;
    else if (ll < k) return findKthQuicksort(nums, ll+1, r, k);
    else return findKthQuicksort(nums, l, ll-1, k);
}
int findKthLargest_quicksort(int* nums, int numsSize, int k){
    return findKthQuicksort(nums, 0, numsSize-1, k-1);
}

int comp_minimal_heap(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}
void heap_up(int* nums, int i, int (* compar)(const void*, const void*)) {
    int j=i, v=nums[j], k;
    for (; (k=(j-1)>>1) >= 0;) {
        if (compar(&nums[k], &v) <= 0) break;
        nums[j] = nums[k], j=k;
    }
    nums[j] = v;
}
void heap_down(int* nums, int numsSize, int (* compar)(const void*, const void*)) {
    int j=0, v=nums[j], k;
    for (; (k=2*j+1)<numsSize; ) {
        if (k<numsSize-1 && compar(&nums[k], &nums[k+1])>0) k++;
        if (compar(&nums[k], &v) >= 0) break;
        nums[j] = nums[k], j = k;
    }
    nums[j] = v;
}
int findKthLargest_minimum_heap(int* nums, int numsSize, int k){
    for (int i=1; i<k; i++) {
        heap_up(nums, i, comp_minimal_heap);
    }
    for (int i=k; i<numsSize; i++) {
        if (nums[i] <= nums[0]) continue;
        nums[0] = nums[i];
        heap_down(nums, k, comp_minimal_heap);
    }
    return nums[0];
}

/*
 *          0
 *     1         2
 *  3     4   5     6
 * 2i+1 2i+2
 */
int comp_maximum_heap(const void* a, const void* b) {
    return *(int*)b - *(int*)a;
}
int findKthLargest_maximum_heap(int* nums, int numsSize, int k) {
    for (int i=1; i<numsSize; i++) {
        heap_up(nums, i, comp_maximum_heap);
    }
    for (int i=1; i<k; i++) {
        nums[0] = nums[numsSize-i];
        heap_down(nums, numsSize-i, comp_maximum_heap);
    }
    return nums[0];
}

int findKthLargest(int* nums, int numsSize, int k){
    /* return findKthLargest_minimum_heap(nums, numsSize, k); */
    return findKthLargest_maximum_heap(nums, numsSize, k);
    /* return findKthLargest_quicksort(nums, numsSize, k); */
}

int main() {
    int nums1[6] = {3, 2, 1, 5, 6, 4};
    assert(findKthLargest(nums1, sizeof(nums1)/sizeof(int), 2) == 5);
    int nums2[7] = {5,4,3,2,1,9,9};
    assert(findKthLargest(nums2, sizeof(nums2)/sizeof(int), 2) == 9);
    int nums3[7] = {5,5,4,3,3,1,1};
    assert(findKthLargest(nums3, sizeof(nums3)/sizeof(int), 4) == 3);
    int nums4[9] = {3,2,3,1,2,4,5,5,6};
    assert(findKthLargest(nums4, sizeof(nums4)/sizeof(int), 4) == 4);
    int nums5[6] = {1,2,3,4,5,6};
    assert(findKthLargest(nums5, sizeof(nums5)/sizeof(int), 1) == 6);
    /* printf("%d\n", findKthLargest(nums, sizeof(nums)/sizeof(int), 2)); */
}
