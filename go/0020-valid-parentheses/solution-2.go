// 20. Valid Parentheses (Easy)
// https://leetcode.com/problems/valid-parentheses/
// Runtime: 1 ms (beats 6.9%) · Memory: 4.3 MB (beats 39.3%)

func isValid(s string) bool {
    if len(s) % 2 != 0 {
        return false
    }
    stack := make([]rune, 0)
    pairs := make(map[rune]rune, 0)
    pairs['('] = ')'
    pairs['{'] = '}'
    pairs['['] = ']'
    for _, c := range s {
        if _, ok := pairs[c]; ok {
            stack = append(stack, c)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if pairs[top] != c {
                return false
            }
        }
    }
    return len(stack) == 0
}
