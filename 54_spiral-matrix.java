import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * > 难度: Medium | 通过率：48.8% | 通过次数：264956 | 提交次数：543013
 * 
 * 给你一个 `m` 行 `n` 列的矩阵 `matrix` ，请按照 **顺时针螺旋顺序** ，返回矩阵中的所有元素。
 */

class Solution {
    public List<Integer> spiralOrder_tblr(int[][] matrix) {
        List<Integer> ret = new ArrayList<>();
        int t=0,b=matrix.length-1;
        int l=0,r=matrix[0].length-1;
        for (;;) {
            for (int i=l; i<=r; i++) ret.add(matrix[t][i]);
            if (++t>b) break;
            for (int i=t; i<=b; i++) ret.add(matrix[i][r]);
            if (--r<l) break;
            for (int i=r; i>=l; i--) ret.add(matrix[b][i]);
            if (--b<t) break;
            for (int i=b; i>=t; i--) ret.add(matrix[i][l]);
            if (++l>r) break;
        }
        return ret;
    }

    public List<Integer> spiralOrder_k(int[][] matrix) {
        List<Integer> ret = new ArrayList<>();
        int i=0,j=-1,k=0;
        Boolean end;
        for (;;) {
            end = true;
            for (;j<matrix[i].length-k-1;) {
                ret.add(matrix[i][++j]);
                end = false;
            }
            if (end) {
                break;
            }
            end = true;
            for (;i<matrix.length-k-1;) {
                ret.add(matrix[++i][j]);
                end = false;
            }
            if (end) {
                break;
            }
            end = true;
            for (;j>0+k;) {
                ret.add(matrix[i][--j]);
                end = false;
            }
            if (end) {
                break;
            }
            end = true;
            for (;i>1+k;) {
                ret.add(matrix[--i][j]);
                end = false;
            }
            if (end) {
                break;
            }
            k++;
        }
        return ret;
    }

    public List<Integer> spiralOrder(int[][] matrix) {
        return spiralOrder_tblr(matrix);
    }
    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.spiralOrder(new int[][]{{1,2,3},{8,9,4},{7,6,5}}).equals(Arrays.asList(1,2,3,4,5,6,7,8,9));
    }
}

