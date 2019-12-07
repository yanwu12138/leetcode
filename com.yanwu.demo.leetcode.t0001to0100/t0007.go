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
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
<P>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {

	return 0
}

/**
class Solution {
    public int reverse(int x) {
        int a = 0;
        try {
            String reverse = "";
            if (x >= 0) {
                String s = String.valueOf(x);
                int length = s.length();
                for (int i = 0; i < length; i++) {
                    reverse = s.charAt(i) + reverse;
                }
            } else {
                String s = String.valueOf(x).split("-")[1];
                int length = s.length();
                for (int i = 0; i < length; i++) {
                    reverse = s.charAt(i) + reverse;
                }
                reverse = "-" + reverse;
            }
            int result = Integer.parseInt(reverse);
            if (result > Integer.MAX_VALUE || result < Integer.MIN_VALUE) {
                a = 0;
            } else {
                a = result;
            }
        } catch (Exception e) {

        }
        return a;
    }
}
*/
