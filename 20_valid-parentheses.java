import java.util.Arrays;
import java.util.Deque;
import java.util.LinkedList;
import java.util.List;

class Solution {
    public boolean isValid(String s) {
        List<Character> open = Arrays.asList('(', '[', '{');
        List<Character> close = Arrays.asList(')', ']', '}');
        Deque<Character> stack = new LinkedList<>();
        for (char c : s.toCharArray()) {
            if (open.contains(c)) {
                stack.add(close.get(open.indexOf(c)));
                continue;
            }
            if (close.contains(c)) {
                if (stack.isEmpty() || !stack.pollLast().equals(c)) {
                    return false;
                }
            }
        }
        if (!stack.isEmpty()) {
            return false;
        }
        return true;
    }

    public static void main(String[] args) {
        Solution sl = new Solution();
        assert sl.isValid("()") == false : "assert failed";
        System.out.println("hello world!");
        System.out.println(sl.isValid("()"));
    }
}

