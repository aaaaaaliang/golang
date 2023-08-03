给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。
你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。
func findDuplicates(nums []int) []int {
     var result []int
     for _,num :=range nums{
         index :=abs(num) -1
         if nums[index] <0{
             result=append(result,abs(num))
         }else{
             nums[index]=-nums[index]
         }
     }
    
     return result
}

func abs(num int) int {
    if num<0  {
        return -num
    }
    return num
}
思路：出现两次取反
