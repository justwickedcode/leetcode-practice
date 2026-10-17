// 3551. Minimum Swaps to Sort by Digit Sum (Medium)
// https://leetcode.com/problems/minimum-swaps-to-sort-by-digit-sum/
// Runtime: 1021 ms (beats 7.1%) · Memory: 18.2 MB (beats 7.1%)

package main

import (
	"sort"
)

func minSwaps(nums []int) int {
	n := len(nums)
	
	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	
	digitSums := make(map[int]int)
	for _, num := range nums {
		digitSums[num] = getDigitSum(num)
	}
	
	sort.Slice(sortedNums, func(i, j int) bool {
		sumI := digitSums[sortedNums[i]]
		sumJ := digitSums[sortedNums[j]]
		if sumI != sumJ {
			return sumI < sumJ
		}
		return sortedNums[i] < sortedNums[j]
	})
	
	valueToTargetIndex := make(map[int]int)
	for i, num := range sortedNums {
		valueToTargetIndex[num] = i
	}
	
	visited := make([]bool, n)
	swaps := 0
	
	for i := 0; i < n; i++ {
		if visited[i] || valueToTargetIndex[nums[i]] == i {
			continue
		}
		
		cycleSize := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = valueToTargetIndex[nums[j]]
			cycleSize++
		}
		
		if cycleSize > 0 {
			swaps += (cycleSize - 1)
		}
	}
	
	return swaps
}

func getDigitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
