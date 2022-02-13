package main

func main() {
	for i := 0; i < 10; i++ {
		res := fibo(i)
		println(res)
	}

}

func fibo(index int) int {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return fibo(index-1) + fibo(index-2)
	}
}
