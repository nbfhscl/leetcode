#include <stdlib.h>
#include <stdio.h>

/*
 * review reverse linkedlist
 *
 *
 * Constrains
 * 1. ListNodeSize? [0,5000]
 * 2. ListNode[i]? [-5000,5000]
 *
 * Ideas and Complexities
 * 1. iteration
 *      reverse the node
 * 2. recursion
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


// 1 -> 2 -> 3 -> 4 -> 5
// 1 -> 2 -> 3    4 <- 5
//                ↓
//               NULL
// NULL 1 <- 2 <- 3 <- 4 <- 5
struct ListNode* reverseList_rc(struct ListNode* head){
    if (head->next == NULL) return head;
    struct ListNode *ret = reverseList_rc(head->next);
    head->next->next = head;
    head->next = NULL;
    return ret;
}

/**
 *       1 -> 2 -> 3 -> NULL;
 *  l    m    r
 *
 *    r = m.next
 *  m.next = l
 *    l = m
 *  m = r
 *          
 */
struct ListNode* reverseList_2(struct ListNode* head){
    struct ListNode *l = NULL;
    struct ListNode *m = head;
    struct ListNode *r;
    for(;m != NULL;) {
        r = m->next;
        m->next = l;
        l = m;
        m = r;
    }
    return l;
}

int main() {
    struct ListNode *listNode;
    listNode = malloc(sizeof(*listNode)*5);
    listNode[0] = (struct ListNode){1, listNode+1};
    listNode[1] = (struct ListNode){2, listNode+2};
    listNode[2] = (struct ListNode){3, listNode+3};
    listNode[3] = (struct ListNode){4, listNode+4};
    listNode[4] = (struct ListNode){5, listNode+5};
    listNode[4].next = NULL;
    struct ListNode* ret = reverseList_rc(listNode);
    for(;ret != NULL;) {
        printf("%d ", ret->val);
        ret = ret->next;
    }
    printf("\n");
    free(listNode);
}

