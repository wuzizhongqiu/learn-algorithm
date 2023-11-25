package main

import (
	"2023_11/code"
	"fmt"
)

func main() {
	root := code.TreeNode{
		Val:   9,
		Right: nil,
		Left:  nil,
	}
	fmt.Println(code.PseudoPalindromicPaths(&root))
}
