// 26. Remove Duplicates from Sorted Array (Easy)
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Runtime: 0 ms (beats 100.0%) · Memory: 2.5 MB (beats 13.7%)


impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.is_empty() {
            return 0;
        }

        let mut i = 0;

        for j in 1..nums.len() {
            if nums[j] != nums[i] {
                i += 1;
                nums[i] = nums[j];
            }
        }

        (i + 1) as i32
    }
}
