// 13. Roman to Integer (Easy)
// https://leetcode.com/problems/roman-to-integer/
// Runtime: 1 ms (beats 59.7%) · Memory: 4.7 MB (beats 39.3%)

func romanToInt(s string) int {
    roman := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    num := 0

    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
            num -= roman[s[i]]
        } else {
            num += roman[s[i]]
        }
    }

    return num
}
