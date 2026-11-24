// 3411. Maximum Subarray With Equal Products (Easy)
// https://leetcode.com/problems/maximum-subarray-with-equal-products/
// Runtime: 3 ms (beats 21.4%) · Memory: 5 MB (beats 28.6%)

import "math"

func maxLength(nums []int) int {
	length := 0
	arrLen := len(nums)

	for i := 0; i < arrLen; i++ {
		prod := 1
		currGCD := 0
		currLCM := 1

		for j := i; j < arrLen; j++ {
			if nums[j] != 0 && prod > math.MaxInt/nums[j] {
				break
			}
			prod *= nums[j]

			if j == i {
				currGCD = nums[j]
				currLCM = nums[j]
			} else {
				currGCD = gcd(currGCD, nums[j])
				currLCM = lcm(currLCM, nums[j])
			}

			if prod == currGCD*currLCM {
				if j-i+1 > length {
					length = j - i + 1
				}
			}
		}
	}

	return length
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a / gcd(a, b)) * b
}
