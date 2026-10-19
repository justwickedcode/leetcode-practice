// 2089. Find Target Indices After Sorting Array (Easy)
// https://leetcode.com/problems/find-target-indices-after-sorting-array/
// Runtime: 0 ms (beats 100.0%) · Memory: 4.4 MB (beats 30.0%)

func targetIndices(nums []int, target int) []int {
	arr := []int{}
	index := 0
	count := 0

    // count the number of time target is present
    // and how many numbers are smaller than it
	for _, val := range nums {
		if val == target {
			count++
		} else if val < target {
			index++
		}
	}

	for offset := range count {
		arr = append(arr, index + offset)
	}

	return arr
}
