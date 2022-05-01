import java.util.Deque;
import java.util.LinkedList;
import java.util.Map;

class Solution {
    public boolean isValid(String s) {
        Map<Character, Character> opMap = Map.of('(', ')', '[', ']', '{', '}');
        Deque<Character> stack = new LinkedList<>();
        for (char c : s.toCharArray()) {
            if (opMap.containsKey(c)) {
                stack.add(opMap.get(c));
                continue;
            }
            if (opMap.values().contains(c)) {
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
        assert sl.isValid("()") == true : "assert failed";
        System.out.println("hello world!");
        System.out.println(sl.isValid("()"));
    }
}

