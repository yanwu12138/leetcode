/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/12/3 22:51.
<p>
description: 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
* 输入: "babad"
* 输出: "bab"
* 注意: "aba" 也是一个有效答案。
示例 2：
* 输入: "cbbd"
* 输出: "bb"
<p>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
)

func main() {
	str := "ccccc"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	subStr := ""
	for i := 0; i < len(s); i++ {
		temp := expandAroundCenter(s, string(s[i]), i)
		if len(temp) > len(subStr) {
			subStr = temp
		}

	}
	return subStr
}

func expandAroundCenter(str, item string, index int) string {
	L, R := index, index
	result := item
	if L > 0 && item == string(str[L-1]) {
		L--
		result = string(str[L]) + result
	}
	if R < len(str)-1 && item == string(str[R+1]) {
		R++
		result = result + string(str[R])
	}
	for {
		if L > 0 && str[L] == str[L-1] && str[L-1] == str[R+1] {
			L--
			result = string(str[L]) + result
		}
		if R < len(str)-1 && str[R] == str[R+1] && str[L-1] == str[R+1] {
			R++
			result = result + string(str[R])
		}
		if L-1 < 0 || R+1 >= len(str) || str[L-1] != str[R+1] {
			break
		}
		L--
		R++
		result = string(str[L]) + result + string(str[R])
	}
	return result
}
