题目描述：给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。
示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:
输入: nums = [0]
输出: [0]
提示:
func moveZeroes(nums []int)  {
    n := len(nums)
    j := 0

    for i :=0;i<n;i++{
        if nums[i] != 0{
          nums[i],nums[j] =nums[j],nums[i]
          j++
        }
    }
}
思路：双指针，遍历，遇到非0元素，与num[j]元素交换，j右移
