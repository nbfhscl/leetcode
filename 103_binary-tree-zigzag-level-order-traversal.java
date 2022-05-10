import java.util.ArrayList;
import java.util.Arrays;
import java.util.Deque;
import java.util.LinkedList;
import java.util.List;

/**
 * Constrains
 * 1. treeSize? [0,2000]
 * 2. val? [-100, 100]
 * Ideas and Complexities
 * out queue order| tree | in queue order of children
 * -|-
 * -> | root null | <-
 * <- | null l r | ->
 * -> | l r l r null| <-
 * <-
 *      1
 *   2     3
 *  4 5   6 7
 * 1 null
 * 从左往右取，左儿子再右儿子，null插最前面
 * null 2 3
 * 从右往左取，右儿子再左儿子，null插最后面
 * 4 5 6 7 null
 * 
 * 其实是一个双向队列，中间有一个null做分割，然后左边放一层，右边放一层。
 * 从左往右取的那层放左右，从右往左取的那一层放右边
 * 取到中间的null则一层取完了。
 * 
 * Test Cases
 * 1. [] => []
 * */

class Solution {

    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> ret = new LinkedList<>();
        if (root == null) return ret;
        // init for root
        Deque<TreeNode> dq = new LinkedList<>();
        // 先放左半边，null放中间
        dq.addLast(root);dq.addLast(null);
        // 用来放root.val
        List<Integer> rett = new ArrayList<>(1);
        // 左半边 or 右半边，root放在左半边
        boolean left = true;
        TreeNode node, node2;
        for (; !dq.isEmpty();) {
            // 左半边从左往右取
            node = left ? dq.pollFirst() : dq.pollLast();
            if (node == null) {
                ret.add(rett);
                // 如果没有下一层了
                if (dq.isEmpty()) break;
                // 下一层
                rett = new ArrayList<>(dq.size());
                if (left) {
                    // 左半边取完了要在中间放null
                    dq.addFirst(null);
                } else {
                    // 右半边取完了要在中间放null
                    dq.addLast(null);
                }
                left = !left;
            } else {
                rett.add(node.val);
                // 左半边从左往右取，先取左儿子
                // 右半边从右往左取，先取右儿子
                node2 = left ? node.left : node.right;
                if (node2 != null) {
                    if (left) {
                        // 左半边的儿子要放右半边
                        dq.addLast(node2);
                    } else {
                        // 右半边的儿子要放左半边
                        dq.addFirst(node2);
                    }
                }
                node2 = left ? node.right : node.left;
                if (node2 != null) {
                    if (left) {
                        dq.addLast(node2);
                    } else {
                        dq.addFirst(node2);
                    }
                }
            }
        }
        return ret;
    }

    public static void main(String[] args) {
        List<Integer> list = Arrays.asList(3, 9, 20, null, null, 15, 7);
       Solution sl = new Solution();
        System.out.println(sl.zigzagLevelOrder(TreeNode.getTree(list, 0)));
    }
}

