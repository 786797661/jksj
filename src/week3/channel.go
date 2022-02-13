package main

//func main() {
//	c := make(chan int,1)
//	go func() {
//		for i := 0; i < 5; i++ {
//			c <- i
//		}
//		close(c)
//	}()
//	for {
//		if data, ok := <-c; ok {
//			time.Sleep(time.Second * 10)
//			fmt.Println(data)
//		} else {
//			break
//		}
//	}
//	fmt.Println("main结束")
//}
