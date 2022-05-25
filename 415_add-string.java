class Solution {
    public String addStrings(String num1, String num2) {
        StringBuilder sb = new StringBuilder();
        int i=num1.length()-1, j= num2.length()-1, k=0, sum=0;
        for (;;) {
            if (i<0 && j<0) {
                if (k != 0) {
                    sb.append(k);
                }
                break;
            } else if (i<0) {
                sum = num2.charAt(j--) - '0' + k;
                sb.append(sum % 10);
                k = sum / 10;
            } else if (j<0) {
                sum = num1.charAt(i--) - '0' + k;
                sb.append(sum % 10);
                k = sum / 10;
            } else {
                sum = num1.charAt(i--) - '0' + num2.charAt(j--) - '0' + k;
                sb.append(sum % 10);
                k = sum / 10;
            }
        }
        return sb.reverse().toString();
    }

    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.addStrings("123", "123").equals("246");
        assert sl.addStrings("1", "1").equals("2");
        assert sl.addStrings("1", "12").equals("13");
        assert sl.addStrings("9", "99").equals("108");
    }
}

