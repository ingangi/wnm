package main

import (
	"fmt"
	"math"
)

//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
//Example 1:
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
//Example 2:
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5

// 思路：
// 题意： 2个排序的数组，找出两个数组合并之后的中位数，如果是奇数就取中间的那个数，如果是偶数就取中间两个数的平均
// 实现时肯定不能真去合并数组，否则效率太低，
// 因为是有序的数组，所以问题等同于求第K小的值，奇数：K=N/2+1  偶数K1=N/2  K2=N/2+1

// 从两个升序数组中找出第K小的数
func findKth(array1 []int, array2 []int, k int) (result int)  {

	if len(array1) == 0 {
		return array2[k-1]
	}
	if len(array2) == 0 {
		return array1[k-1]
	}

	if k == 1 {
		// 相当于求最小值
		return int(math.Min(float64(array1[0]), float64(array2[0])))
	}

	// 分别在两个数组中取K的一半的位置对应的数字做比较
	// 假设1对应的值小于2对应的值，那么说明目标值不会再1的前k/2个中
	// 将1的前k/2个元素排除后，范围缩小为求剩余元素的第 K - K/2 小的元素，递归
	index := k/2-1

	// 个数不够的情况
	if len(array1) < index+1 {
		// 目标值 必然不在arr1的前index+1个元素中
		return findKth(array1[index+1:], array2, k-len(array1))
	} else if len(array2) < index+1 {
		// 目标值 必然不在arr2的前index+1个元素中
		return findKth(array1, array2[index+1:], k-len(array2))
	}

	// 个数够
	if array1[index] < array2[index] {
		//
		return findKth(array1[index+1:], array2, k-index-1)
	}
	return findKth(array1, array2[index+1:], k-index-1)
}

// 解
func medianOfTwoSortedArrays(array1 []int, array2 []int) (f float32) {

	length := len(array1)+len(array2)
	if length == 0 {
		return 0
	}

	if length % 2 == 0 {
		f = float32(findKth(array1, array2, length/2) + findKth(array1, array2, length/2+1))/2.0
	} else {
		f = float32(findKth(array1, array2, length/2+1))
	}
	return 
}

func main()  {

	array1 := []int{1,2}
	array2 := []int{3,4,33,55,77,99}

	fmt.Println("median is", medianOfTwoSortedArrays(array1, array2))
}
