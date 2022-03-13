#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

/*
 * Constrains
 * 1. dividend? 32 bit signed integer
 * 2. divisor? 32 bit signed integer except 0
 * 3. should not use multiplication, division and mod.
 *
 * Ideas and Complexities
 * 1. substraction, brute force. timeout
 *      continuesly substract dividend by divisor until the dividend is less
 *      than or equal to zero.
 *      if the dividend and divisor is negative, should convert them to positive
 *      numbers first.
 * 2. bit shift.
 *      check overflow before shift
 *      Time O(logn)
 * 3. binary search.
 *      x / y = z ==> z*y <= x < (z+1)*y
 *      if x<0 && y<0 then z*y >= x > (z+1)*y
 *      find the maximum z for z*y >= x.
 *      using quick_add_check to check condition, essencially as same as bit
 *      shift.
 *    Time O(log^2n)
 *
 */

int quick_add_check(int x, int y, int z) {
    for (;z>0;) {
        if (z & 1) {
            x -= y;
        }
        if (z != 1) {
            // check if y+y already less than x
            // avoid y+y overflow
            if (y < x - y) {
                return 0;
            }
            y += y;
        }
        z >>= 1;
    }
    return 1;
}
int divide_bs(int dividend, int divisor){
    if (divisor == 1) return dividend;
    if (divisor == -1) {
        if (dividend == INT32_MIN) return INT32_MAX;
        else return -dividend;
    }
    int sign = 1;
    if (dividend > 0) {
        dividend = -dividend;
        sign = -sign;
    }
    if (divisor > 0) {
        divisor = -divisor;
        sign = -sign;
    }
    if (dividend > divisor) {
        return 0;
    }
    int l=1, r=INT32_MAX, m;
    for (;l<r;) {
        m = l + ((r-l+1)>>1);
        if (quick_add_check(dividend, divisor, m)) {
            l = m;
        } else {
            r = m - 1;
        }
    }
    return l*sign;
}

int divide_shift(int dividend, int divisor){
    if (divisor == INT32_MIN && dividend == INT32_MIN) {
        return 1;
    }
    if (divisor == INT32_MIN || dividend == 0) {
        return 0;
    }
    if (divisor == 1) {
        return dividend;
    }
    if (divisor == -1) {
        if (dividend == INT32_MIN) return INT32_MAX;
        else return -dividend;
    }
    int sign = 1;
    int ret = 0;
    if (dividend > 0) {
        dividend = -dividend;
        sign = -sign;
    }
    if (divisor > 0) {
        divisor = -divisor;
        sign = -sign;
    }
    int x = 0, y;
    for (;dividend<=0;) {
        for (x=0, y=divisor; y >= dividend - y;x++, y+=y);
        dividend -= y;
        ret += 1 << x;
    }
    return (ret-1)*sign;
}

int divide_bf(int dividend, int divisor) {
    int sign = 1;
    int ret = 0;
    if (dividend > 0) {
        dividend = -dividend;
        sign = -sign;
    }
    if (divisor > 0) {
        divisor = -divisor;
        sign = -sign;
    }
    for (;dividend<=0;ret++) {
        dividend -= divisor;
    }
    return (ret-1)*sign;
}

int divide(int dividend, int divisor){
    return divide_shift(dividend, divisor);
}

int main(int argc, char* argv[]) {
    /* printf("%d / %d = %d\n", 6, 2, divide(-2147483648, -1)); */
    for (;;) {
        char* dividendStr = malloc(sizeof(char)*20);
        scanf("%s", dividendStr);
        char* divisorStr = malloc(sizeof(char)*20);
        scanf("%s", divisorStr);
        int dividend = atoi(dividendStr);
        int divisor = atoi(divisorStr);
        printf("%d / %d = %d\n", dividend, divisor, divide(dividend, divisor));
    }
}

