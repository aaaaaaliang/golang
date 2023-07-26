题目描述：
217. 存在重复元素
简单
1K
相关企业
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
示例 1：

输入：nums = [1,2,3,1]
输出：true
示例 2：

输入：nums = [1,2,3,4]
输出：false
示例 3：

输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true
func containsDuplicate(nums []int) bool {
    numMap :=make(map[int]bool)

    for _,num:=range nums{
        if numMap[num]{
            return true
        }
        numMap[num]=true
    }
    return false
}

思路：使用哈希表统计出现次数，出现过返回true说明至少出现两次
