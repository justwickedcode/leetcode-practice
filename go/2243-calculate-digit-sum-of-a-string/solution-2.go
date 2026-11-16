// 2243. Calculate Digit Sum of a String (Easy)
// https://leetcode.com/problems/calculate-digit-sum-of-a-string/
// Runtime: 0 ms (beats 100.0%) · Memory: 4.8 MB (beats 28.0%)

package main

import (
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {
	var sb strings.Builder

	for len(s) > k {
		sb.Reset()
		sb.Grow((len(s)/k + 1) * 3)
		groupToSumString(&sb, s, k)
		s = sb.String()
	}

	return s
}

func groupToSumString(sb *strings.Builder, s string, k int) {
	var buf [20]byte

	for i := 0; i < len(s); i += k {
		end := min(i+k, len(s))
		sum := stringToSum(s[i:end])
		sb.Write(strconv.AppendInt(buf[:0], int64(sum), 10))
	}
}

func stringToSum(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i] - '0')
	}
	return sum
}
