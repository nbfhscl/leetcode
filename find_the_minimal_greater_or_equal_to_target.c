#include <assert.h>

/*
 * review binarysearch
 * 左边取右边，右边取左边
 * 左边左边等，右边右边等
 * 有等也有等，没等也没等
 */

int find_the_minimal_greater_than_or_equal_to_target(int* a, int size, int target) {
    int l=0, r=size-1, m;
    for (;l<r;) {
        m = l + ((r-l)>>1);
        if (a[m] == target) {
            r = m;
        } else if (a[m] < target) {
            l = m + 1;
        } else {
            r = m;
        }
    }
    if (a[l] >= target) return l;
    else return -1;
}

int main() {
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){0,1,2,3,4,5}, 6, 3) == 3); 
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){0,2,2,3,3,5}, 6, 3) == 3);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){2,2,2,2,2,2}, 6, 3) == -1);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){2,2,2,2,2,2}, 6, 2) == 0);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){2,2,2,2,2,2}, 6, 3) == -1);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){2,2,2,2,2,2}, 6, 1) == 0);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){1,2,3,3,5,5}, 6, 1) == 0);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){1,2,3,3,5,5}, 6, 0) == 0);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){1,2,3,3,5,5}, 6, 2) == 1);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){1,2,3,3,5,5}, 6, 5) == 4);
    assert(find_the_minimal_greater_than_or_equal_to_target
            ((int[6]){1,2,3,3,5,5}, 6, 6) == -1);
}
