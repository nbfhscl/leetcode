#include <stdlib.h>
#include <stdio.h>

/*
 * 实现过程中出现几个低级错误
 * 1. malloc没有memset，用到了index=0的判断必须初始化内存
 * 2. hashmap的size和cap的关系, size = 2*cap, 分配空间=size+1
 * 3. 链表操作应该封装成独立的小函数，更清晰，方便调用
 * 4. free操作只需要把数组释放就行了，不需要根据链表去走
 * 5. 大小要么在链表中处理，要么在hashmap中处理，不能都处理
 * 6. 明确get方法返回的到底是index，key还是value，不要错位了
 */

typedef struct LinkedHashMap {
    int index;
    int key;
    int value;
    struct LinkedHashMap* prev;
    struct LinkedHashMap* next;
} LinkedHashMap;

typedef struct {
    LinkedHashMap* linkedHashMap;
    LinkedHashMap* head;
    LinkedHashMap* tail;
    int mapSize;
    int capacity;
    int size;
} LRUCache;

int get_index(LRUCache* obj, int key) {
    return key % (obj->mapSize) + 1;
}

void add_to_head(LRUCache* obj, LinkedHashMap* t) {
    t->prev = obj->head;
    t->next = obj->head->next;
    obj->head->next->prev = t;
    obj->head->next = t;
    obj->size++;
}

void remove_tail(LRUCache* obj) {
    obj->tail->prev->index = 0;
    obj->tail->prev = obj->tail->prev->prev;
    obj->tail->prev->next = obj->tail;   
    obj->size--;
}

void remove_node(LRUCache* obj, LinkedHashMap* t) {
    t->prev->next = t->next;
    t->next->prev = t->prev;
    obj->size--;
}

LRUCache* lRUCacheCreate(int capacity) {
    LRUCache* lruCache = calloc(sizeof(LRUCache), 1);
    lruCache->capacity = capacity;
    lruCache->mapSize = capacity<<1;
    lruCache->linkedHashMap = calloc(sizeof(LinkedHashMap), (lruCache->mapSize+1));
    lruCache->size = 0;
    lruCache->head = calloc(sizeof(LinkedHashMap), 1);
    lruCache->tail = calloc(sizeof(LinkedHashMap), 1);
    lruCache->head->next = lruCache->tail;
    lruCache->tail->prev = lruCache->head;
    return lruCache;
}

int lRUCacheGet(LRUCache* obj, int key) {
    int index = get_index(obj, key);
    for (;obj->linkedHashMap[index].index && obj->linkedHashMap[index].key != key;)
        index = get_index(obj, index);
    if (obj->linkedHashMap[index].index && obj->linkedHashMap[index].key == key) {
        remove_node(obj, &obj->linkedHashMap[index]);
        add_to_head(obj, &obj->linkedHashMap[index]);
        return obj->linkedHashMap[index].value;
    } else {
        return -1;
    }
}

void lRUCachePut(LRUCache* obj, int key, int value) {
    int index = get_index(obj, key);
    for (;obj->linkedHashMap[index].index && obj->linkedHashMap[index].key != key;)
        index = get_index(obj, index);
    if (obj->linkedHashMap[index].index && obj->linkedHashMap[index].key == key) {
        obj->linkedHashMap[index].value = value;
        remove_node(obj, &obj->linkedHashMap[index]);
        add_to_head(obj, &obj->linkedHashMap[index]);
    } else {
        obj->linkedHashMap[index].index = index;
        obj->linkedHashMap[index].key = key;
        obj->linkedHashMap[index].value = value;
        add_to_head(obj, &obj->linkedHashMap[index]);
        if (obj->size > obj->capacity) remove_tail(obj);
    }
}

void lRUCacheFree(LRUCache* obj) {
    free(obj->linkedHashMap);
    free(obj->head);
    free(obj->tail);
    free(obj);
}

/**
 * Your LRUCache struct will be instantiated and called as such:
 * LRUCache* obj = lRUCacheCreate(capacity);
 * int param_1 = lRUCacheGet(obj, key);
 
 * lRUCachePut(obj, key, value);
 
 * lRUCacheFree(obj);
*/

int main() {
   LRUCache* lru = lRUCacheCreate(2);
   lRUCachePut(lru, 2, 1);
   lRUCachePut(lru, 2, 2);
   printf("%d", lRUCacheGet(lru, 2));
   lRUCachePut(lru, 3, 3);
}
