#include <stdint.h>
#include <stdio.h>

/*
 * Constrains
 * 1. n? [-2^32, 2^32-1]
 * 2. x? (-100.0,100.0)
 * 3. x^n [-10000,10000]
 *
 * Ideas and Complexities
 * 1. quick power 
 *
 */

double myPow(double x, int n){
    double ret=1;
    // use type long to handle special case of INT32_MIN
    long nn = n < 0 ? -(long)n : n;
    for (;nn>0;) {
        if (nn & 1) ret *= x;   
        x *= x;
        nn >>= 1;
    }
    return n < 0 ? 1/ret : ret;
}

int main() {
    double x = 2;
    int n = -1;
    printf("%d power of %lf = %lf\n", n, x, myPow(x, n));
}

