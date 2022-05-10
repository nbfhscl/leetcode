import java.lang.Integer;

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
class Solution {
    public boolean hasCycle(ListNode head) {
        return hasCycle_cycle(head);
    }

    public boolean hasCycle_rec(ListNode head) {
        if (head == null) {
            return false;
        }
        if (head.val == Integer.MAX_VALUE) {
            return true;
        }
        head.val = Integer.MAX_VALUE;
        return hasCycle(head.next);
    }

    public boolean hasCycle_cycle(ListNode head) {
        for (; head != null;) {
            if (head.val == Integer.MAX_VALUE) {
                return true;
            }
            head.val = Integer.MAX_VALUE;
            head = head.next;
        }
        return false;
    }

    public boolean hasCycle_tp(ListNode head) {
        ListNode i = head, j = head;
        for (;j != null && j.next != null;) {
            i = i.next;
            j = j.next.next;
            if (i == j) {
                return true;
            }
        }
        return false;
    }

}

