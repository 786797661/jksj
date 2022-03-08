/**
* @Author: Gosin
* @Date: 2022/3/4 21:15
 */

package main

import (
	"fmt"
	"strings"
)

type TypeNode struct {
	Val    string
	IsSao  int
	childs []*TypeNode
}

func main() {

	goodser := ""
	errser := ""

	fmt.Scan(&goodser)
	fmt.Scan(&errser)
	goodsersplit := strings.Split(goodser, ",")
	errSers := strings.Split(errser, ",")

	sers := make([][]string, 0)

	for _, s := range goodsersplit {
		ss := strings.Split(s, "-")
		ss[0], ss[1] = ss[1], ss[0]
		sers = append(sers, ss)
	}

	res := tree(sers, errSers)
	resS := ""
	for i, re := range res {
		if i == len(res)-1 {
			resS += re
		} else {
			resS = resS + re + ","
		}
	}
	fmt.Print(resS)
}

func tree(sers [][]string, errSers []string) []string {
	nodes := []TypeNode{}
	for _, ser := range sers {
		val := ser[0]
		val2 := ser[1]
		indexN := -1
		indexC := -1
		for i, node := range nodes {
			if node.Val == val {
				indexN = i
			}
		}
		for i, node := range nodes {
			if node.Val == val2 {
				indexC = i
			}
		}
		if indexN > -1 {
			if indexC > -1 {
				nodes[indexN].childs = append(nodes[indexN].childs, &nodes[indexC])
			} else {
				child := TypeNode{
					Val:    val2,
					childs: nil,
				}
				nodes = append(nodes, child)
				nodes[indexN].childs = append(nodes[indexN].childs, &child)
			}
		} else {
			f := TypeNode{
				Val:    val,
				childs: nil,
			}
			if indexC > -1 {
				f.childs = append(f.childs, &nodes[indexC])
			} else {
				child := TypeNode{
					Val:    val2,
					childs: nil,
				}
				nodes = append(nodes, child)
				f.childs = append(f.childs, &child)
			}
			nodes = append(nodes, f)
		}
	}
	find(nodes, errSers)
	res := []string{}
	for _, node := range nodes {
		if node.IsSao != 1 {
			res = append(res, node.Val)
		}
	}
	return res
}

func find(nodes []TypeNode, errSers []string) []TypeNode {
	for _, ser := range errSers {
		for _, node := range nodes {
			if ser == node.Val {
				dfs(&node)
			}
		}
	}
	return nodes
}

func dfs(node *TypeNode) {
	node.IsSao = 1
	if node.childs == nil {
		return
	}
	for _, i2 := range node.childs {
		dfs(i2)
	}
}
