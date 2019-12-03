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
	"strings"
)

func main() {
	str := "babad"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	subStr := ""
	size, index := 0, 0
	for i := 0; i < len(s); i++ {
		temp := s[index:i]
		surplus := reverseString(s[i:])
		fmt.Println("temp: ", temp, " surplus: ", surplus)
		if !strings.Contains(surplus, temp) {
			index = i
			continue
		}
		if len(temp) > size {
			subStr = temp
		}
	}
	temp := subStr
	if strings.Contains(s[len(subStr)+1:], s[strings.Index(s, subStr):len(subStr)+1]) {
		for i := len(temp); i > 0; i-- {
			subStr = subStr + string(temp[i-1])
		}
	} else {
		for i := len(temp) - 1; i > 0; i-- {
			subStr = subStr + string(temp[i-1])
		}
	}
	return subStr
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
