package action

import (
	"fmt"
	"strings"
)

type MedianFinder struct {
	Number int
	Pre    *MedianFinder
	Node   *MedianFinder
}
// 困难版 链表中计算中间数，如1 2中间数是1.5
/** initialize your data structure here. */
func Constructor() MedianFinder {
	node := new(MedianFinder)
	return MedianFinder{
		Node: node,
		Pre:  node,
	}
}
func (this *MedianFinder) AddNum(num int) {
	if this.Number < num {
		this.Number = num
		temp := this.Node
		this.Node = this
		this.Pre = temp
	}
	fmt.Println(this.Number, this.Node.Number)
	fmt.Println(this.Number, this.Pre.Number)
}

func (this *MedianFinder) FindMedian() float64 {
	var point float64
	if this.Node != nil {
		point = float64(this.Node.Number+this.Number) / 2
	} else if this.Node.Node != nil {
		point = float64(this.Node.Node.Number+this.Node.Number) / 2
	}
	return point
}

// 移动零
func moveZeroes(nums []int) {
	var num int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			num++
		} else if nums[i] != 0 {
			nums[i-num], nums[i] = nums[i], nums[i-num]
		}
	}
}

// 查找最大数的的值
func findDuplicate(nums []int) int {
	var point int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				point = nums[i]
			}
		}
	}
	return point
}

// 匹配单词
func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	strmap := make(map[rune]string)
	strnummap := make(map[string]int)
	if len(pattern) != len(strs) {
		return false
	}
	for i, s := range pattern {
		v, ok := strmap[s]
		if ok {
			if v != strs[i] {
				return false
			}
		} else {
			strmap[s] = strs[i]
			strnummap[strs[i]]++
			if strnummap[strs[i]] == 2 {
				return false
			}
		}
	}
	return true
}

/*
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:

给定 nums = [3,2,2,3], val = 3,

函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。

你不需要考虑数组中超出新长度后面的元素。
*/

func removeElement(nums []int, val int) int {
	var point int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == val {
				if nums[i] == nums[j] {
					point = j
				} else {
					point = i
				}
			} else if nums[j]==val {
				point = j
			}
		}
	}
	return point
}

// 求最长子串
func lengthOfLongestSubstring(s string) int {
	strMap := make(map[int32]int, 0)
	start := 0
	max := -1
	for k, v := range s {
		//如果值存在表示存在
		if value, ok := strMap[v]; ok {
			//开始指针
			if start <= value {
				start = value + 1
			}

		}
		strMap[v] = k
		if k-start > max {
			max = k - start
		}
	}
	return max + 1
}

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	return 2.2
}