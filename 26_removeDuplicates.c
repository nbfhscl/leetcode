#include <stdio.h>
#include <limits.h>

/*
 * 简单(easy)
 * 数组(array)
 * 双指针(double pointer)
 *
 * Constrains
 *
 * Ideas and Complexities
 * 1. Mark eles to be moved. then use double pointer move ele. (bad idea).
 * 2. Use double pointer to find every unique ele and copy into array start
 *    from index 0.
 *
 * Test Cases
 * [] => 0
 * [0] => 1
 * [1,2,3] => 3
 * [1,1,2] => 2
 * [1,2,2] => 2
 */

int removeDuplicates(int* nums, int numsSize){
    if (numsSize == 0) return 0;
    if (numsSize == 1) return 1;
    int i,j;
    for(i=0,j=1;j<numsSize;j++) {
        if (nums[j] != nums[j-1]) {
            nums[++i] = nums[j];
        }
    }
    return i+1;
}

/* int removeDuplicates(int* nums, int numsSize){ */
/*     if (numsSize == 0) return 0; */
/*     if (numsSize == 1) return 1; */
/*     int n = numsSize; */
/*     for (int i=1, hold=INT_MAX; i<numsSize; i++) { */
/*         if ( nums[i] == nums[i-1] || nums[i] == hold ) { */
/*             hold = nums[i]; */
/*             nums[i] = INT_MAX; */
/*             n--; */
/*         } */
/*     } */
/*     int i=1,j=1; */
/*     for ( ;i<numsSize && j<numsSize; ) { */
/*         for (;i<numsSize && nums[i] != INT_MAX;i++); */
/*         for (j=i+1;j<numsSize && nums[j] == INT_MAX;j++); */
/*         if (i<numsSize && j<numsSize) { */
/*             nums[i] = nums[j]; */
/*             nums[j] = INT_MAX; */
/*         } */
/*     } */
/*     return n; */
/* } */

int main(int argc, char* argv[]) {
    int nums[3] = {1,1,2};
    int n = removeDuplicates(nums, 3);
    printf("[");
    for (int i=0; i<n; i++) {
        if (i>0) printf(",");
        printf("%d", nums[i]);
    }
    printf("]\n");
}
