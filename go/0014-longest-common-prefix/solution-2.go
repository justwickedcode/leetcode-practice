// 14. Longest Common Prefix (Easy)
// https://leetcode.com/problems/longest-common-prefix/
// Runtime: 0 ms (beats 100.0%) · Memory: 4.2 MB (beats 59.3%)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    for i := 0; i < len(strs[0]); i++ {
        char := strs[0][i]

        for _, s := range strs[1:] {
            if i >= len(s) || s[i] != char {
                return strs[0][:i]
            }
        }
    }

    return strs[0]
}
