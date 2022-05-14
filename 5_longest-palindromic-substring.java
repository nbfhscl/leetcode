/**
 * Medium(中等)
 * 
 * Constrains
 * 1. length? [1,1000]
 * 2. charset? number and english letter
 *
 * Ideas and Complexities
 * 1. traversal through the string, try every character as center, search to both sides for longest equal substring, if there are duplicated characters, make them all as center.
 * 2. todo: dynamic programming
 * 3. todo: Manacher algorithm
 *
 */

class Solution {
    public String longestPalindrome(String s) {
        char[] cList = s.toCharArray();
        int len = cList.length;
        int rl=0, rr=0;
        // m<len-1 cant parse “a" -> "a"
        for (int m=0,l=0,r=0;m<len;) {
            // l=m-1 hard to handle duplicated char at index 0
            l=m; r=m+1;
            for (;r<len && cList[r] == cList[m];r++);
            m = r;
            for (;l>0 && r<len && cList[l-1] == cList[r];l--,r++);
            if (r-l > rr-rl) {
                rl=l;
                rr=r;
            }
        }
        return String.copyValueOf(cList, rl, rr-rl);
    }
    public static void main (String[] args) {
        Solution sl = new Solution();
        assert sl.longestPalindrome("abcaba").equals("aba");
        assert sl.longestPalindrome("a").equals("a");
    }
}

