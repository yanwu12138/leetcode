/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/12/3 19:28.
<p>
description: 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
* 输入: "abcabcbb"
* 输出: 3
* 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
* 输入: "bbbbb"
* 输出: 1
* 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
* 输入: "pwwkew"
* 输出: 3
* 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
* 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
<p>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLength := 0
	for i := 0; i < len(s); i++ {
		temp := 1
		item := string(s[i])
		for j := i + 1; j < len(s); j++ {
			if strings.Contains(item, string(s[j])) {
				break
			}
			item = item + string(s[j])
			temp++
		}
		if temp > maxLength {
			maxLength = temp
		}
	}
	return maxLength
}
