题目描述：
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
示例 2：

输入：ransomNote = "aa", magazine = "ab"
输出：false
示例 3：

输入：ransomNote = "aa", magazine = "aab"
输出：true

func canConstruct(ransomNote string, magazine string) bool {
    charCount :=make(map[rune]int)

    for _,ch := range magazine{
        charCount[ch]++
    }

    for _,ch := range ransomNote{
        charCount[ch]--
        if charCount[ch] <0 {
        return false
       }
    }

     return true

}
思路：统计每个字符出现的次数，在遍历另一个字符串，如果有相同的就减减
