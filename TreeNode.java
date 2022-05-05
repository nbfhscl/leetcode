import java.util.List;

public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
    public static TreeNode getTree(List<Integer> nodes, int i) {
        if (nodes == null || nodes.isEmpty()) return null;
        if (i > nodes.size()) return null;
        if (nodes.get(i) == null) return null;
        TreeNode node = new TreeNode(nodes.get(i).intValue());
        node.left = getTree(nodes, 2*i+1);
        node.right = getTree(nodes, 2*i+2);
        return node;
    }
}

