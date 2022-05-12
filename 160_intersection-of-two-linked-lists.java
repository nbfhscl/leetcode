
/**
 * 给你两个单链表的头节点 `headA` 和 `headB` ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 `null` 。
 * 图示两个链表在节点 `c1` 开始相交**：**
 *
 * Easy(简单)
 * LinkedList(链表)
 *
 * Constrains
 * 1. no modify to input list
 * 2. how long the lists could be? [1,30000]
 * 3. list node val? [1,100000]
 * 4. node1.val == node2.val ? happens
 * */

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode la = headA, lb = headB;
        int sa=1, sb=1;
        for (; la.next != null; ) {
            la = la.next;
            sa++;
        }
        for (; lb.next != null; ) {
            lb = lb.next;
            sb++;
        }
        if (la != lb) {
            return null;
        }
        if (sa > sb) {
            for (;sa>sb;) {
                headA = headA.next;
                sa--;
            }
        } else {
            for (;sa<sb;) {
                headB = headB.next;
                sb--;
            }
        }
        for (; sa > 0;sa--) {
            if (headA == headB) {
                return headA;
            }
            headA = headA.next;
            headB = headB.next;
        }
        return null;
    }
}

