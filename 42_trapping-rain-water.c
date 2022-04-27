#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <math.h>
#include "utstack.h"

/*
 * 困难(hard)
 * 数组(array)
 * 动态回归(dynamic programming)
 * 双指针(double pointer)
 * 单调栈(monotone stack)
 *
 * Constrains
 * 1. n? [1, 20000]
 * 2. height? [0,105]
 *
 * Test Cases
 * 1. [0,   1,    0,    2,    1,    0,    1,    3,    2,    1,    2,    1] => 6
 *    [0,3],[0,3],[1,3],[1,3],[2,3],[2,3],[2,3],[2,2],[3,2],[3,2],[3,1],[3,0] =>
 *    0+    0+    1+    0+    1+    2+    1+    0+    0+    1+    0+    0 = 6
 *
 * 2. [4,2,0,3,2,5] => 9
 *    [0,5],[4,5],[4,5],[4,5],[4,5],[4,0] => 2+4+1+2=9
 *
 * Ideas and Complexities
 * maxLeftHeight[i] = max(maxLeftHeight[i-1], height[i])
 * maxRightHeight[i] = max(maxRightHeight[i+1], height[i])
 * max[i] = max(min(maxLeft[i],maxRight[i])-height[i], 0)
 *
 * Ideas and Complexities
 * 1. dynamic programming, see above. Time O(n), Memory O(n).
 * 2. double pointer, variety of dynamic programming to reduce memory cost to
 *    O(1).
 * 3. monotone stack, Time O(n), Memory O(n).
 *
 */

/* #define max(a,b) ((a) > (b) ? (a) : (b)) */
/* #define min(a,b) ((a) < (b) ? (a) : (b)) */
/* #define max(a,b) ({ \ */
/*             __typeof__ (a) _a = (a); \ */
/*             __typeof__ (b) _b = (b); \ */
/*             _a > _b ? _a : _b; \ */
/*         }) */
/* #define min(a,b) ({ \ */
/*             __typeof__ (a) _a = (a); \ */
/*             __typeof__ (b) _b = (b); \ */
/*             _a > _b ? _b : _a; \ */
/*         }) */ 
#define max(a,b) \
({ \
    __auto_type _a = (a); \
    __auto_type _b = (b); \
    _a > _b ? _a : _b;  \
})
#define min(a,b) \
({ \
    __auto_type _a = (a); \
    __auto_type _b = (b); \
    _a > _b ? _b : _a;  \
})

int trap(int* height, int heightSize){
    int *maxLeftHeight = malloc(sizeof(int)*heightSize); maxLeftHeight[0] = 0;
    int *maxRightHeight = malloc(sizeof(int)*heightSize); maxRightHeight[heightSize-1] = 0;
    for (int i=1;i<heightSize;i++) {
        maxLeftHeight[i] = max(maxLeftHeight[i-1], height[i-1]);
        maxRightHeight[heightSize-i-1] = max(maxRightHeight[heightSize-i], height[heightSize-i]);
    }
    int sum=0;
    for (int i=0;i<heightSize;i++) {
        sum += max(min(maxLeftHeight[i], maxRightHeight[i])-height[i], 0);
    }
    free(maxLeftHeight);
    free(maxRightHeight);
    return sum;
}

int trap_double_pointer(int* height, int heightSize) {
    int maxLeft = 0, maxRight = 0, sum = 0, l = 0, r = heightSize-1;
    for (;l<=r;) {
        // no need to maintain two array for maxLeft and maxRight
        // instead two var is enough
        if (maxLeft <= maxRight) {
            sum += max(0, maxLeft-height[l]);
            maxLeft = max(maxLeft, height[l]);
            l++;
        } else {
            sum += max(0, maxRight-height[r]);
            maxRight = max(maxRight, height[r]);
            r--;
        }
    }
    return sum;
}

int trap_monotone_stack(int *height, int heightSize) {
    int sum=0;
    typedef struct item {int i; struct item *next;} Item; Item *stack = NULL;
    Item *item;
    for (int i=0;i<heightSize;i++) {
        for (;!STACK_EMPTY(stack) && height[i] > height[STACK_TOP(stack)->i];) {
            STACK_POP(stack, item);
            if (STACK_EMPTY(stack)) {
                free(item);
                break;
            }
            int w = i-STACK_TOP(stack)->i-1;
            int h = min(height[STACK_TOP(stack)->i], height[i]) - height[item->i];
            sum += w*h;
            free(item);
        }
        item = malloc(sizeof(*item));
        item->i = i;
        STACK_PUSH(stack, item);
    }
    for (;!STACK_EMPTY(stack);) {
        STACK_POP(stack, item);
        free(item);
    }
    return sum;
}

int main() {
    /* int height[12] = {0,1,0,2,1,0,1,3,2,1,2,1}, len = 12; */
    int height[6] = {4,2,0,3,2,5}, len = 6;
    int res = trap(height, len);
    printf("%d\n", res);
    res = trap_double_pointer(height, len);
    printf("%d\n", res);
    res = trap_monotone_stack(height, len);
    printf("%d\n", res);
}

