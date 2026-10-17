// 9. Palindrome Number (Easy)
// https://leetcode.com/problems/palindrome-number/
// Runtime: 10 ms (beats 11.0%) · Memory: 6 MB (beats 68.7%)

func isPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }

    revertedNumber := 0

    for x > revertedNumber {
        revertedNumber = (revertedNumber * 10) + (x % 10)
        x /= 10
    }
    
    return x == revertedNumber || x == revertedNumber/10
}
