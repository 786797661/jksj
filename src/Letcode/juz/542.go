package juz

import "container/list"

func updateMatrix(mat [][]int) [][]int {

	queeu := list.New()

	res := make([][]int, 0)
	for level, v := range mat {
		wid := make([]int, 0)
		for i, k := range v {
			if k == 0 {
				point := []int{level, i}
				queeu.PushBack(point)
			}
			wid = append(wid, 0)
		}
		res = append(res, wid)
	}

	for queeu.Len() > 0 {
		point := queeu.Front().Value.([]int)
		queeu.Remove(queeu.Front())
		//for level, v := range mat {
		//	for i, k := range v {
		//		if k==0{
		//			res[level][i]=0
		//			continue
		//		}
		//		height:=0
		//		width:=0
		//		if level>point[0]{
		//			height =level-point[0]
		//		}else {
		//			height =point[0]-level
		//		}
		//		if i>point[1]{
		//			width=i-point[1]
		//		}else {
		//			width=point[1]-i
		//		}
		//
		//		res[level][i]=min(height+width,res[level][i])
		//	}
		//}
	}
	return res
}
func min(x, y int) int {
	if y == 0 {
		return x
	}
	if x < y {
		return x
	}
	return y
}
