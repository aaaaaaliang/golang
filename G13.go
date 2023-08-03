 

示例 1：

输入：words = ["Hello","Alaska","Dad","Peace"]
输出：["Alaska","Dad"]
示例 2：

输入：words = ["omk"]
输出：[]
示例 3：

输入：words = ["adsdf","sfd"]
输出：["adsdf","sfd"]

func findWords(words []string) []string {
    rowMap := map[rune]int{
        'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
        'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
        'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
    }

    var result []string
    for _, word := range words {
        sameRow := true
        row := rowMap[unicode.ToLower(rune(word[0]))]
        for i := 1; i < len(word); i++ {
            if rowMap[unicode.ToLower(rune(word[i]))] != row {
                sameRow = false
                break
            }
        }
        if sameRow {
            result = append(result, word)
        }
    }

    return result
}
思路：构建哈希表，将每个字母与对应的行数对应，遍历字符判断是否在一行
