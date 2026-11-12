// 3105. Longest Strictly Increasing or Strictly Decreasing Subarray (Easy)
// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/
// Runtime: 1 ms (beats 33.3%) · Memory: 2.2 MB (beats 16.7%)


impl Solution {
    pub fn longest_monotonic_subarray(nums: Vec<i32>) -> i32 {
        let mut increasing = 1;
        let mut decreasing = 1;
        let mut max_len = 1;

        for i in 1..nums.len() {
            if nums[i] > nums[i - 1] {
                increasing += 1;
                decreasing = 1;
            } else if nums[i] < nums[i - 1] {
                decreasing += 1;
                increasing = 1;
            } else {
                increasing = 1;
                decreasing = 1;
            }

            max_len = max_len.max(increasing).max(decreasing);
        }

        max_len
    }
}
