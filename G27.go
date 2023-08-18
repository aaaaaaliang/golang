题目描述：给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库内置的 sort 函数的情况下解决这个问题。
示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]
func sortColors(nums []int) {
    red, blue := 0, len(nums)-1
    curr := 0

    for curr <= blue {
        if nums[curr] == 0 {
            nums[curr], nums[red] = nums[red], nums[curr]
            red++
            curr++
        } else if nums[curr] == 2 {
            nums[curr], nums[blue] = nums[blue], nums[curr]
            blue--
        } else {
            curr++
        }
    }
}
思路：定义三个指针一个指向开头，一个指向末尾，一个用来遍历数组，最后将0移到红色区域，2移到蓝色区域
