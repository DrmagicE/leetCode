package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	var start, end, maxlen int
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i) //以 i 为中心两边扩展
		len2 := expandAroundCenter(s, i, i+1) //以 i和i+1之间为中心，两边扩展
		if len1 > len2 {
			maxlen = len1
		} else {
			maxlen = len2
		}

		if maxlen > end - start {
			start = i - (maxlen - 1) / 2
			end = i + maxlen / 2
		}
	}
	return s[start:end+1]
}

//以
func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func main() {
	fmt.Printf(longestPalindrome("abbba"))
}
