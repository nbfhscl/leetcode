import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

import javax.management.Query;


class TreeNode {
  
  int val;
  TreeNode left;
  TreeNode right;

  TreeNode() {
  }

  TreeNode(int val) {
      this.val = val;
  }

  TreeNode(int val, TreeNode left, TreeNode right) {
         this.val = val;
         this.left = left;
         this.right = right;
     }
}

class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        if (root == null) return null;
        List<List<Integer>> ret = new LinkedList<>();
        Queue<TreeNode> q = new LinkedList<>();
        q.add(root);
        for (; !q.isEmpty();) {
            List<Integer> layer = new LinkedList<>();
            for (int i = q.size(); i>0; i--) {
                TreeNode t = q.poll();
                layer.add(t.val);
                if (t.left != null) {
                    q.add(t.left);
                }
                if (t.right != null) {
                    q.add(t.right);
                }
            }
            ret.add(layer);
        }
        return ret;
    }
}