
/**
 * Constrains
 * 1. numsSize? [1, 5000]
 * 2. nums[i]? [-10000, 10000]
 * 3. nums[i] == nums[j], i != j? not exist
 * 4. target? [-10000, 10000]
 *
 * Ideas and Complexities
 * 1. 利用数组的有序性和二分查找进行排除
 */

class Solution {
    public int search(int[] nums, int target) {
        int l=0, r=nums.length-1, mid=0, v=0;
        for (; l<=r;) {
            mid = l + ((r-l)>>1);
            v = nums[mid];
            if (v == target) {
                return mid;
            }
            if (v >= nums[l]) {
                if (target >= nums[l] && target < v) {
                    r = mid - 1;
                } else {
                    l = mid + 1;
                }
            } else {
                if (target > v && target <= nums[r]) {
                    l = mid + 1;
                } else {
                    r = mid - 1;
                }
            }
        }
        return -1;
    }
    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.search(new int[]{4,5,6,7,0,1,2}, 0) == 4;
        assert sl.search(new int[]{4,5,6,7,0,1,2}, 9) == -1;
        assert sl.search(new int[]{}, 0) == -1;
        assert sl.search(new int[]{1}, 0) == -1;
        assert sl.search(new int[]{1}, 1) == 0;
        // when v == nums[l], v is located in left part
        assert sl.search(new int[]{3,1}, 1) == 1;
    }
}

