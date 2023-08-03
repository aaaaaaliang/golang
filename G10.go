给定两个字符串 s 和 t ，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。
示例 1：

输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。
示例 2：

输入：s = "", t = "y"
输出："y"
func findTheDifference(s string, t string) byte {
     var freq [26]int
     for _,ch:=range s{
         freq[ch-'a']++
     }

     for  _,ch:=range t{

         freq[ch-'a']--

         if freq[ch-'a']<0{

             return byte(ch)
         }
     } 

     return ' '

}
思路：遍历数组 如果统计每个字母出现次数  在遍历另外一个 出现就-- 小于0的就是新增的
