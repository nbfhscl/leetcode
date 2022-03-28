#include <stdlib.h>
#include <stdio.h>

/*
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


// 1 -> 2 -> 3 -> 4 -> 5 -> NULL
//                          ↑
// 1 -> 2 -> 3 -> 4 -> 5 -> NULL
//                     ↑
// 1 -> 2 -> 3 NULL <- 4 <- 5
//                     ↑
// 1 -> 2 -> 3 -> 4 <- 5
//           ↑
// 1 -> 2 NULL <- 3 <- 4 <- 5
//                ↑
// 1 -> 2 -> 3 <- 4 <- 5
//      ↑
// 1 NULL <- 2 <- 3 <- 4 <- 5
//           ↑
// 1 -> 2 <- 3 <- 4 <- 5
// ↑
// NULL <- 1 <- 2 <- 3 <- 4 <- 5
//         ↑
struct ListNode* reverseList_rc(struct ListNode* head){
    if (head == NULL) return head;
    struct ListNode* newHead = reverseList_rc(head->next);
    if (head->next == NULL) return head;
    head->next->next = head;
    head->next = NULL;
    return newHead;
}


struct ListNode* reverseList_2(struct ListNode* head){
    struct ListNode *prev = NULL;
    struct ListNode *curr = head;
    struct ListNode *next;
    for(;curr != NULL;) {
        next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }
    return prev;
}

struct ListNode* reverseList(struct ListNode* head){
    if (head == NULL) return head;
    if (head->next == NULL) return head;
    struct ListNode *l=head;
    head = head->next;
    l->next = NULL;
    struct ListNode *r = head->next;
    head->next = l;
    for (;r != NULL;) {
        l = head;
        head = r;
        r = head->next;
        head->next = l;
    }
    return head;
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
    struct ListNode* ret = reverseList(listNode);
    for(;ret != NULL;) {
        printf("%d ", ret->val);
        ret = ret->next;
    }
    printf("\n");
    free(listNode);
}

