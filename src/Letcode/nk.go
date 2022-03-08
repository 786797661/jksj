/**
* @Author: Gosin
* @Date: 2022/3/4 20:02
 */

package main

import (
	"sort"
)

//100 10
//95 96 97 98 99 101 102 103 104 105
//func main()  {
//	hs:=[]int{95,105,102,98,99,101,97,103,104,96}
//
//	res:=sortlist(hs,100,10)
//	for _, re := range res {
//		fmt.Printf("%v ",re)
//	}
//
//	//h:=0
//	//n:=0
//	//ans:=make([]int,0)
//	//
//	//fmt.Scan(&h,&n)
//	//for i := 0; i < n; i++ {
//	//		x:=0
//	//		fmt.Scan(&x)
//	//		ans = append(ans, x)
//	//}
//	//res:=sortlist(ans,h,n)
//	////res:=sortlist(hs,100,10)
//	//for _, re := range res {
//	//	fmt.Printf("%v ",re)
//	//}
//}
/**
hs：新班级其他小朋友身高
h:小明身高
n:班级个数
99 101 98 102 97 103 96 104 95 105
*/
func sortlist(hs []int, h, n int) []int {
	if hs == nil {
		return nil
	}
	sort.Ints(hs)
	i := 0
	j := n - 1
	temp := make([]int, 0)
	for i <= j {
		left := Abs(h, hs[i])
		right := Abs(h, hs[j])
		if left > right {
			temp = append(temp, hs[i])
			i++
		} else {
			temp = append(temp, hs[j])
			j--
		}
	}
	res := make([]int, 0)
	for i := len(temp) - 1; i >= 0; i-- {
		res = append(res, temp[i])
	}
	return res
}
func Abs(h, hi int) int {
	if h > hi {
		return h - hi
	}
	return hi - h
}
