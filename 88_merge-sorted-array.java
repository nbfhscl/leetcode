class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int i=m-1,j=n-1,k=m+n-1;
        for (;k >= 0;) {
            if (i<0) {
                nums1[k--] = nums2[j--];
            } else if (j<0) {
                nums1[k--] = nums1[i--];
            } else if (nums1[i] > nums2[j]) {
                nums1[k--] = nums1[i--];
            } else {
                nums1[k--] = nums2[j--];
            }
        }
    }

    public static void main(String[] args) {
        Solution sl = new Solution();
        int[] nums1 = new int[]{1,3,5,0,0,0};
        sl.merge(nums1,3, new int[]{2,4,6}, 3);
        for (int i=0;i<nums1.length;i++){
            System.out.println(nums1[i]);
        }
    }
}

