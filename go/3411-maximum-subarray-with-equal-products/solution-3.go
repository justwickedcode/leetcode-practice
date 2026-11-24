// 3411. Maximum Subarray With Equal Products (Easy)
// https://leetcode.com/problems/maximum-subarray-with-equal-products/
// Runtime: 2 ms (beats 35.7%) · Memory: 4.9 MB (beats 50.0%)

func maxLength(nums []int) int {
	// Any length-2 subarray is valid, so start best at 2
	i, currLCM, best := 0, 1, 2

	for j, v := range nums {
		// Shrink window from left while v shares a factor (> 1) with currLCM
		for gcd(currLCM, v) > 1 {
			currLCM /= nums[i]
			i++
		}
		// Since v is coprime to currLCM, new LCM is just currLCM * v
		currLCM *= v
		best = max(best, j-i+1)
	}

	return best
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
