package main

import "fmt"


//滑动窗口法,自己实现版
//时间和空间复杂度同解法二
func lengthOfLongestSubstring(s string) int {
	var maxlength int
	var currentLength int
	for s != "" { //左窗口
		t := make(map[uint8]struct{})
		length := len(s)
		for i := 0; i < length; i++ { //右窗口
			if _, ok := t[s[i]]; ok {
				currentLength = 0
				break;
			}
			t[s[i]] = struct{}{}
			currentLength++
			if maxlength <= currentLength {
				maxlength = currentLength
			}
		}
		s = s[1:]
		if maxlength >= len(s) {
			break;
		}
	}
	return maxlength
}

//解答方法二：
//时间复杂度：O(2n) = O(n)O(2n)=O(n)，在最糟糕的情况下，每个字符将被 ii 和 jj 访问两次。
//空间复杂度：O(min(m, n))O(min(m,n))，
func lengthOfLongestSubstringA(s string) int {
	n := len(s)
	t := make(map[uint8]struct{})
	var ans,i,j int
	for (i < n && j < n) {
		// try to extend the range [i, j]
		if _, ok := t[s[j]]; !ok {
			t[s[j]] = struct {}{}
			j++
			if j - i >= ans {
				ans = j -i
			}
		} else {
			delete(t,s[i])
			i++
		}

	}
	return ans;
}

//解法三：
//优化过的滑动窗口
func lengthOfLongestSubstringB(s string) int {
	n := len(s)
	t := make(map[uint8]int)
	var ans,i,j int
	for ; j < n; j++ {
		// try to extend the range [i, j]
		if index, ok := t[s[j]]; ok {
			if index > i {
				i = index
			}
		}
		if j - i + 1 > ans {
			ans = j - i + 1
		}
		t[s[j]] = j + 1

	}
	return ans;
}


func main() {
	fmt.Println(lengthOfLongestSubstringA("abcabcbb"))

}