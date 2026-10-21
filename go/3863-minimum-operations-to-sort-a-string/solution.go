// 3863. Minimum Operations to Sort a String (Medium)
// https://leetcode.com/problems/minimum-operations-to-sort-a-string/
// Runtime: 2 ms (beats 100.0%) · Memory: 9.5 MB (beats 0.0%)

func minOperations(s string) int {
    n := len(s)
    sorted := true

    // verifiy if the string is already sorted
    for i := 1; i < n; i++ {
        if s[i] < s[i - 1]{
            sorted = false
            break
        }
    }
    // if sorted return 0
    if sorted {
        return 0
    }

    // if not sorted and len = 2 => can't be sorted
    if n == 2 {
        return -1
    }

    // compute mini and maxi
    mini, maxi := s[0], s[0]

    for i := 1; i < n; i++ {
        if s[i] < mini {
            mini = s[i]
        }

        if s[i] > maxi {
            maxi = s[i]
        }
    }

    // if the first is mini or last is maxi, return 1
    if s[0] == mini || s[n-1] == maxi {
        return 1
    }

    // if mini or maxi is in the middle of the string, return 2
    for i := 1; i < n-1; i++ {
        if s[i] == mini || s[i] == maxi {
            return 2
        }
    }
    // return 3 in any other case
    return 3
}
