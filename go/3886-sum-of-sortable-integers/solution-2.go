// 3886. Sum of Sortable Integers (Hard)
// https://leetcode.com/problems/sum-of-sortable-integers/
// Runtime: 31 ms (beats 53.3%) · Memory: 9.9 MB (beats 93.3%)

func sortableIntegers(nums []int) int {
	n := len(nums)

	check := func(k int) bool {
		prev := 0

		for i := 0; i < n; i += k {
			cnt := 0
			mini, maxi := nums[i], nums[i]

			for j := i; j < i+k-1; j++ {
				if nums[j] > nums[j+1] {
					cnt++
				}

				if nums[j+1] < mini {
					mini = nums[j+1]
				}
				if nums[j+1] > maxi {
					maxi = nums[j+1]
				}
			}

			if nums[i+k-1] > nums[i] {
				cnt++
			}

			if cnt > 1 || prev > mini {
				return false
			}

			prev = maxi
		}

		return true
	}

	ans := 0

	for k := 1; k*k <= n; k++ {
		if n%k == 0 {
			if check(k) {
				ans += k
			}

			other := n / k
			if other != k && check(other) {
				ans += other
			}
		}
	}

	return ans
}
