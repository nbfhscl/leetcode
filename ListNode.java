import java.util.Comparator;

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
    public static ListNode from(int[] nodes) {
        return getList(nodes);
    }
    public static void assertEqual(ListNode a, ListNode b) {
        for (;a!=null && b!=null && a.val == b.val; a=a.next, b=b.next);
        assert a==null && b==null;
    }
    public static ListNode merge(ListNode a, ListNode b, Comparator<ListNode> comp) {
        if (comp == null) {
            comp = new Comparator<ListNode>() {
                public int compare(ListNode a, ListNode b) {
                    return a.val - b.val;
                }
            };
        }
        ListNode ret = new ListNode();
        ListNode head = ret;
        for (; ;) {
            if (a!=null && b!=null) {
                if (comp.compare(a,b) < 0) {
                    head.next = a;
                    head = head.next;
                    a = a.next;
                } else {
                    head.next = b;
                    head = head.next;
                    b = b.next;
                }
            } else if (a!=null) {
                head.next = a;
                break;
            } else if (b!=null) {
                head.next = b;
                break;
            } else {
                break;
            }
        }
        return ret.next;
    }
    public static void main (String[] args) {
        assertEqual(merge(from(new int[]{1,3,7}), from(new int[]{2,5,8}), null), from(new int[]{1,2,3,5,7,8}));
    }
}
