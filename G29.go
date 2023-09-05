题目描述：
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
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
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串
func lengthOfLongestSubstring(s string) int {
    m:=map[byte]int{}
    start:=0
    maxLen:=0
    for i,ch:=range []byte(s){
       if lastIndex,ok:=m[ch];ok&&lastIndex>=start{
           start=lastIndex+1
       }
       m[ch]=i
       maxLen=max(maxLen,i-start+1)
    }
    return maxLen
}
func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}
思路：定义map,lastIndex,ok:=m[ch],如果重复更新起始位置,且lastIndex>=start，求出不重复的最长字符串
