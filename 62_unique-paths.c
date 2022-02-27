#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/**
 * 中等(medium)
 * 数组(array)
 * 动态规划(dynamic programming)
 * 组合(combination)
 *
 * Constrains
 * 1. m? [1,100]
 * 2. n? [1,100]
 * 3. i++ or j++
 * 4. path? [0, 2*10^9]
 *
 * Ideas and Complexities
 * 1. dfs. time consuming.
 * 2. p[i][j] = p[i-1][j] + p[i][j-1]
 * 3. combination. total=m-1+n-1; down=m-1; right=n-1; result = choose m-1 from
 *    m-1+n-1;
 *
 * Test Cases
 * 1. [1,1] => 1;
 * 2. [1,2] => 1;
 * 2. [2,1] => 1;
 * 3. [2,2] => 2;
 * 4. [2,3] => 3;
 * 5. [3,2] => 3;
 * 6. [3,3] => 6;
 */

int uniquePaths_combination(int m, int n) {
    long long ans = 1;
    for (int x=n, y=1; y<m; ++x, ++y) {
        ans = ans * x / y;
    }
    return ans;
}

int uniquePaths_dp_norecursion(int m, int n) {
    int** map = malloc(sizeof(int*)*m);
    for (int i=0; i<m; i++) {
        map[i] = malloc(sizeof(int)*n);
        memset(map[i], 0, sizeof(int)*n);
    }
    for (int i=0; i<m; i++) {
        for (int j=0; j<n; j++) {
            if (i==0 || j==0) {
                map[i][j] = 1;
            } else map[i][j] = map[i-1][j] + map[i][j-1];
        }
    }
    return map[m-1][n-1];
}

int dp(int m, int n, int** map) {
    if (m==0 || n==0) {
        return 1;
    }
    if ( map[m][n] != 0) {
        return map[m][n];
    }
    map[m][n] = dp(m-1,n,map) + dp(m,n-1,map);
    return map[m][n];
}

int uniquePaths_dp(int m, int n) {
    int** map = malloc(sizeof(int*)*m);
    for (int i=0; i<m; i++) {
        map[i] = malloc(sizeof(int)*n);
        memset(map[i], 0, sizeof(int)*n);
    }
    return dp(m-1,n-1,map);
}

void dfs(int m, int n, int* path) {
    if (m==0 && n==0) {
        (*path)++;
        return;
    }
    if (m>0) dfs(m-1,n,path);
    if (n>0) dfs(m,n-1,path);
}

int uniquePaths_dfs(int m, int n){
    int path = 0;
    dfs(m-1, n-1, &path);
    return path;
}
int uniquePaths(int m, int n){
    return uniquePaths_dp(m,n);
}

int main() {
    printf("%d\n", uniquePaths(51,9));
}

