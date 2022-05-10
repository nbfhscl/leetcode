import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution {
    private TreeNode find_node_in_tree_by_index(TreeNode root, int index) {
        for (; root != null && index != 1; index >>= 1) {
            if (index%2 == 1) {
                root = root.right;
            } else {
                root = root.left;
            }
        }
        return root;
    }
    private int get_node_index_in_tree(TreeNode root, int index, TreeNode p) {
        if (root == null) {
            return 0;
        }
        if (root.val == p.val) {
            return index;
        }
        int ret = 0;
        ret = get_node_index_in_tree(root.left, index<<1, p);
        if (ret != 0) { return ret; }
        ret = get_node_index_in_tree(root.right, (index<<1)+1, p);
        if (ret != 0) { return ret; }
        return 0;
    }
    // 没有通过全部用例，待排查问题
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        int ip, iq;
        ip = get_node_index_in_tree(root, 1, p);
        iq = get_node_index_in_tree(root, 1, q);
        List<Integer> ppList = new ArrayList<>();
        for (; ip != 0;) {
            ppList.add(ip);
            ip >>= 1;
        }
        for (; iq != 0; iq >>= 1){
            if (ppList.contains(iq)) {
                return find_node_in_tree_by_index(root, iq);
            }
        }
        return root;
    }

    /**
     * fl == left child tree contains p or q
     * (fl && fr) || ((x==p || x==q) && (fl || fr))
     * */
    private TreeNode ans;
    private boolean dfs(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null) {
            return false;
        }
        boolean fl = dfs(root.left, p, q);
        boolean fr = dfs(root.right, p, q);
        boolean ret;
        if ((fl && fr) || (root.val == p.val || root.val == q.val) && (fl || fr)) {
            ans = root;
        }
        return fl || fr || (root.val == p.val || root.val == q.val);
    }
    public TreeNode lowestCommonAncestor_better(TreeNode root, TreeNode p, TreeNode q) {
        dfs(root, p, q);
        return ans;
    }

    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.lowestCommonAncestor(TreeNode.getTree(Arrays.asList(3,5,1,6,2,0,8,null,null,7,4), 0), new TreeNode(5), new TreeNode(4)).val == 5;
    }

}

