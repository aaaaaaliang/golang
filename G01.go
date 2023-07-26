题目描述：
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2

func majorityElement(nums []int) int {
    replace := nums[0]
    count := 1

    for i := 1; i < len(nums); i++ {
        if replace == nums[i] {
            count++
        } else {
            count--
        }

        if count == 0 {
            replace = nums[i]
            count = 1
        }
    }

    return replace
}

思路：数组中的第一个 在定义一个计数器count，遇到相同的加一 遇到不同的减一
