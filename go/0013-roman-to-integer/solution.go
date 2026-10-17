// 13. Roman to Integer (Easy)
// https://leetcode.com/problems/roman-to-integer/
// Runtime: 1 ms (beats 59.7%) · Memory: 4.7 MB (beats 39.3%)

func romanToInt(s string) int {
    roman := map[string]int{
    "I":  1,
    "V":  5,
    "X":  10,
    "L":  50,
    "C":  100,
    "D":  500,
    "M":  1000,
    }

    num := roman[string(s[0])]

    for i := 1; i < len(s); i++ {
        if (s[i-1] == 'I' && (s[i] == 'V' || s[i] == 'X')) ||
            (s[i-1] == 'X' && (s[i] == 'L' || s[i] == 'C')) ||
            (s[i-1] == 'C' && (s[i] == 'D' || s[i] == 'M')) {
            num -= 2 * roman[string(s[i-1])]
        }

        num += roman[string(s[i])]
    }

    return num
}
