/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019-12-03 9:31.
<p>
description: 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
* 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
* 输出：7 -> 0 -> 8
* 原因：342 + 465 = 807
<p>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func listNode(arr []int) *ListNode {
	var result *ListNode
	for i := 0; i < len(arr); i++ {
		result = getNode(arr[i], result)
	}
	return result
}

func getNode(Val int, Next *ListNode) *ListNode {
	return &ListNode{Val: Val, Next: Next}
}

func main() {
	l1 := listNode([]int{5})
	l2 := listNode([]int{5})
	numbers := addTwoNumbers(l1, l2)
	fmt.Printf("addTwoNumbers result: %v \n", numbers)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	return getListNode(l1, l2, 0)
}

/**
 * l1 链表1
 * l2 链表2
 * ind 是否进一位
 */
func getListNode(l1 *ListNode, l2 *ListNode, ind int) *ListNode {
	if l1 == nil && l2 == nil {
		// ----- 如果两个链表都为nil时，判断ind是否为1
		if ind == 1 {
			// ----- 为1说明前一次两个val的和大于10，需要进一
			return &ListNode{Val: ind, Next: nil}
		} else {
			return nil
		}
	}
	// ----- 处理l1和l2的val、next
	Val1, Val2 := 0, 0
	var Next1 *ListNode = nil
	var Next2 *ListNode = nil
	if l1 != nil {
		Val1 = l1.Val
		Next1 = l1.Next
	}
	if l2 != nil {
		Val2 = l2.Val
		Next2 = l2.Next
	}
	// ----- 将两个val相加，并将结果与10比较，判断结果是否需要进一
	val := Val1 + Val2 + ind
	if val < 10 {
		ind = 0
	} else {
		ind = 1
		val = val % 10
	}
	// ----- 递归获取next节点
	return &ListNode{Val: val, Next: getListNode(Next1, Next2, ind)}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
