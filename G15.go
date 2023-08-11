给你一个字符串 s，找到 s 中最长的回文子串。
如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：
输入：s = "cbbd"
输出："bb"
func longestPalindrome(s string) string {
    start := 0
    end := 0    
    for i := 0; i < len(s); i++ {
        len1 := expand(s, i, i)
        len2 := expand(s, i, i+1)
        length := max(len1, len2)
        
        if length > end - start + 1 {
            start = i - (length - 1) / 2
            end = i + length / 2
        }
    }  
    return s[start:end+1]
}
func expand(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
思路：从中间往外扩展，不断更新长度，返回最长长度
