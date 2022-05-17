import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 * 困难(hard)
 * 1. 归并merge
 * 2. 分治. O(n*logk), O(1)
 * 3. 优先级队列，priority queue. O(n*logk), O(1)
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
    public ListNode mergeKLists_priority_queue(ListNode[] lists) {
        if (lists.length == 0) {
            return null;
        }
        ListNode ret = new ListNode();
        ListNode head = ret;
        PriorityQueue<ListNode> pq = new PriorityQueue<>(
                lists.length,
                new Comparator<ListNode>() {
                    public int compare(ListNode a, ListNode b) {
                        return a.val - b.val;
                    }
                }
        ); 
        for (int i=0;i<lists.length;i++) {
            if (lists[i] != null) {
                pq.add(lists[i]);
            }
        }
        for (;!pq.isEmpty();) {
            head.next = pq.poll();
            head = head.next;
            if (head.next != null) {
                pq.add(head.next);
            }
        }
        return ret.next;
    }

    public ListNode mergeKLists_divide_and_conqure(ListNode[] lists) {
        for (int k=lists.length;k>1;k = k/2 + k%2) {
            for (int i=0, j=k-1;i<j;i++,j--) {
                lists[i] = ListNode.merge(lists[i], lists[j], null);
            }
        }
        return lists[0];
    }
    public static void main(String[] args) {
        Solution sl = new Solution();
        ListNode.assertEqual(sl.mergeKLists_merge(new ListNode[]{
            ListNode.getList(new int[]{3,4,9}),
                ListNode.getList(new int[]{1,2,5})
        }), ListNode.getList(new int[]{1,2,3,4,5,9}));
        ListNode.assertEqual(sl.mergeKLists_priority_queue(new ListNode[]{
            ListNode.getList(new int[]{3,4,9}),
                ListNode.getList(new int[]{1,2,5})
        }), ListNode.getList(new int[]{1,2,3,4,5,9}));
        ListNode.assertEqual(sl.mergeKLists_divide_and_conqure(new ListNode[]{
            ListNode.getList(new int[]{3,4,9}),
                ListNode.getList(new int[]{1,2,5}),
                ListNode.from(new int[]{0,3,10})
        }), ListNode.getList(new int[]{0,1,2,3,3,4,5,9,10}));
    }
}

