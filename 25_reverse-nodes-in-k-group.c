#include <stdlib.h>

/**
 * 困难(hard)
 * 链表(LinkedList)
 *
 */

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

struct ListNode {
    int val;
    struct ListNode *next;
};

/**
 * h -> 1 -> 2 -> 3 -> 4 -> NULL
 * h -> 1 -> 2 -> NULL
 *
 */
void reverseKLinkedList_swap(struct ListNode* h, int k) {
    struct ListNode *i=h, *j=h;
    for (k--;k>0;) {
        j = i->next;
        i->next = j->next;
        j->next = j->next->next;
        i->next->next = j;
        i = i->next;
        k--;
    }
}
struct ListNode* reverseKGroup_swap(struct ListNode* head, int k){
    if (k == 1) {
        return head;
    }
    struct ListNode* h = malloc(sizeof(*h));
    h->next = head;
    struct ListNode *i = h, *j = h;
    int cnt;
    for (;;) {
        j = i;
        cnt = 0;
        for (;cnt < k && j != NULL;) {
            j = j->next;
            cnt++;
        }
        if (cnt < k || j == NULL) {
            return h->next;
        }
        j = i->next;
        for (; cnt>1;) {
            reverseKLinkedList_swap(i, cnt);
            cnt--;
        }
        i = j;
    }
}

/**
 * dummy 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
 *   |   1 <- 2    3 <- 4    6 <- 5    7
 *   |   |____↑____|____↑    |    ↑    ↑
 *   |        |    |_________|____|    |
 *   |________|              |_________|
 * return dummy->next
 */
struct ListNode* reverseKGroup_reverse_iter(struct ListNode* head, int k){
    struct ListNode *hh, *ll, *mm;
    hh = malloc(sizeof(*hh)), ll = hh;
    struct ListNode *t, *l, *m = head, *r;
    int cnt;
    for (;;) {
        t = m, cnt = k;
        for (;cnt>0 && t != NULL;cnt--) t = t->next;
        if (cnt>0) {
            ll->next = m;
            return hh->next;
        }
        mm = m;
        for (;m != t;) {
            r = m->next;
            m->next = l;
            l = m;
            m = r;
        }
        ll->next = l;
        ll = mm;
    }
}

struct ListNode* reverseKGroup(struct ListNode* head, int k){
    return reverseKGroup_reverse_iter(head, k);
}

