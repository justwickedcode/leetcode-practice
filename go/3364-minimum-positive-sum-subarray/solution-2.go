// 3364. Minimum Positive Sum Subarray  (Easy)
// https://leetcode.com/problems/minimum-positive-sum-subarray/
// Runtime: 0 ms (beats 100.0%) · Memory: 5.9 MB (beats 7.4%)

func minimumSumSubarray(nums []int, l int, r int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)

	for i, num := range nums {
		prefixSum[i+1] = prefixSum[i] + num
	}

	minSum := math.MaxInt
	for length := l; length <= r; length++ {
		for start := 0; start+length <= n; start++ {
			sum := prefixSum[start+length] - prefixSum[start]
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
