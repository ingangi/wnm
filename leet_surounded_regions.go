package main

import "fmt"

//Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
//
//A region is captured by flipping all 'O's into 'X's in that surrounded region.
//
//Example:
//
//X X X X
//X O O X
//X X O X
//X O X X
//After running your function, the board should be:
//
//X X X X
//X X X X
//X X X X
//X O X X
//Explanation:
//
//Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'.
//Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'.
//Two cells are connected if they are adjacent cells connected horizontally or vertically.

func dfsMark(r [][]int, row int, col int) {
	rows := len(r)
	cols := len(r[0])

	if row < 0 || row > rows-1 || col < 0 || col > cols-1 {
		return
	}

	if r[row][col] != 'O' {
		return
	}

	r[row][col] = 'Y'
	dfsMark(r, row-1, col)
	dfsMark(r, row+1, col)
	dfsMark(r, row, col-1)
	dfsMark(r, row, col+1)
}

func surroundedRegions(r [][]int) {
	fmt.Println(r)
	rows := len(r)
	cols := len(r[0])
	// 遍历4条边
	for i, j := 0, 0; j < cols; j++ {
		dfsMark(r, i, j)
	}
	for i, j := rows-1, 0; j < cols; j++ {
		dfsMark(r, i, j)
	}
	for i, j := 0, 0; i < rows; i++ {
		dfsMark(r, i, j)
	}
	for i, j := 0, cols-1; i < rows; i++ {
		dfsMark(r, i, j)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if r[i][j] == 'O' {
				r[i][j] = 'X'
			} else if r[i][j] == 'Y' {
				r[i][j] = 'O'
			}
		}
	}

	fmt.Println(r)
}

func main() {
	region := [][]int{[]int{'X', 'X', 'X', 'X'}, []int{'X', 'O', 'O', 'X'}, []int{'X', 'X', 'O', 'X'}, []int{'X', 'O', 'X', 'X'}}
	surroundedRegions(region)
}
