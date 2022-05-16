/**
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 * 困难(hard)
 * 1. 归并merge
 * 2. todo: 分治. O(n*logk), O(1)
 * 3. todo: 优先级队列，priority queue. O(n*logk), O(1)
 */
class Solution {
    // O(kn), O(1)
    public ListNode mergeKLists_merge(ListNode[] lists) {
        ListNode ret = new ListNode();
        ListNode head = ret;
        for (;;) {
            int which = -1;
            for (int i=0, min=20000; i<lists.length; i++) {
                if (lists[i] != null && lists[i].val < min) {
                    min = lists[i].val;
                    which = i;
                }
            }
            if (which == -1) {
                break;
            }
            head.next = lists[which];
            head = head.next;
            lists[which] = lists[which].next;
        }
        return ret.next;
    }
    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.mergeKLists_merge(new ListNode[]{
                ListNode.getList(new int[]{3,4,9}),
                ListNode.getList(new int[]{1,2,5})
        }).equals(ListNode.getList(new int[]{1,2,3,4,5,9}));
    }
}

