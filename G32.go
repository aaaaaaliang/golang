题目描述:
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
示例 1：
输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。
示例 2：
输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。
示例 3：
输入：s = " "
输出：true
解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
由于空字符串正着反着读都一样，所以是回文串。
代码：
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Compare(s string) bool {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return s == string(r)
}
func main() {
	s := "Avsjfncklafmfuha hoi,a.d a/ a.';f l;aks"
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	s = reg.ReplaceAllString(s, "")
	fmt.Println(s)
	fmt.Println(Compare(s))
}
思路：先去除字符串里面的非字符部分，然后全部转成小写，最后双指针遍历，将两边的字符串交换，然后判断是否一样
