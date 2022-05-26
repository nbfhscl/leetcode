/**
 *
 * > 难度: Medium | 通过率：53.5% | 通过次数：527004 | 提交次数：985435
 *
 * 给你一个整数数组 `nums` ，找到其中最长严格递增子序列的长度。
 *子序列 **是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，`[3,6,2,7]` 是数组 `[0,3,1,6,2,2,7]` 的子序列。
 *
 * */
class Solution {
    public int lengthOfLIS_dp(int[] nums) {
        int ret = 1;
        int []dp = new int[nums.length];
        dp[0] = 1;
        for (int i=1;i<nums.length;i++) {
            int max = 0;
            for (int j=i-1;j>=0;j--) {
                if (nums[j]<nums[i] && dp[j]>max) {
                    max = dp[j];
                }
            }
            dp[i] = max+1;
            if (dp[i] > ret) {
                ret = dp[i];
            }
        }
        return ret;
    }

    private int find_index_of_the_maximum_less_than_target(int target, int[] nums, int r) {
        int l=0;
        for (;l<r;) {
            int m = l + ((r-l+1)>>1);
            int v = nums[m];
            if (v == target) {
                r = m-1;
            } else if (v < target) {
                l = m;
            } else {
                r = m-1;
            }
        }
        if (nums[l] < target) return l;
        return -1;
    }
    public int lengthOfLIS_bs(int[] nums) {
        int ret=0;
        int min[] = new int[nums.length];
        min[0] = nums[0];
        for (int i=1; i<nums.length; i++) {
            if (nums[i] > min[ret]) {
                min[++ret] = nums[i];
            } else if (nums[i] == min[ret]) {
                continue;
            } else{
                min[find_index_of_the_maximum_less_than_target(nums[i], min, ret)+1] = nums[i];
            }
        }
        return ret+1;
    }

    public int lengthOfLIS(int[] nums) {
        return lengthOfLIS_bs(nums);      
    }
    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.lengthOfLIS(new int[]{0,3,1,6,2,2,7}) == 4;
        assert sl.lengthOfLIS(new int[]{0,1,0,3,2,3}) == 4;
    }
}

