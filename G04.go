349. 两个数组的交集
简单
792
相关企业
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的

func intersection(nums1 []int, nums2 []int) []int {
    numSet := make(map[int]bool)
    var result []int

    for _, num := range nums1 {
        numSet[num] = true
    }

    for _, num := range nums2 {
        if numSet[num] {
            result = append(result, num)
            numSet[num] = false
        }
    }

    return result
}
思路：用哈希表记录数组中的元素，遍历另一个数组判断元素是否在第一个数组中出现
