题目描述：给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
    return str == reverse(str)  
}

func reverse(s string) string {
    runes := []rune(s)
    n := len(runes)
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    return string(runes)
}
思路：将整数转成字符串，然后翻转字符串，判断是否一样
