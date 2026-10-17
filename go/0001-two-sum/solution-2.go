// 1. Two Sum (Easy)
// https://leetcode.com/problems/two-sum/
// Runtime: 1 ms (beats 51.9%) · Memory: 5.8 MB (beats 70.3%)

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int, len(nums)) // value -> index

    for i, n := range nums {
        if j, ok := seen[target-n]; ok {
            return []int{j, i}
        }
        seen[n] = i
    }
    return nil
}
