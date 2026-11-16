// 2243. Calculate Digit Sum of a String (Easy)
// https://leetcode.com/problems/calculate-digit-sum-of-a-string/
// Runtime: 0 ms (beats 100.0%) · Memory: 4.7 MB (beats 28.0%)

import (
	"strconv"
)

func digitSum(s string, k int) string {
	for len(s) > k {
		s = groupToSumString(s, k)
	}

	return s
}

func groupToSumString(s string, k int) string {
	str := ""
	lenS := len(s)

	for i := 0; i+k <= lenS; i += k {
		str += stringToSum(s[i : i+k])
	}

	if lenS%k != 0 {
		str += stringToSum(s[(lenS - (lenS % k)):lenS])
	}

	return str
}

func stringToSum(s string) string {
	sum := 0
	for _, ch := range s {
        num, _ := strconv.Atoi(string(ch))
        sum += num
    }

	return strconv.Itoa(sum)
}
