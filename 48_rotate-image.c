/*
 * 中等(medium)
 * 数组(array)
 * 矩阵(matrix)
 *
 * clockwise 90 degree by transpose and left right symmetric transformation
 * 1 2    -    1 3   -   3 1
 * 3 4    -    2 4   -   4 2 
 * counter clockwise 90 degree by transpose and top bottom symmetric
 * transformation
 *                   -   2 4
 *                   -   1 3
 */

void rotate(int** matrix, int matrixSize, int* matrixColSize){
    int tmp;
    for (int i=0; i<matrixSize; i++) {
        for (int j=i+1; j<matrixColSize[i]; j++) {
            tmp = matrix[i][j];
            matrix[i][j] = matrix[j][i];
            matrix[j][i] = tmp;
        }
    }
    for (int i=0; i<matrixSize; i++) {
        for (int l=0,r=matrixColSize[i]-1; l<r; l++,r--) {
            tmp = matrix[i][l];
            matrix[i][l] = matrix[i][r];
            matrix[i][r] = tmp;
        }
    }
}

