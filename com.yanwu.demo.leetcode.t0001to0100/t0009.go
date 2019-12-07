/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/12/7 23:27.
<p>
description: 回文数
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
示例 1:
* 输入: 121
* 输出: true
示例 2:
* 输入: -121
* 输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:
* 输入: 10
* 输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:
* 你能不将整数转为字符串来解决这个问题吗？
<p>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(1012101))
}

func isPalindrome(x int) bool {
	// ----- 当数字只有一位时不做计算
	if x >= 0 && x <= 9 {
		return true
	}
	// ----- 当数字小于0时不做计算
	if x < 0 {
		return false
	}
	// ----- 将数字反转过来进行比较
	return x == reversal(x)
}

/**
反转数字：123 >> 321
*/
func reversal(x int) int {
	result := 0
	for x != 0 {
		// ----- 逐位取模，并将计算结果赋值给result
		result = result*10 + x%10
		x /= 10
	}
	return result
}
