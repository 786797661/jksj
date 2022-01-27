package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func main() {
	slic := []int{1, 0, 3, 10}
	res := sortedSquares(slic)
	fmt.Println(slic[0:len(slic)])
	fmt.Println(slic[1:len(slic)])
	fmt.Println(res)
	_, err := opfile("1")
	fmt.Printf("original error : %T %v\n", errors.Cause(err), errors.Cause(err))
	fmt.Printf("stack error: \n%+v\n", err)
}

type queryError struct {
	query string
	err   error
}

func opfile(path string) ([]byte, error) {
	_, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open filed")
	}
	return nil, err
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	lastNegIndex := -1
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}

	ans := make([]int, 0, n)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}

	return ans
}

//func sortedSquares(nums []int) []int {
//	sort.Ints(nums)
//	fs:=make([]int,0,0)
//	zs:=make([]int,0,0)
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i]<0 {
//			fs=append(fs, nums[i])
//		}else {
//			zs=append(zs,nums[i])
//		}
//	}
//
//	for i := 0; i < len(fs); i++ {
//		fs[i]=fs[i]*fs[i]
//	}
//	for i := 0; i < len(zs); i++ {
//		zs[i]=zs[i]*zs[i]
//	}
//	res:=make([]int,0,0)
//	fsIndex:=0
//	zsIndex:=0
//	for fsIndex<= len(fs)-1 && zsIndex <=len(zs)-1 {
//		if fs[fsIndex]<=zs[zsIndex] {
//			res=append(res, fs[fsIndex])
//			fsIndex++
//		}else {
//			res=append(res, zs[zsIndex])
//			zsIndex++
//		}
//	}
//	if(fsIndex<= len(fs)-1 ){
//		res=append(res, fs[fsIndex:len(fs)]...)
//	}
//	if (zsIndex<= len(zs)-1 ) {
//		res=append(res, zs[zsIndex:len(zs)]...)
//	}
//	return res
//}
