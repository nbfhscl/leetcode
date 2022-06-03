import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution {
    private void dfs(int[] nums, List<List<Integer>> ret, List<Integer> selected) {
        if (selected.size() == nums.length) {
            ret.add(new ArrayList<>(selected));
        }
        for (int i=0;i<nums.length;i++) {
            if (!selected.contains(nums[i])) {
                selected.add(nums[i]);
                dfs(nums, ret, selected);
                selected.remove(selected.size()-1);
            }
        }
    }
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> ret = new ArrayList<>();
        List<Integer> selected = new ArrayList<>(nums.length);
        dfs(nums, ret, selected);
        return ret;
    }

    public static void main (String[] args) {
        Solution sl = new Solution();
        assert sl.permute(new int[]{1,2,3}).equals(Arrays.asList(
                    Arrays.asList(1,2,3),
                    Arrays.asList(1,3,2),
                    Arrays.asList(2,1,3),
                    Arrays.asList(2,3,1),
                    Arrays.asList(3,1,2),
                    Arrays.asList(3,2,1)));
    }
}

