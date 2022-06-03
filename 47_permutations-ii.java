import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution {
    private void dfs(int[] nums, List<Integer> rett, boolean[] selected, List<List<Integer>> ret) {
        if (rett.size() == nums.length) {
            ret.add(new ArrayList<>(rett));
        }
        boolean[] selectedThisPlace = new boolean[50];
        for (int i=0; i<nums.length; i++) {
            if (!selectedThisPlace[nums[i]] && !selected[i]) {
                selected[i] = true;
                rett.add(nums[i]);
                dfs(nums, rett, selected, ret);
                rett.remove(rett.size()-1);
                selected[i] = false;
                selectedThisPlace[nums[i]+25] = true;
            }
        }
    }
    public List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> ret = new ArrayList<>();
        List<Integer> rett = new ArrayList<>(nums.length);
        boolean[] selected = new boolean[nums.length];
        dfs(nums, rett, selected, ret);
        return ret;
    }

    public static void main (String[] args) {
        Solution sl = new Solution();
        assert sl.permuteUnique(new int[]{1,1,3}).equals(Arrays.asList(
                    Arrays.asList(1,1,3),
                    Arrays.asList(1,3,1),
                    Arrays.asList(3,1,1)));
    }
}

