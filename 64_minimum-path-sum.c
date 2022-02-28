#include <stdlib.h>

/**
 * 中等(medium)
 * 数组(array)
 * 动态规划(dynamic programming)
 * 滚动数组(rolling array) 常规解法用二维数组，滚动数组只需要用一维数组，方法是
 * 一行一行往下计算，始终只需要保存前一行的最小和。
 *
 * Constrains
 * 1. gridSize? [1,200]
 * 2. gridColSize? [1,200]
 * 3. grid[i][j]? [0,100]
 * 4. sum?
 *
 * Ideas and Complexities
 * minPathSum[i][j] = MIN(minPathSum[i-1][j], minPathSum[i][j-1]) + grid[i][j]
 * minPathSum[0][0] = grid[0][0]
 * minPathSum[0][j] = minPathSum[0][j-1]+grid[0][j-1], j>0
 * minPathSum[i][0] = minPathSum[i-1][0]+grid[i-1][0], i>0
 */

#define MIN(a,b) ((a) < (b) ? (a) : (b))

int minPathSum(int** grid, int gridSize, int* gridColSize){
    int** map = malloc(sizeof(int*)*gridSize);
    for (int i=0; i<gridSize; i++) {
        map[i] = malloc(sizeof(int)*gridColSize[i]);
    }
    for (int i=0; i<gridSize; i++) {
        for (int j=0; j<gridColSize[i]; j++) {
            if (i==0 && j==0) map[i][j] = grid[i][j];
            else if (i==0) map[i][j] = map[i][j-1] + grid[i][j];
            else if (j==0) map[i][j] = map[i-1][j] + grid[i][j];
            else map[i][j] = MIN(map[i-1][j], map[i][j-1])+grid[i][j];
        }
    }
    return map[gridSize-1][gridColSize[gridSize-1]-1];
}

