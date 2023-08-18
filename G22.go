题目描述：给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:

输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func lengthOfLongestSubstring(s string) int {
    n := len(s)
    charSet := make(map[byte]bool) 
    ans, left, right := 0, 0, 0  

    for right < n { 
        if charSet[s[right]] { 
            charSet[s[left]] = false
            left++                    
        } else { 
            charSet[s[right]] = true 
            right++                  
            ans = max(ans, right-left) 
        }
    }
    return ans 
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
思路：通过滑动窗口保证没有重复的字符串，不断更新最大长度
