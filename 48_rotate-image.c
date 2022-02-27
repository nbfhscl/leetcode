/*
 * 中等(medium)
 * 数组(array)
 * 矩阵(matrix)
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

