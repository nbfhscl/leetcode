package test

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tcs := []struct {
		name   string
		input  string
		output bool
	}{
		{
			name:   "empty input",
			input:  "",
			output: false,
		},
		{
			name:   "single left bracket",
			input:  "{",
			output: false,
		},
		{
			name:   "single right bracket",
			input:  "}",
			output: false,
		},
		{
			name:   "}{",
			input:  "}{",
			output: false,
		},
		{
			name:   "{}",
			input:  "{}",
			output: true,
		},
		{
			name:   "(){}[]",
			input:  "(){}[]",
			output: true,
		},
		{
			name:   "(]",
			input:  "(]",
			output: false,
		},
		{
			name:   "{[()]}",
			input:  "{[()]}",
			output: true,
		},
		{
			name:   "{[](())}()[]({}){}",
			input:  "{[](())}()[]({}){}",
			output: true,
		},
		{
			name:   "{[]((]))}()[]({}){}",
			input:  "{[]((]))}()[]({}){}",
			output: false,
		},
		{
			name:   "{[](())}()[]({}){}[",
			input:  "{[](())}()[]({}){}[",
			output: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := isValid(tc.input)
			if tc.output != output {
				t.Errorf("input: %v, output: %v, expected: %v", tc.input, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	rl := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	filo := make([]byte, 0, len(s)/2)
	for _, b := range []byte(s) {
		if b == '(' || b == '[' || b == '{' {
			if len(filo) >= len(s)/2 {
				return false
			}
			filo = append(filo, b)
			continue
		}
		if b == ')' || b == ']' || b == '}' {
			if len(filo) == 0 || rl[b] != filo[len(filo)-1] {
				return false
			}
			filo = filo[:len(filo)-1]
			continue
		}
		return false
	}
	return len(filo) == 0
}

// @lc code=end
