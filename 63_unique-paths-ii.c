#include <stdlib.h>
#include <stdio.h>

/*
 * 中等(medium)
 * 数组(array)
 * 动态规划(dynamic programming)
 *
 * Constrains
 * 1. m?n? [1,100]
 *
 * Ideas and Complexities
 * 注意(i==1 || j==1)的条件在这里要做改变
 */


int uniquePathsWithObstacles(int** obstacleGrid, int obstacleGridSize, int* obstacleGridColSize){
    int** map = malloc(sizeof(int*)*obstacleGridSize);
    for (int i=0; i<obstacleGridSize; i++) {
        map[i] = malloc(sizeof(int)*obstacleGridColSize[i]);
    }
    map[0][0] = obstacleGrid[0][0] == 1 ? 0 : 1;
    for (int i=0; i<obstacleGridSize; i++) {
        for (int j=0; j<obstacleGridColSize[i]; j++) {
            if (obstacleGrid[i][j] == 1) map[i][j] = 0;
            else if (i==0 && j==0) map[i][j] = 1;
            else if (i==0) map[i][j] = map[i][j-1];
            else if (j==0) map[i][j] = map[i-1][j];
            else map[i][j] = map[i-1][j] + map[i][j-1];
        }
    }
    return map[obstacleGridSize-1][obstacleGridColSize[obstacleGridSize-1]-1];
}

