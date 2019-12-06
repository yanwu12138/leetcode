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

import (
	"fmt"
)

func main() {
	nums1 := []int{1}
	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	if len(nums1) == 0 && len(nums2) > 0 {
		return getMedianSorted(nums2)
	} else if len(nums1) > 0 && len(nums2) == 0 {
		return getMedianSorted(nums1)
	} else {
		// ----- 将两个数组合并起来，然后找到中位数
		return getMedianSorted(short(nums1, nums2))
	}
}

/**
使用归并排序法将两个数组组合成一个数组
*/
func short(nums1, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	nums := make([]int, len1+len2)
	i, j, k := 0, 0, 0
	for {
		if i >= len1 || j >= len2 {
			break
		}
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}
	for {
		if i >= len1 {
			break
		}
		nums[k] = nums1[i]
		k++
		i++
	}
	for {
		if j >= len2 {
			break
		}
		nums[k] = nums2[j]
		k++
		j++
	}
	return nums
}

func getMedianSorted(nums []int) float64 {
	if len(nums) == 1 {
		// ----- 当长度为1时直接返回
		return float64(nums[0])
	}
	median := len(nums) / 2
	if len(nums)&1 == 1 {
		// ----- 奇数
		return float64(nums[median])
	} else {
		// ----- 偶数
		return (float64(nums[median-1]) + float64(nums[median])) / 2

	}
}
