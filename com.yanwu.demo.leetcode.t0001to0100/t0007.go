/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/12/7 17:42.
<p>
description: 整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例 1:
* 输入: 123
* 输出: 321
示例 2:
* 输入: -123
* 输出: -321
示例 3:
* 输入: 120
* 输出: 21
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31, 2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
<p>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1))
}

func reverse(x int) int {
	result := 0
	if x >= -9 && x <= 9 {
		// ----- 个位数直接返回
		return x
	}
	for x != 0 {
		// ----- 逐位取模，并将计算结果赋值给result
		pop := x % 10
		result = result*10 + pop
		x /= 10
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		// ----- 如果数字结果超过[−2^31, 2^31 − 1]返回0
		result = 0
	}
	return result
}
