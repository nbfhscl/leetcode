import java.util.HashSet;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Set;

/**
 * 岛屿数量
 * 给你一个由 `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 * 此外，你可以假设该网格的四条边均被水包围。

 * Medium(中等)
 *
 * Ideas and Complexities
 * 1. dfs
 * 2. bfs
 * 3. 并查集(joint search set)
 *
 * */

class Solution {
    private void dfs(char[][] grid, int i, int j) {
        if (i<0 || j<0 || i>=grid.length || j>=grid[i].length) {
            return;
        }
        if (grid[i][j] == '0') {
            return;
        }
        grid[i][j] = '0';
        dfs(grid, i+1, j);
        dfs(grid, i-1, j);
        dfs(grid, i, j+1);
        dfs(grid, i, j-1);
    }
    public int numIslands_dfs(char[][] grid) {
        int ret=0;
        for (int i=0; i<grid.length; i++) {
            for (int j=0; j<grid[i].length; j++) {
                if (grid[i][j] == '1') {
                    dfs(grid, i, j);
                    ret++;
                }
            }
        }
        return ret;
    }

    class Pair {
        int i;
        int j;
        Pair(int i, int j) {
            this.i = i;
            this.j = j;
        }
    }
    public int numIslands_bfs(char[][] grid) {
        int ret = 0;
        for (int i=0; i<grid.length; i++) {
            for (int j=0; j<grid[i].length; j++) {
                if (grid[i][j] == '1') {
                    Queue<Pair> q = new LinkedList<>();
                    grid[i][j] = '0';
                    q.add(new Pair(i,j));
                    for (;!q.isEmpty();) {
                        Pair p = q.poll();
                        if (p.i+1<grid.length && grid[p.i+1][p.j] == '1') {
                            // must set 0 when put in instead of when poll out
                            // otherwise an infinite loop could happen
                            grid[p.i+1][p.j] = '0';
                            q.add(new Pair(p.i+1, p.j));
                        }
                        if (p.i-1>=0 && grid[p.i-1][p.j] == '1') {
                            grid[p.i-1][p.j] = '0';
                            q.add(new Pair(p.i-1, p.j));
                        }
                        if (p.j+1<grid[p.i].length && grid[p.i][p.j+1] == '1') {
                            grid[p.i][p.j+1] = '0';
                            q.add(new Pair(p.i, p.j+1));
                        }
                        if (p.j-1>=0 && grid[p.i][p.j-1] == '1') {
                            grid[p.i][p.j-1] = '0';
                            q.add(new Pair(p.i, p.j-1));
                        }
                    }
                    ret++;
                }
            }
        }
        return ret;
    }

    
    private void merge(Pair p1, Pair p2, Pair[][] fa) {
        fa[p1.i][p1.j] = fa[p2.i][p2.j];
    }
    // 并查集，以下代码未完成，带实现
    public int numIslands_jsc(char[][] grid) {
        int rs = grid.length, cs = grid[0].length;
        Pair[][] fa = new Pair[rs][cs];
        for (int i=0; i<rs; i++) {
            for (int j=0; j<cs; j++) {
                fa[i][j] = new Pair(i,j);
            }
        }
        for (int i=0; i<rs; i++) {
            for (int j=0; j<cs; j++) {
                if (grid[i][j] == '1') {
                    grid[i][j] = '0';
                    if (i+1<rs && grid[i+1][j] == '1') {
                        merge(new Pair(i+1, j), new Pair(i,j), fa);
                    }
                    if (i-1>=0 && grid[i-1][j] == '1') {
                        merge(new Pair(i-1, j), new Pair(i,j), fa);
                    }
                    if (j+1<rs && grid[i][j+1] == '1') {
                        merge(new Pair(i, j+1), new Pair(i,j), fa);
                    }
                    if (j-1>=0 && grid[i][j-1] == '1') {
                        merge(new Pair(i, j-1), new Pair(i,j), fa);
                    }
                }
            }
        }
        Set<Pair> set = new HashSet<>();
        for (int i=0; i<rs; i++) {
            for (int j=0; j<cs; j++) {
                set.add(fa[i][j]);
            }
        }
        return set.size();
    }

    private int numIslands(char[][] grid) {
        return numIslands_jsc(grid);
    }

    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.numIslands(new char[][]{{'1','1','1'},{'0','1','0'},{'1','1','1'}}) == 1;
        assert sl.numIslands(new char[][]{{'1','0','0'},{'0','1','0'},{'0','0','1'}}) == 3;
        assert sl.numIslands(new char[][]{{'1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','0','1','0','1','1'},{'0','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','0'},{'1','0','1','1','1','0','0','1','1','0','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','0','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','0','1','1','1','1','1','1','0','1','1','1','0','1','1','1','0','1','1','1'},{'0','1','1','1','1','1','1','1','1','1','1','1','0','1','1','0','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','0','1','1'},{'1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'0','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','0','1','1','1','1','1','1','1','0','1','1','1','1','1','1'},{'1','0','1','1','1','1','1','0','1','1','1','0','1','1','1','1','0','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','0'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','0','0'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'}}) == 1;
    }
}

