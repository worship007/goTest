// easyGoroutine project main.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello, %s.\n", who)
		}(name) //赋name的值给who
	}

	runtime.Gosched() //让出CPU时间片
}
