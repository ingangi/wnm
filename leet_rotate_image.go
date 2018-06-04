package main

import "fmt"

//You are given an n x n 2D matrix representing an image.
//
//Rotate the image by 90 degrees (clockwise).
//
//Note:
//
//You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
//
//Example 1:
//
//Given input matrix =
//[
//[1,2,3],
//[4,5,6],
//[7,8,9]
//],
//
//rotate the input matrix in-place such that it becomes:
//[
//[7,4,1],
//[8,5,2],
//[9,6,3]
//]

func printImg(img [][]int)  {
	for _,line := range img {
		fmt.Println(line)
	}
}

func rotateImage(img [][]int) {
	printImg(img)
	// 转置
	for i, line := range img {
		for j, _ := range line {
			if j > i {
				img[i][j], img[j][i] = img[j][i], img[i][j]
			}
		}
	}

	// 以中间列为中心翻转
	n := len(img)
	for j:=0; j<n/2; j++ {
		for i:=0; i < n; i++ {
			img[i][j], img[i][n-1-j] = img[i][n-1-j], img[i][j]
		}
	}

	printImg(img)
}

func main()  {

	img := [][]int{[]int{1,2,3,11},[]int{4,5,6,22},[]int{7,8,9,33},[]int{10,11,12,44}}
	rotateImage(img)
}