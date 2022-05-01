#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <assert.h>

int isContain(const char s[], int l, char t) {
    for (; l>0; l--) {
        if (s[l-1] == t) return l-1;
    }
    return -1;
}

const char open[] = {'(', '[', '{'};
const char clos[] = {')', ']', '}'};

bool isValid(char *s) {
    char stack[10000];
    int si=0;
    int index;
    for (; *s != '\0'; s++) {
        if (-1 != (index = isContain(open, sizeof(open)/sizeof(char), *s))) {
            stack[si++] = open[index];
        } else if (-1 != (index = isContain(clos, sizeof(clos)/sizeof(char), *s))) {
            if (si == 0 || stack[--si] != open[index]) return false;
        }
    }
    if (si != 0) return false;
    return true;
}

bool isValid_simple(char *s) {
    char stack[10000];
    int si=0;
    int index;
    for (; *s != '\0'; s++) {
        if (*s == '(' || *s == '[' || *s == '{') {
            stack[si++] = *s;
        } else if (*s == ')' || *s == ']' || *s == '}') {
            if (si == 0 || stack[--si] != (*s==')'?'(':*s==']'?'[':*s=='}'?'{':'\0')) return false;
        }
    }
    if (si != 0) return false;
    return true;
}

int main() {
    assert(isValid("{[]}\0") == true);
    assert(isValid_simple("{[]}\0") == true);
}
