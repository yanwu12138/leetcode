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

import "fmt"

func main() {
	str := "aba"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	// ----- 起始位置和结束位置
	var ind = []int{0, 0}
	for i := 0; i < len(s); i++ {
		// ----- 把回文看成中间的部分全是同一字符，左右部分相对称
		// ----- 找到下一个与当前字符不同的字符
		i = expandAroundCenter(s, i, ind)
	}
	return s[ind[0] : ind[1]+1]
}

func expandAroundCenter(str string, L int, ind []int) int {
	// ----- 找到中间相同字符部分
	R := L
	for R < len(str)-1 && str[R+1] == str[L] {
		R++
	}
	// ----- 定位中间同字符部分的最后一个字符
	ans := R
	// ----- 从中间向左右扩散
	for L > 0 && R < len(str)-1 && str[L-1] == str[R+1] {
		L--
		R++
	}
	// ----- 记录开始与结束位置
	if R-L > ind[1]-ind[0] {
		ind[0] = L
		ind[1] = R
	}
	// ----- 返回进行下一轮的查找
	return ans
}
