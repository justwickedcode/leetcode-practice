// 14. Longest Common Prefix (Easy)
// https://leetcode.com/problems/longest-common-prefix/
// Runtime: 1 ms (beats 9.9%) · Memory: 4.4 MB (beats 23.6%)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    minLen := len(strs[0])
    for _, s := range strs[1:] {
        minLen = min(minLen, len(s))
    }

    prefix := strs[0][:minLen]

    for i := 0; i < minLen; i++ {
        for _, s := range strs[1:] {
            if strs[0][i] != s[i] {
                return prefix[:i]
            }
        }
    }

    return prefix
}
