import java.util.HashMap;
import java.util.Map;

/**
 * Medium(中等)
 * HashMap
 * Two Pointer
 *
 *给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 * 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
 * 不允许修改 链表。
 *
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
public class Solution {
    public ListNode detectCycle(ListNode head) {
        return detectCycle_tp(head);
    }
    
    public ListNode detectCycle_hashmap(ListNode head) {
        Map<ListNode, Boolean> map = new HashMap<>();
        for (;head != null;) {
            if (map.containsKey(head)) {
                return head;
            }
            map.put(head, true);
            head = head.next;
        }
        return null;
    }

    public ListNode detectCycle_tp(ListNode head) {
        if (head == null) {
            return null;
        }
        if (head.next == null) {
            return null;
        }
        ListNode i=head, j=head;
        for (;j != null && j.next != null;) {
            j = j.next.next;
            i = i.next;
            if (i == j) {
                break;
            }
        }
        if (i != j) {
            return null;
        }
        ListNode k=head;
        for (;;) {
            if (k == i) {
                return k;
            }
            k = k.next;
            i = i.next;
        }
    }
}

