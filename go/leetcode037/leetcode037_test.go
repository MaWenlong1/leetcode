package leetcode037

import (
	"fmt"
	"testing"
)

func Test036(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	print(board)
}
func print(bo [][]byte) {
	for i := 0; i < len(bo); i++ {
		for j := 0; j < len(bo[i]); j++ {
			fmt.Printf("%c\t", bo[i][j])
		}
		println()
	}
}
func TestDiv(t *testing.T) {
	num := 9
	a := int(num / 3)
	println(a)

	fmt.Println(byteToInt('3'))
	fmt.Printf("%c", intToByte(3))
}
