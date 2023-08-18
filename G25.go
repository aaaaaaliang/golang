题目描述：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if opening, exists := mapping[char]; exists {
			var topElement rune
			if len(stack) > 0 {
				topElement, stack = stack[len(stack)-1], stack[:len(stack)-1]
			} else {
				topElement = '#'
			}

			if topElement != opening {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
思路：创建栈，遇到左括号的入栈，判断右括号是否匹配，不匹配就返回false
