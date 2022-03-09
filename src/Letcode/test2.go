/**
* @Author: Gosin
* @Date: 2022/3/4 20:47
 */

package main

import (
	"container/list"
)

//6
//10 20 30 15 23 12
//3
//func main()  {
//	//hs:=[]int{10,20,30,15,23,12}
//	//fmt.Print(sumAll(hs,6,3))
//
//		n:=0
//		ans:=make([]int,0)
//
//		fmt.Scan(&n)
//		for i := 0; i < n; i++ {
//				x:=0
//				fmt.Scan(&x)
//				ans = append(ans, x)
//		}
//		size:=0
//		fmt.Scan(&size)
//		res:=sumAll(ans,n,size)
//		//res:=sortlist(hs,100,10)
//
//		fmt.Printf("%v",res)
//
//}

func sumAll(nums []int, n, size int) int {
	stack := list.New()
	remove := list.New()
	sum := 0
	for _, num := range nums {
		stack.PushFront(num)
		if stack.Len() >= size {
			targer := 0
			for i := 0; i < size; i++ {
				v := stack.Back().Value.(int)
				targer += v
				stack.Remove(stack.Back())
				stack.PushFront(v)
			}
			sum = max(sum, targer)
			remove.PushBack(stack.Back())
			stack.Remove(stack.Back())
		}
	}
	return sum
}
func max(sum, targer int) int {
	if targer > sum {
		return targer
	}
	return sum
}
