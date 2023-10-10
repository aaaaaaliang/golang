题目描述：
给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x

func isPowerOfFour(n int) bool {
    if n <= 0 {
		return false
	}
    
	if n&(n-1) != 0 {
		return false
	}
	return n&0x55555555 == n
} 

思路：排除小于等于0的情况，然后&只有同时为1时才取1，求出是否是二进制，最后检查n是否为4的幂次方
