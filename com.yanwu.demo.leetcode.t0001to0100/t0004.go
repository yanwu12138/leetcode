/**
@author <a herf="mailto:yanwu0527@163.com">XuBaofeng</a>
@date 2019/12/3 20:18.
<p>
description: 寻找两个有序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。
示例 1:
* nums1 = [1, 3]
* nums2 = [2]
* 则中位数是 2.0
示例 2:
* nums1 = [1, 2]
* nums2 = [3, 4]
* 则中位数是 (2 + 3)/2 = 2.5
<p>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) > 0 {
		return 0.0
	} else if len(nums1) == 0 && len(nums2) > 0 {
		return getMedianSorted(nums2)
	} else if len(nums1) > 0 && len(nums2) == 0 {
		return getMedianSorted(nums1)
	} else {
		var temp = make([]int, len(nums1)+len(nums2))
		if len(nums1) > len(nums2) {
			for i := 0; i < len(nums1); i++ {
				for j := 0; j < len(nums2); j++ {
					if nums1[i] < nums2[j] {
						temp[i+j] = nums1[i]
					} else {
						continue
					}
				}
			}
		} else {
			for i := 0; i < len(nums2); i++ {
				for j := 0; j < len(nums1); j++ {
					if nums2[i] < nums1[j] {
						temp[i+j] = nums2[i]
					} else {
						continue
					}
				}
			}
		}
		return getMedianSorted([]int{})
	}
}

func getMedianSorted(nums []int) float64 {
	median := len(nums) / 2
	fmt.Println(median)
	return 0.0
}
