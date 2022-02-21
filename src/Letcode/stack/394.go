package stack

import (
	"container/list"
	"strconv"
)

func decodeString(s string) string {
	statck := list.New()
	statck2 := list.New()
	res := ""
	if len(s) > 0 {
		//for i := len(s) - 1; i >= 0; i-- {
		//	statck.PushBack(string(s[i]))
		//}
		var number = ""
		for _, v := range s {
			_, err := strconv.Atoi(string(v))
			if err == nil {
				//statck2.PushBack(value)
				number = number + string(v)
			} else {
				if number != "" {
					value, _ := strconv.Atoi(number)
					statck2.PushBack(value)
					number = ""
				}
				if v != ']' {
					statck.PushBack(string(v))
				} else {
					temp := ""
					s := ""
					for statck.Len() > 0 {
						top := statck.Back().Value.(string)
						statck.Remove(statck.Back())
						if top != "[" {
							s = top + s
						} else {
							break
						}
					}
					if statck2.Len() > 0 {
						top_num := statck2.Back().Value.(int)
						statck2.Remove(statck2.Back())
						for top_num > 0 {
							temp = temp + s
							top_num--
						}
					}
					statck.PushBack(temp)
				}
			}

		}

	}
	for statck.Len() > 0 {
		top := statck.Front().Value.(string)
		statck.Remove(statck.Front())
		res += top
	}
	return res
}
