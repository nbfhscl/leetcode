/*
 * 中等(medium)
 * 数组(array)
 * 
 */

void setZeroes(int** matrix, int matrixSize, int* matrixColSize){
    int m1=1, n1=1;
    for (int j=0; j<matrixColSize[0]; j++) {
       if (matrix[0][j] == 0) {
           m1 = 0;
           break;
       }
    }
    for (int i=0; i<matrixSize; i++) {
        if (matrix[i][0] == 0) {
            n1 = 0;
            break;
        }
    }
    for (int i=0; i<matrixSize; i++) {
        for (int j=0; j<matrixColSize[i]; j++) {
            if (matrix[i][j] == 0) {
                matrix[0][j] = 0;
                matrix[i][0] = 0;
            }
        }
    }
    for (int i=1; i<matrixSize; i++) {
        for (int j=1; j<matrixColSize[0]; j++) {
            if (matrix[i][0] == 0 || matrix[0][j] == 0) {
                matrix[i][j] = 0;
            }
        }
    }
    if (m1==0) {
        for (int j=0; j<matrixColSize[0]; j++) {
            matrix[0][j] = 0;
        }
    }
    if (n1==0) {
        for (int i=0; i<matrixSize; i++) {
            matrix[i][0] = 0;
        }
    }
}
