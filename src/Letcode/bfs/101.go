package main

import "fmt"

func main() {
	n := 3
	for i := 0; i < n*2-1; i++ {
		fmt.Print(i % n)
	}
}
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)

}
