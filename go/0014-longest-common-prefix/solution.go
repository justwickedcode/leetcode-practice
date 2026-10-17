// 14. Longest Common Prefix (Easy)
// https://leetcode.com/problems/longest-common-prefix/
// Runtime: 0 ms (beats 100.0%) · Memory: 4.2 MB (beats 59.3%)

func longestCommonPrefix(strs []string) string {
    longestPrefix := ""
    prefixIndex := 0

    minLen := 202
    for _, s := range strs {
        minLen = min(minLen, len(s))
    }

    firstStr := strs[0]

    for prefixIndex < minLen {
        prefix := firstStr[prefixIndex]

        for _, str := range strs {
            if prefix != str[prefixIndex] {
                return longestPrefix
            } 
        }

        longestPrefix += string(prefix)

        prefixIndex++
    }

    return longestPrefix  
}
