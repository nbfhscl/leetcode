import java.util.LinkedList;
import java.util.List;

/*
 *
 * > 难度: Medium | 通过率：63.8% | 通过次数：181734 | 提交次数：285070
 * 
 * 给定一个单链表 `L`**的头节点 `head` ，单链表 `L` 表示为：
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
    
    /**
     * O(N), O(N)
     * */
    public void reorderList(ListNode head) {
        ListNode tail = head;
        List<ListNode> list = new LinkedList<>();
        for (;tail != null;) {
            list.add(tail);
            tail = tail.next;
        }
        ListNode temp, last;
        for (int i=0; i<list.size()/2; i++) {
            last = list.get(list.size()-i-1);
            temp = head.next;
            head.next = last;
            last.next = temp;
            head = temp;
        }
        head.next = null;
    }

    /**
     * todo:
     * O(N), O(1)
     * 1. 快慢指针找到中间节点
     * 2. 逆序后半部分链表
     * 3. 合并两部分链表
     */
    public void reorderList_mid_reverse_merge(ListNode head) {
    }

    public static void main(String[] args) {
        Solution sl = new Solution();
        ListNode in = ListNode.from(new int[]{1,2,3,4});
        sl.reorderList(in);
        ListNode.assertEqual(
                in,
                ListNode.from(new int[]{1,4,2,3}));
    }
}

