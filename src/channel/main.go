// main.go
package main

import "fmt"

func Count(ch chan int, arg_int int, arg_ints *int) {
	*arg_ints = arg_int
	ch <- 1
}

func Set(arg_int int, arg_ints *int) {
	*arg_ints = arg_int
}

func main() {
	chs := make([]chan int, 5)
	ints := make([]int, 5)

	/*for k := 0; k < 5; k++ {
		Set(k, &(ints[k]))
		fmt.Println(ints[k])
	}*/

	for i := 0; i < 5; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i, &(ints[i]))
	}

	for num, ch := range chs {
		<-ch
		fmt.Println(ints[num])
	}
}
