// 3963. Create Grid With Exactly One Path (Easy)
// https://leetcode.com/problems/create-grid-with-exactly-one-path/
// Runtime: 1 ms (beats 25.0%) · Memory: 6.4 MB (beats 33.3%)

func createGrid(m int, n int) []string {
    grid := make([][]byte, m)

    // Fill the grid with obstacles.
    for i := 0; i < m; i++ {
        grid[i] = make([]byte, n)

        for j := 0; j < n; j++ {
            grid[i][j] = '#'
        }
    }

    // Make the first row free.
    for j := 0; j < n; j++ {
        grid[0][j] = '.'
    }

    // Make the last column free.
    for i := 0; i < m; i++ {
        grid[i][n-1] = '.'
    }

    // Convert each row into a string.
    result := make([]string, m)

    for i := 0; i < m; i++ {
        result[i] = string(grid[i])
    }

    return result
}
