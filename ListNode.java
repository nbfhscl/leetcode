public class ListNode {
    int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    public static ListNode getList(int[] nodes) {
        ListNode ret = new ListNode();
        ListNode head = ret;
        for (int i=0; i<nodes.length; i++) {
            head.next = new ListNode(nodes[i]);
            head = head.next;
        }
        return ret.next;
    }
}
