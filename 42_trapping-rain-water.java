import java.util.Deque;
import java.util.LinkedList;

/**
 *
 *难度: Hard | 通过率：61.0% | 通过次数：488225 | 提交次数：800615
 *
 *给定 `n` 个非负整数表示每个宽度为 `1` 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 */

class Solution {
    public int trap_dp(int[] height) {
        int size = height.length;
        int[] left = new int[size];
        int[] right = new int[size];
        left[0] = 0; right[size-1] = 0;
        for (int i=1; i<size; i++) {
            left[i] = Integer.max(left[i-1], height[i-1]);
        }
        for (int i=size-2; i>=0; i--) {
            right[i] = Integer.max(right[i+1], height[i+1]);
        }
        int ret=0;
        for (int i=0; i<size; i++) {
            ret += Integer.max(0, (Integer.min(left[i],right[i]) - height[i]));
        }
        return ret;
    }

    public int trap_tp(int[] height) {
        int ret = 0;
        int l=0, r=height.length-1;
        int maxLeft=0, maxRight=0;
        for (;l<=r;) {
            if (maxLeft < maxRight) {
                ret += Integer.max(maxLeft-height[l], 0);
                maxLeft = Integer.max(maxLeft, height[l]);
                l++;
            } else {
                ret += Integer.max(maxRight-height[r], 0);
                maxRight = Integer.max(maxRight, height[r]);
                r--;
            }
        }
        return ret;
    }

    public int trap_monotone_stack(int[] height) {
        int ret=0;
        Deque<Integer> stack = new LinkedList<>();
        for (int i=0; i<height.length; i++) {
            if (stack.isEmpty()) {
                stack.addFirst(i);
                continue;
            }
            if (height[stack.getFirst()] >= height[i]) stack.addFirst(i);
            else {
                for (;!stack.isEmpty() && height[stack.getFirst()]<height[i];) {
                    // 注意这里for循环里面要get，进来后才poll
                    int h = stack.pollFirst();
                    if (stack.isEmpty()) break;;
                    ret += (Integer.min(height[stack.getFirst()], height[i])-height[h]) * (i-stack.getFirst()-1);
                }
                stack.addFirst(i);
            }
        }
        return ret;
    }

    public int trap(int[] height) {
        return trap_monotone_stack(height);
    }
    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.trap(new int[]{0,1,0,2,1,0,1,3,2,1,2,1}) == 6;
        assert sl.trap(new int[]{4,2,0,3,2,5}) == 9;
    }
}

