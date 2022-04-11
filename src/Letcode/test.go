package main

import (
	"container/list"
	"fmt"
)

func main() {
	//[1,1,1,1,0,1,0,0,1,1]]
	mat := [][]int{{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0, 1, 0}, {1, 1, 1, 1, 0, 1, 0, 0, 1, 1}}
	updateMatrix(mat)
}

func updateMatrix(mat [][]int) [][]int {
	if mat == nil {
		return nil
	}
	level := len(mat)
	wid := len(mat[0])
	queu := list.New()
	res := make([][]int, 0)
	//将所有0插入
	for i := 0; i < level; i++ {
		l := make([]int, 0)
		for k := 0; k < wid; k++ {
			if mat[i][k] == 0 {
				point := []int{i, k}
				queu.PushBack(point)
			}
			l = append(l, 0)
		}
		res = append(res, l)
	}
	for queu.Len() > 0 {
		point := queu.Front().Value.([]int)
		queu.Remove(queu.Front())
		levels := point[0]
		k := point[1]
		if levels == 8 && k == 0 {
			fmt.Println("向下")
		}
		value := res[levels][k]
		value++
		//上，判断是否存在 level-1,k
		if levels-1 >= 0 {
			if mat[levels-1][k] == 1 && res[levels-1][k] == 0 {
				res[levels-1][k] = value
				points := []int{levels - 1, k}
				queu.PushBack(points)
			}
		}
		//下
		if levels+1 <= level-1 {
			if mat[levels+1][k] == 1 && res[levels+1][k] == 0 {
				res[levels+1][k] = value
				points := []int{levels + 1, k}
				queu.PushBack(points)
			}
		}
		//左
		if k-1 >= 0 {
			if mat[levels][k-1] == 1 && res[levels][k-1] == 0 {
				res[levels][k-1] = value
				points := []int{levels, k - 1}
				queu.PushBack(points)
			}
		}
		//右
		if k+1 <= wid-1 {
			if mat[levels][k+1] == 1 && res[levels][k+1] == 0 {
				res[levels][k+1] = value
				points := []int{levels, k + 1}
				queu.PushBack(points)
			}
		}
	}
	return res
}
