// 3364. Minimum Positive Sum Subarray  (Easy)
// https://leetcode.com/problems/minimum-positive-sum-subarray/
// Runtime: 6 ms (beats 0.0%) · Memory: 5.5 MB (beats 11.1%)

import (
	"fmt"
)

func minimumSumSubarray(nums []int, l int, r int) int {
	minSum := math.MaxInt
	length := len(nums)

	for rangeLen := l; rangeLen <= r; rangeLen++ {
		for i := 0; i <= length-rangeLen; i++ {
			sum := 0

			for j := i; j < i+rangeLen; j++ {
				sum += nums[j]
			}

			if sum > 0 && sum < minSum {
				minSum = sum
			}
		}
	}

	if minSum == math.MaxInt {
		return -1
	}
	return minSum
}
