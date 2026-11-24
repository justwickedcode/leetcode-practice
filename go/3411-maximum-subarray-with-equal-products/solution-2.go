// 3411. Maximum Subarray With Equal Products (Easy)
// https://leetcode.com/problems/maximum-subarray-with-equal-products/
// Runtime: 0 ms (beats 100.0%) · Memory: 5.1 MB (beats 28.6%)

func maxLength(nums []int) int {
	// Packed 4-bit prime factor masks for 10..1 (bits: 7, 5, 3, 2)
	const MASK = 0x5218341210
	// Extract 4-bit prime mask for x
	prime := func(x int) int { return MASK >> (uint(x-1)*4&63) & 0xF }

	// Any length-2 subarray is valid, so start best at 2
	i, mask, best := 0, 0, 2

	for j, v := range nums {
		m := prime(v)
		// Shrink window from left while v shares prime factors with current window
		for mask&m != 0 {
			mask &^= prime(nums[i])
			i++
		}
		// Add v's primes to window and update max length
		mask |= m
		best = max(best, j-i+1)
	}

	return best
}
