// 1. Two Sum (Easy)
// https://leetcode.com/problems/two-sum/
// Runtime: 59 ms (beats 5.0%) · Memory: 5.3 MB (beats 98.5%)

func twoSum(nums []int, target int) []int {
    for i := range nums {
        for j := range nums {
            if i != j && nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil 
}
