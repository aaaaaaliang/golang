给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
var phoneMap = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

func letterCombinations(digits string) []string {
    combinations := []string{}
    if len(digits) == 0 {
        return combinations
    }
    
    backtrack(&combinations, "", digits)
    return combinations
}

func backtrack(combinations *[]string, current string, digits string) {
    if len(current) == len(digits) {
        *combinations = append(*combinations, current)
        return
    }
    
    digit := string(digits[len(current)])
    letters := phoneMap[digit]
    for _, letter := range letters {
        backtrack(combinations, current+string(letter), digits)
    }
}
思路：调用函数使用递归生成字母组合  根据数字到字母的映射 ，不同的字母逐渐添加进当前组合
