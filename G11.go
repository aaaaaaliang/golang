给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。
示例 1：

输入：nums = [1,1,0,1,1,1]
输出：3
解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
示例 2:

输入：nums = [1,0,1,1,0,1]
输出：2
func findMaxConsecutiveOnes(nums []int) int {
    count:=0
    maxCount:=0
    for _,num:=range nums{
        if num==1{
            count++
        }else{
            maxCount=max(maxCount,count)
            count=0
        }
    }
    maxCount = max(maxCount, count)
    return maxCount

}
func max(a,b int) int {
   if a>b {
       return a
   }
   return b
}
思路：定义两个计数器，如果是1，count++，否则更新最大计数器，并把count清零
