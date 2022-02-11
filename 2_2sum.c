#include <stdlib.h>
#include <string.h>

/*
 * Constrains
 * 1. Value involved is all integer
 * 2. one solution or none
 * 3. no reusing
 * 4. the order of returned array does not matters
 * 5. how large can each values be and how long can the array be?
 * 6. what should I return if there is not any solution found?
 * 7. will value repeat in array?
 *
 * Ideas and Complexities
 * 1. double cycle checking
 *    time: n2. memory: 1.
 * 2. a hashmap stores all value, go through the array to check if
 *    (target-current) exists in hashmap.
 *    time: n. memory: n.
 * 3. go through the array while store and checking in the mean time.
 *    time: n. memory: n.
 *
 * Test Cases
 * 1. [1, 2, 3, 4, 5], 9 => [4, 5]
 * 2. [30, 33, 87, 108], 120 => [1, 2]
 * 3. [2, 2], 4 => [0, 1]
 * 4. [1, 2, 3, 4, 5], 10 => none
 * 5. [1, 2], 2 => none
 * 6. [1], 1 => none
 * 7. [], 0 => none
 * 8. [0,4,3,0], 0 => [0,3], this is a special case for hashmap implementation.
 * 9. [2,5,5,1], 10 => [1,2], this too, lead to heap overflow
 *
 * Summary
 * when implementing hashmap, I encountered folling problems:
 * 1. what should be on heap while others be on stack?
 * 2. edge cases of first and last space of the map buffer. First space have an
 *    index 0, key 0 and value 0, while the value to be stored could also be 0,
 *    the key could be 0, the index hashed from the key could be 0. This could
 *    mess everything up if you take this values as checking condition while do
 *    not consider this edge cases. For example, if you trying to check whether
 *    place is takened by comparing index to 0, you should ensure buffer[0]
 *    never be used. If you trying to check the value is already stored by
 *    comparing provided value to the found value, you should also ensure the
 *    found value is actually found, that means to check the index not 0 too.
 *    Last space is always related to buffer overflow.
 *
 */

struct HashNode {
    int index;
    int key;
    int value;
};

struct HashMap {
    int size;
    struct HashNode* node;
};

int calculate_hash_index(int size, int key) {
    return abs(key)%size+1;
}

int found_in_hashmap(struct HashMap* hm, int key) {
    int index = calculate_hash_index(hm->size, key);
    while (hm->node[index].index && hm->node[index].key != key) {
        index = calculate_hash_index(hm->size, index);
    }
    if (hm->node[index].index && hm->node[index].key == key) {
        return index;
    }
    return -1;
}

int store_in_hashmap(struct HashMap* hm, int key, int value) {
    int index = calculate_hash_index(hm->size, key);
    while (hm->node[index].index) {
        index = calculate_hash_index(hm->size, index);
    }
    /* hm->node[index] = (struct HashNode){.index=index, .key=key, .value=value}; */
    /* direct assignment will save a little memory. Testing on leetcode shows
     * about 0.3aM */
    hm->node[index].index = index;
    hm->node[index].key = key;
    hm->node[index].value = value;
    return index;
}

int free_hashmap(struct HashMap* hm) {
    free(hm->node);
    /* free(hm); */
    return 0;
}

int* twoSum_n1(int* nums, int numsSize, int target, int* returnSize) {
    struct HashMap hm = {.size=numsSize*2};
    hm.node = calloc(sizeof(struct HashNode), hm.size+1);
    for (int i=0;i<numsSize;i++) {
        int index = found_in_hashmap(&hm, target-nums[i]);
        if (index != -1) {
            int* res = (int*)malloc(sizeof(int)*2);
            res[0] = hm.node[index].value, res[1] = i;
            *returnSize = 2;
            free_hashmap(&hm);
            return res;
        } else {
            store_in_hashmap(&hm, nums[i], i);
        }
    }
    *returnSize = 0;
    free_hashmap(&hm);
    return 0;
}

int* twoSum_n2(int* nums, int numsSize, int target, int* returnSize){
    for (int i=0;i<numsSize;i++) {
        for (int j=i+1;j<numsSize;j++) {
            if (*(nums+i) + *(nums+j) == target) {
                int *res = (int*)malloc(sizeof(int)*2);
                *res = i;
                *(res+1) = j;
                *returnSize = 2;
                return res;
            }
        }
    }
    *returnSize = 0;
    return 0;
}


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    return twoSum_n1(nums, numsSize, target, returnSize);
}

