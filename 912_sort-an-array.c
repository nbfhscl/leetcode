#include <stdio.h>

/**
 * review quicksort
 *
 * 中等(medium)
 * 数组(array)
 * 快排(quicksort)
 *
 * Constrains
 * 1. numsSize? [1, 50000]
 * 2. nums[i]? [-50000, 50000]
 *
 */

#define SWAP(a, b) \
({ \
    __auto_type _c = (a);\
    ( a ) = ( b ); \
    ( b ) = _c; \
})

void quicksort(int* nums, int l, int r) {
    // [5,25] works about the same
    if (r-l < 15) {
        for (int i=l+1;i<=r;i++) {
            int j=i, v=nums[i];
            for (; j>0 && nums[j-1]>v;j--) nums[j] = nums[j-1];
            nums[j]=v;
        }
        // dont forget to return here
        return;
    }
    int mid = l + ((r-l)>>1);
    // make nums[mid] the largest
    if (nums[mid] < nums[l]) SWAP(nums[mid], nums[l]);
    if (nums[mid] < nums[r]) SWAP(nums[mid], nums[r]);
    // make nums[r] the medium
    if (nums[r] < nums[l]) SWAP(nums[r], nums[l]);
    int ll = l, rr = r, v = nums[rr];
    for (;ll<rr;) {
        for (;ll<rr && nums[ll] <= v;) ll++;
        nums[rr] = nums[ll];
        for (;ll<rr && nums[rr] >= v;) rr--;
        nums[ll] = nums[rr];
    }
    // ll==rr
    nums[ll] = v;
    quicksort(nums, l, ll-1);
    quicksort(nums, ll+1, r);
}

void quicksort_stack(int* nums, int l, int r) {
    int stack[64] = {}, si=0;
    for (;;) {
        if (r-l < 15) {
            for (int i=l+1;i<=r;i++) {
                int j=i, v=nums[i];
                for (; j>0 && nums[j-1]>v;j--) nums[j] = nums[j-1];
                nums[j]=v;
            }
            if (si<=0) return;
            // pop from stack instead of return
            // pop using --si rather than si--
            l = stack[--si], r = stack[--si];
            continue;
        }
        int mid = l + ((r-l)>>1);
        if (nums[mid] < nums[l]) SWAP(nums[mid], nums[l]);
        if (nums[mid] < nums[r]) SWAP(nums[mid], nums[r]);
        if (nums[r] < nums[l]) SWAP(nums[r], nums[l]);
        int ll = l, rr = r, v = nums[rr];
        for (;ll<rr;) {
            for (;ll<rr && nums[ll] <= v;) ll++;
            nums[rr] = nums[ll];
            for (;ll<rr && nums[rr] >= v;) rr--;
            nums[ll] = nums[rr];
        }
        nums[ll] = v;
        // end-recursion-removal and the
        // policy of processing the smaller of the two subfiles
        // first ensures the room needed for stack to be as small
        // as possible(logn). processing the larger part first will need
        // more rooms for stack typically.
        if (ll-l > r-ll) {
            // push using si++ rather than ++si
            stack[si++] = ll-1, stack[si++] = l, l=ll+1;
        } else {
            stack[si++] = r, stack[si++] = ll+1, r=ll-1;
        }
    }
}

int* sortArray(int* nums, int numsSize, int* returnSize){
    *returnSize = numsSize;
    quicksort_stack(nums, 0, numsSize-1);
    return nums;
}

int main(int argc, char *argv[])
{
    int nums[98] = {-74,48,-20,2,10,-84,-5,-9,11,-24,-91,2,-71,64,63,80,28,
        -30,-58,-11,-44,-87,-22,54,-74,-10,-55,-28,-46,29,10,50,-72,34,26,25,8,
        51,13,30,35,-8,50,65,-6,16,-2,21,-78,35,-13,14,23,-3,26,-90,86,25,-56,
        91,-13,92,-25,37,57,-20,-69,98,95,45,47,29,86,-28,73,-44,-46,65,-84,
        -96,-24,-12,72,-68,93,57,92,52,-45,-2,85,-63,56,55,12,-85,77,-39};
    int returnSize;
    sortArray(nums, sizeof(nums)/sizeof(int), &returnSize);
    printf("[");
    for (int i=0; i<returnSize; i++) {
        if (i>0) printf(",");
        printf("%d", nums[i]);
    }
    printf("]\n");
    return 0;
}
