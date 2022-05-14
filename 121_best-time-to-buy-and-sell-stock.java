class Solution {
    public int maxProfit(int[] prices) {
        int max = 0;
        int min[] = new int[prices.length];
        min[0] = 20000;
        for (int i=1; i<prices.length; i++) {
            min[i] = Math.min(min[i-1], prices[i-1]);
            max = Math.max(prices[i]-min[i], max);
        }
        return max;
    }

    public static void main(String[] args) {
        Solution sl = new Solution();
        assert(sl.maxProfit(new int[]{7,1,5,3,6,4}) == 5);
    }
}
