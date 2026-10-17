// 20. Valid Parentheses (Easy)
// https://leetcode.com/problems/valid-parentheses/
// Runtime: 0 ms (beats 100.0%) · Memory: 4.2 MB (beats 39.3%)

func isValid(s string) bool {
    if len(s)%2 == 1 {
        return false
    }

    stack := []rune{}
    pair := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 {
                return false
            }

            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if pair[top] != char {
                return false
            }
        }
    }

    return len(stack) == 0
}
