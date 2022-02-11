/*
 * 中等(medium)
 * 数组(array)
 * 双指针(double pointer)
 *
 */

/*
 * Constrains
 * 1. How large can n be? 2 <= n <= 105
 * 2. How large can an be? 0 <= an <= 104
 *
 * Ideas and Complexities
 * Find i,j maximize min(ai, aj) * (j-i).
 * O(n2) solution is obvious. But n and an is too large.
 *
 * Test Cases
 */

int maxArea(int* height, int heightSize){
    int max = 0, area;
    for (int i=1; i<heightSize; i++) {
        for (int j=i-1; j>=0; j--) {
            area = (height[i] > height[j] ? height[j] : height[i]) * (i-j);
            if (area > max) {
                max = area;
            }
        }
    }
    return max;
}

int maxArea2(int* height, int heightSize) {
    int l=0, r=heightSize-1, max=0, area;
    while (l < r) {
        area = (height[l] > height[r] ? height[r] : height[l]) * (r-l);
        if (area > max) {
            max = area;
        }
        if (height[l] >= height[r]) {
            r--;
        } else {
            l++;
        }
    }
    return max;
}

