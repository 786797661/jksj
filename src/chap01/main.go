/**
* @Author: Gosin
* @Date: 2022/1/16 21:51
 */

package main

import "fmt"

func main() {
	slice := []string{"你", "好", "Go"}
	//slice=add(slice,2,"Go")
	//slice=add(slice,2,"好")
	//slice=add(slice,3,"！")

	slice = del(slice, 1)
	for i, v := range slice {
		fmt.Println("%s  %d", v, i)
	}
}

/**
构建ArrayList加减
*/
func add(slice []string, index int, value string) []string {
	//["你","好",""]
	slice = append(slice, "")
	//slice[index+1:]= [""]
	//slice[index:]= ["GO",""]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func del(slice []string, index int) []string {
	//["你","好","Go"]
	//slice[index+1:]= ["Go"]
	//slice[index:]= ["好","Go"]
	copy(slice[index:], slice[index+1:])
	slice = slice[0 : len(slice)-1]
	return slice
}
