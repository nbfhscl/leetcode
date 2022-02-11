/*
 * @lc app=leetcode.cn id=1 lang=java
 *
 * [1] 两数之和
 */

// @lc code=start
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> m = new HashMap<>();
        for (int i=0;i<nums.length;i++) {
            if (m.containsKey(nums[i])) {
                return new int[]{m.get(nums[i]), i};
            } else {
                m.put(target-nums[i], i);
            }
        }
        return null;
    }
}
// @lc code=end

