#include <stdio.h>


/*
 * 中等(medium)
 * 数组(array)
 * 贪心(greedy)
 *
 * Constrains
 * 1. numsSize?
 * 2. nums[i]?
 */

/* find the next max choice and jump to it
 */
int jump(int* nums, int numsSize){
    int res=0;
    for (int i=0;i<numsSize-1;) {
        int max = i+nums[i], nextI = max; res++;
        if (nextI >= numsSize-1) break;
        for (int j=i+1; j<=i+nums[i] && j<numsSize; j++) {
            if (j+nums[j]>max) {
                max = j + nums[j];
                nextI = j;
            }
        }
        i = nextI;
    }
    return res;
}

#define MAX(a,b) ((a) > (b) ? (a) : (b))

/* a more simple code base
 * [1,3,4,1,1,1,2]
 *  - initialize
 *    - first step
 *      - - - second step
 *            - - third step
 *  
 */
int jump_2(int* nums, int numsSize) {
    int max=0, end=0, step=0;
    for (int i=0; i<numsSize-1; i++) {
        max = MAX(max, i+nums[i]);
        if (i == end) {
            end = max;
            step++;
        }
    }
    return step;
}

int main() {
    int nums[] = {1,2,3};
    int res = jump(nums, sizeof(nums)/sizeof(int));
    printf("%d\n", res);
}

