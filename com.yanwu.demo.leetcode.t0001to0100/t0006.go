/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/12/6 15:46.
<p>
description: Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
0   2   4   6
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
请你实现这个将字符串进行指定行数变换的函数：
string convert(string s, int numRows);
示例 1:
* 输入: s = "LEETCODEISHIRING", numRows = 3
* 输出: "LCIRETOESIIGEDHN"
示例 2:
* 输入: s = "LEETCODEISHIRING", numRows = 4
* 输出: "LDREOEIIECIHNTSG"
解释:
0     3     6
L     D     R
E   O E   I I
E C   I H   N
T     S     G
<P>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	str, numRows := "PAYPALISHIRING", 4
	fmt.Println(convert(str, numRows))
}

func convert(s string, numRows int) string {
	if len(s) <= 1 || len(s) <= numRows {
		return s
	}
	// ----- result[返回值]；tempArr[Z字每一行的字符串]；curRow[]：goingDown[]
	result, tempArr, curRow, goingDown := "", make([]string, len(s)), 0, false
	for index := range s {
		tempArr[curRow] += string(s[index])
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			// ----- 判断往上走
			curRow++
		} else {
			// -----
			curRow--
		}
	}
	for index := range tempArr {
		result += tempArr[index]
	}
	return result
}
