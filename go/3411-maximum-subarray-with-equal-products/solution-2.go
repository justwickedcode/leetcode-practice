// 3411. Maximum Subarray With Equal Products (Easy)
// https://leetcode.com/problems/maximum-subarray-with-equal-products/
// Runtime: 0 ms (beats 100.0%) · Memory: 5 MB (beats 28.6%)

func maxLength(nums []int) int {
	const MASK = 0x5218341210
	prime := func(x int) int { return MASK >> (uint(x-1)*4&63) & 0xF }
	i, mask, best := 0, 0, 2
	for j, v := range nums {
		m := prime(v)
		for mask&m != 0 {
			mask &^= prime(nums[i])
			i++
		}
		mask |= m
		best = max(best, j-i+1)
	}
	return best
}
