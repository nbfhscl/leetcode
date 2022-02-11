#include "two-sum.c"
#include <assert.h>
#include <stdio.h>

struct TC {
    int* num;
    int numsSize;
    int target;

    int* ans;
};

int array_equal(int* arr1, int size1, int* arr2, int size2) {
    if (size1 != size2) {
        return -1;
    }
    for (int i=0;i<size1;i++) {
        if (arr1[i] != arr2[i]) {
            return -1;
        }
    }
    return 0;
}
int main() {
    struct TC tc[] = {
        /* (struct TC){.num=(int[2]){0,1}, .numsSize=2, .target=1, (int[2]){0,1}}, */
        /* (struct TC){.num=(int[4]){0,1,2,3}, .numsSize=4, .target=5, (int[2]){2,3}}, */
        /* (struct TC){.num=(int[2]){2,2}, .numsSize=2, .target=4, (int[2]){0,1}}, */
        /* (struct TC){.num=(int[5]){1,2,3,4,5}, .numsSize=5, .target=10, NULL}, */
        /* (struct TC){.num=(int[5]){30, 33, 87, 108, 99}, .numsSize=5, .target=120, (int[2]){1,2}}, */
        (struct TC){.num=(int[4]){2,5,5,11}, .numsSize=4, .target=10, (int[2]){1,2}},
        (struct TC){.num=(int[4]){0,4,3,0}, .numsSize=4, .target=0, (int[2]){0,3}}
    };
    int TCSIZE = 2;

    /* printf("%d", tc[1].num[0]); */
    /* for (int i=0;i<tc[1].numsSize;i++) { */
    /*     printf("%d", tc[1].num[i]); */
    /* } */

    for (int i=0;i<TCSIZE;i++) {
        printf("Test Case %d: ", i+1);
        printf("[");
        int j;
        for (j=0;j<tc[i].numsSize-1;j++){
            printf("%d,", tc[i].num[j]);
        }
        printf("%d", tc[i].num[j]);
        printf("]");
        printf(", ");
        printf("%d", tc[i].target);
        printf("\t");

        int returnSize=2;
        int* res = twoSum_n1(tc[i].num, tc[i].numsSize, tc[i].target, &returnSize);
        if ((tc[i].ans == NULL && res != NULL)
                || (tc[i].ans != NULL && array_equal(tc[i].ans,2,res,returnSize) != 0)) {
            printf("Failed\n");
            return 0;
        }
        printf("Succeed!\n");
    }

    printf("All Test Cases Passed!\n");
    return 0;
}
