/**
 *
 * > 难度: Medium | 通过率：55.3% | 通过次数：297485 | 提交次数：537605
 * 
 * `head``left``right``left  。请你反转从位置 `left` 到位置 `right` 的链表节点，返回 **反转后的链表** 。
 * 
 */

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode reverseBetween(ListNode head, int left, int right) {
        if (left >= right) return head;
        ListNode dh = new ListNode(); dh.next = head; head=dh;
        ListNode l, r, ll, rr, last=dh, next;
        int cnt=0;
        for (;cnt++ < left;) {
            last = head;
            head = head.next;
        }
        l = last;
        ll = head;
        last = head;
        head = head.next;
        for (;cnt++ < right;) {
            next = head.next;
            head.next = last;
            last = head;
            head = next;
        }
        next = head.next;
        head.next = last;
        last = head;
        head = next;
        r = head;
        rr = last;
        l.next = rr;
        ll.next = r;
        return dh.next;
    }

    public static void main(String[] args) {
        Solution sl = new Solution();   
        ListNode.assertEqual(sl.reverseBetween(ListNode.from(new int[]{1,2,3,4,5}), 2,4), ListNode.from(new int[]{1,4,3,2,5}));
    }
}

