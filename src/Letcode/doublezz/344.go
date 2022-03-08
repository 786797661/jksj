/**
* @Author: Gosin
* @Date: 2022/3/1 21:48
 */

package doublezz

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
