/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/12/7 22:36.
<p>
description: 字符串转换整数 (atoi)
请你来实现一个 atoi 函数，使其能将字符串转换成整数。
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；
假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回 0。
说明：
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
示例 1:
* 输入: "42"
* 输出: 42
示例 2:
* 输入: "   -42"
* 输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例 3:
* 输入: "4193 with words"
* 输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例 4:
* 输入: "words and 987"
* 输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。因此无法执行有效的转换。
示例 5:
* 输入: "-91283472332"
* 输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。因此返回 INT_MIN (−2^31) 。
<P>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-to-integer-atoi
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(myAtoi(" +"))
}

const PlusSign string = "+"
const MinusSign string = "-"
const NumStr string = "0123456789"

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	item := ""
	for index := range str {
		item = string(str[index])
		if item == " " {
			// ----- 过滤字符串前段的空格
			continue
		}
		if !strings.Contains(NumStr, item) && item != MinusSign && item != PlusSign {
			// ----- 当第一个字符不是数字也不是正负号时：返回0
			return 0
		}
		if (item == PlusSign || item == MinusSign) && (index == len(str)-1 || !strings.Contains(NumStr, string(str[index+1]))) {
			// ----- 当第一个字符是正负号，但后面的字符不是数字时：返回0
			// ----- 当第一个字符为正负号，且该字符为最后一个字符时：返回0
			return 0
		}
		if item == PlusSign {
			item = ""
		}
		for index < len(str)-1 {
			// ----- 找到所有的数字字符
			if !strings.Contains(NumStr, string(str[index+1])) {
				break
			}
			index++
			item += string(str[index])
		}
		break
	}
	i, _ := strconv.Atoi(item)
	return processInt(i)
}

/**
处理超过32位有符号整数范围的情况
*/
func processInt(i int) int {
	if i > math.MaxInt32 {
		return math.MaxInt32
	}
	if i < math.MinInt32 {
		return math.MinInt32
	}
	return i
}
