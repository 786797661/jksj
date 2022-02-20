/**
* @Author: Gosin
* @Date: 2022/2/20 23:16
 */

package stack

import "container/list"

func isValid(s string) bool {
	stack := list.New()
	bytes := []byte(s)
	for _, v := range bytes {
		if v == '(' || v == '[' || v == '{' {
			stack.PushBack(v)
		} else {
			if stack.Len() > 0 {
				if v == ')' {
					if stack.Back().Value.(byte) == '(' {
						stack.Remove(stack.Back())
						continue
					} else {
						return false
					}
				} else if v == ']' {
					if stack.Back().Value.(byte) == '[' {
						stack.Remove(stack.Back())
						continue
					} else {
						return false
					}
				} else if v == '}' {
					if stack.Back().Value.(byte) == '{' {
						stack.Remove(stack.Back())
						continue
					} else {
						return false
					}
				} else {
					return false
				}
			} else {
				return false
			}

		}
	}

	if stack.Len() > 0 {
		return false
	} else {
		return true
	}

}
