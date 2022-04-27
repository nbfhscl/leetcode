#include "io.h"

void printRet(int** ret, int returnSize, int* returnColumnSizes) {
    printf("[");
    for (int i=0; i<returnSize; i++) {
        if (i > 0) printf(",");
        printf("[");
        for (int j=0; j<returnColumnSizes[i]; j++) {
            if (j > 0) printf(",");
            printf("%d", ret[i][j]);
        }
        printf("]");
        free(ret[i]);
    }
    printf("]\n");
    free(returnColumnSizes);
    free(ret);
}
