#include <stdlib.h>
#include <stdio.h>

/**
 * 简单(easy)
 * 数组(array)
 * 算术(arithmetic)
 */

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* plusOne(int* digits, int digitsSize, int* returnSize){
    *returnSize = digitsSize + 1;
    int* res = malloc(sizeof(int)*(*returnSize));
    int k=1, i=0, sum;
    for (i=digitsSize-1; i>=0; i--) {
        sum = digits[i] + k;
        digits[i] = sum % 10;
        k = sum / 10;
    }
    int j=0;
    if (k != 0) {
        res[j++] = k;
    } else {
        *returnSize = digitsSize;
    }
    for (;j<*returnSize;) {
        res[j++] = digits[++i];
    }
    return res;
}

int main() {
    int arr[4] = {1,2,3,4};
    int returnSize;
    int* res = plusOne(arr, sizeof(arr)/sizeof(int), &returnSize);
    printf("[");
    for (int i=0; i<returnSize; i++) {
        if (i>0) printf(",");
        printf("%d", res[i]);
    }
    printf("]\n");
    free(res);
}

