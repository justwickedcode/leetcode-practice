// 3963. Create Grid With Exactly One Path (Easy)
// https://leetcode.com/problems/create-grid-with-exactly-one-path/
// Runtime: 2 ms (beats 8.3%) · Memory: 6.1 MB (beats 91.7%)

func createGrid(m int, n int) []string {
    res := make([]string, m)
    for i := range m {
        row := make([]byte, n)
        for j := range n {
            if i == 0 || j == n-1 {
                row[j] = '.'
            } else {
                row[j] = '#'
            }
        }
        res[i] = string(row)
    }
    return res
}
