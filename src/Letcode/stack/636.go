package stack

func exclusiveTime(n int, logs []string) {

}

//func statck636(n int, logs []string) []int {
//	stack := list.New()
//	mymap := map[string]int{}
//	for _, log := range logs {
//		loga := strings.Split(log, ":")
//		if loga[1] == "end" {
//			for stack.Len() > 0 {
//				indx := stack.Back().Value.(int)
//				ind, _ := strconv.Atoi(loga[2])
//				length := ind - indx + 1
//				value, ok := mymap[loga[0]]
//				if !ok {
//					mymap[loga[0]] = length
//				} else {
//					mymap[loga[0]] = length + value
//				}
//			}
//		} else {
//
//			stack.PushBack(loga[2])
//		}
//
//	}
//}
