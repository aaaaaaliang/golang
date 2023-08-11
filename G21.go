给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    
    left, right := 0, len(height)-1
    leftMax, rightMax := height[left], height[right]
    totalWater := 0
    
    for left <= right {
        if leftMax <= rightMax {
            if height[left] > leftMax {
                leftMax = height[left]
            } else {
                totalWater += leftMax - height[left]
            }
            left++
        } else {
            if height[right] > rightMax {
                rightMax = height[right]
            } else {
                totalWater += rightMax - height[right]
            }
            right--
        }
    }
    
    return totalWater
}
思路：双指针 left right遍历数组  更新leftmax和rightmax 累加到总雨水量
