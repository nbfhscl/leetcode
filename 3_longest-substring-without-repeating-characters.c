#include <stdlib.h>
#include <string.h>
#include <stdio.h>

/*
 * Constrains
 * 1. sSize?
 * 2. retSize?
 *
 * Ideas and Complexities
 * 1. iteration and hashmap
 *      len[i] = end of longest sub string start with s[i]
 *      if s[e[i]] = s[i+len[i]], e[i]<i+len[i] 
 *      then len[i+j] = len[i] - j, 0 < j < e[i]-i
 *      Time O(n), Memory O(m)
 *
 *      
 */

typedef struct MAP {
    int index;
    int key;
    int value;
} Map;

int get_index(int size, int key) {
    return abs(key)%size + 1;
}

int store_in_map(Map* map, int size, int key, int value) {
    int index = get_index(size, key);
    for (;map[index].index;) index = get_index(size, index);
    map[index].index = index;
    map[index].key = key;
    map[index].value = value;
    return index;
}

int find_in_map(Map* map, int size, int key) {
    int index = get_index(size, key);
    for (;map[index].index && map[index].key != key;) index = get_index(size, index);
    if (map[index].index) return index;
    else return 0;
}

int remove_in_map(Map* map, int size, int key) {
    int index = find_in_map(map, size, key);
    if (index) {
        map[index].index = 0;
        map[index].key = 0;
        map[index].value = 0;
        return index;
    }
    return 0;
}

int lengthOfLongestSubstring_hash(char * s){
    if (s == NULL) return 0;
    int ret = 0;
    int mapSize = 256;
    Map* map = malloc(sizeof(Map) * (mapSize+1));
    memset(map, 0, sizeof(Map) * (mapSize+1)); 
    int i=0, j=0;
    for (; s[j] != '\0'; j++) {
        int index = find_in_map(map, mapSize, s[j]);
        if (!index) {
            store_in_map(map, mapSize, s[j], j);
        } else {
            if (ret < j-i) {
                ret = j-i;
            }
            for (; i < map[index].value; i++) remove_in_map(map, mapSize, s[i]);
            i++;
        }
    }
    return ret > j-i ? ret : j-i;
}

int lengthOfLongestSubstring(char * s){ 
    int ret = 0;
    int map[128] = {0};
    int i=0, j;
    for (j=0; s[j] != '\0'; j++) {
        if (map[s[j]]) {
            ret = ret > j-i ? ret : j-i;
            for (;i<=map[s[j]]-1; i++) map[s[i]] = 0;
        }
        map[s[j]] = j+1;
    }
    return ret > j-i ? ret : j-i;
}

int main() {
    /* char* s = "abcabcbb"; */
    /* char* s = "abc"; */
    char* s = "umvejcuuk";
    printf("%d\n", lengthOfLongestSubstring(s));
}
