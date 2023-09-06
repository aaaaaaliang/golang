func isValidSudoku(board [][]byte) bool {
    rows := [9][9]int{}
    cols := [9][9]int{}
    boxes := [3][3][9]int{}
    
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == '.' {
                continue
            }
            
            num := board[r][c] - '1'
            if rows[r][num] > 0 {
                return false
            }
            rows[r][num]++
            
            if cols[c][num] > 0 {
                return false
            }
            cols[c][num]++
            
            boxR, boxC := r/3, c/3
            if boxes[boxR][boxC][num] > 0 {
                return false
            }
            boxes[boxR][boxC][num]++
        }
    }
    
    return true
}
