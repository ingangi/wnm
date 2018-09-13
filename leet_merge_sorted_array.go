package main

import "fmt"

//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//
//Note:
//
//The number of elements initialized in nums1 and nums2 are m and n respectively.
//You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
//Example:
//
//Input:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//Output: [1,2,2,3,5,6]

func mergeSortedArray(nums1 []int, len1 int, nums2 []int, len2 int) {
	l := len1+len2
	i1, i2 := len1-1, len2-1
	for i := l-1; i >= 0 ; i-- {

		if i1 < 0 && i2 < 0 {
			break
		}
		if i1 < 0 {
			nums1[i] = nums2[i2]
			i2--
			continue
		} else if i2 < 0 {
			nums1[i] = nums1[i1]
			i1--
			continue
		}

		if nums1[i1] > nums2[i2] {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[i2]
			i2--
		}
	}
}

func main()  {
	nums1 := []int{11,12,13,0,0,0}
	nums2 := []int{4,5,6}
	mergeSortedArray(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}