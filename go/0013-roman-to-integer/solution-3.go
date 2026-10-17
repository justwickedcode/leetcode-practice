// 13. Roman to Integer (Easy)
// https://leetcode.com/problems/roman-to-integer/
// Runtime: 4 ms (beats 31.7%) · Memory: 4.6 MB (beats 86.8%)

func romanToInt(s string) int {
    roman := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

	result := 0

	j := 1

	for i := len(s) - 1; i >= 1; i-- {
		char1 := rune(s[i]) 
		char2 := rune(s[i-1])

		if roman[char1] > roman[char2] {
			result = result + roman[char1] - roman[char2]
			i--
		} else {
			result += roman[char1]
		}

		j = i
	}
    
	if j != 0 {
		result += roman[rune(s[0])]
	}

	return result
}
