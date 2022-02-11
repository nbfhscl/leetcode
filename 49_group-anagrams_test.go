package test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestGroupAnagrams(t *testing.T) {
	myt := easytest.New(t).F(groupAnagrams)
	myt.Input([]string{"eat", "tea", "tan", "ate", "nat", "bat"}).Output([][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}})
	myt.Input([]string{""}).Output([][]string{{""}})
	myt.Input([]string{"a"}).Output([][]string{{"a"}})

	myt.Input([]string{"ac", "d"}).Output([][]string{{"ac"}, {"d"}})
}

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start

var primes []int

func getPrimes(num int) []int {
	if num == 0 {
		return []int{}
	}
	res := make([]int, num)
	res[0] = 2
	cnt := 1
	last := 1
outerfor:
	for cnt < num {
		last += 2
		if last%2 == 0 {
			continue
		}
		for i := 3; i <= int(math.Sqrt(float64(last))); i += 2 {
			if last%i == 0 {
				continue outerfor
			}
		}
		res[cnt] = last
		cnt++
	}
	return res
}

func hashCode(s string) int {
	res := 1
	for _, c := range s {
		res *= primes[c-'a']
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	primes = getPrimes(26)
	fmt.Println(primes)
	m := make(map[int][]string)
	for i := range strs {
		h := hashCode(strs[i])
		if _, ok := m[h]; ok {
			m[h] = append(m[h], strs[i])
		} else {
			m[h] = []string{strs[i]}
		}
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// @lc code=end

func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for i := range strs {
		h := make([]byte, 26)
		for _, v := range strs[i] {
			h[v-'a']++
		}
		strb := strings.Builder{}
		for i := range h {
			strb.WriteByte(h[i])
		}
		str := strb.String()
		if _, ok := m[str]; ok {
			m[str] = append(m[str], strs[i])
		} else {
			m[str] = []string{strs[i]}
		}
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
