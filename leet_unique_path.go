package main

import "fmt"

//A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
//
//The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
//
//How many possible unique paths are there?

type Point struct {
	x int
	y int
}

var total int

func walkTo(p Point, endX int, endY int)  {
	fmt.Println("walk:", p)

	if p.x == endX && p.y == endY {
		total++
		return
	}

	// 还能往下走
	if p.x < endX {
		next := p
		next.x++
		walkTo(next, endX, endY)
	}

	// 还能往右走
	if p.y < endY {
		next := p
		next.y++
		walkTo(next, endX, endY)
	}
}

func main()  {
	p := Point{0,0}
	walkTo(p, 2, 6)
	fmt.Println(total)
}