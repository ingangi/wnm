package main

import "fmt"

//Given a 2D board and a word, find if the word exists in the grid.
//
//The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
//
//Example:
//
//board =
//[
//['A','B','C','E'],
//['S','F','C','S'],
//['A','D','E','E']
//]
//
//Given word = "ABCCED", return true.
//Given word = "SEE", return true.
//Given word = "ABCB", return false.

func search(board [][]rune, rows int, cols int, rowindex int, colindex int, fitlength int, word string) bool {
	if rowindex < 0 || colindex < 0 {
		return false
	}

	if rowindex >= rows || colindex >= cols {
		return false
	}

	if fitlength >= len(word) {
		return true
	}

	if board[rowindex][colindex] == '*' {
		return false
	}

	if board[rowindex][colindex] != rune(word[fitlength]) {
		return false
	}

	// 到这里表示本字符符合要求
	fmt.Println("found: '", string(word[fitlength]), "' at", rowindex, colindex)
	tmp := board[rowindex][colindex]
	board[rowindex][colindex] = '*'		//标记一下表示已经访问过

	// 上
	if search(board, rows, cols, rowindex-1, colindex, fitlength+1, word) {
		return true
	}

	// 下
	if search(board, rows, cols, rowindex+1, colindex, fitlength+1, word) {
		return true
	}

	// 左
	if search(board, rows, cols, rowindex, colindex-1, fitlength+1, word) {
		return true
	}

	// 右
	if search(board, rows, cols, rowindex, colindex+1, fitlength+1, word) {
		return true
	}


	board[rowindex][colindex] = tmp		//还原回去
	return false
}

func searchWord(board [][]rune, word string) bool {
	// 行列数
	rows, cols := 0, 0
	rows = len(board)
	if rows > 0 {
		cols = len(board[0])
	}
	if rows == 0 || cols == 0 {
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if search(board, rows, cols, i, j, 0, word) {
				return true
			}
		}
	}

	return false
}

func main()  {
	board := [][]rune{[]rune{'A','B','C','E'}, []rune{'S','F','C','S'}, []rune{'A','D','E','E'}}
	word := "ASFCCESEEDF"
	fmt.Println(board, "has", word, "?", searchWord(board, word))
}
