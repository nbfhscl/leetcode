import javax.swing.tree.TreeNode;

/**
 *
 * > 难度: Hard | 通过率：45.0% | 通过次数：234232 | 提交次数：520232
 * 
 * **路径** 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 **至多出现一次** 。该路径** 至少包含一个 **节点，且不一定经过根节点。
 * 
 * **路径和** 是路径中各节点值的总和。
 * 
 * 给你一个二叉树的根节点 `root` ，返回其 **最大路径和** 。
 * */

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private int max = -1001;
    public int maxPathSum(TreeNode root) {
        maxPathSum_rc(root);
        return max;
    }
    public int maxPathSum_rc(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int maxLeft = maxPathSum_rc(root.left);
        int maxRight = maxPathSum_rc(root.right);
        maxLeft = Integer.max(0,maxLeft);
        maxRight = Integer.max(0,maxRight);
        int maxThis = maxLeft + maxRight + root.val;
        max = Integer.max(max, maxThis);
        return root.val + Integer.max(maxLeft, maxRight);
    }

    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.maxPathSum(TreeNode.from(new Integer[]{1,2,3}, 0)) == 6;
    }
}

