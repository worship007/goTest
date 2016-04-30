// Timer project main.go
package main

import "time"

import (
	"fmt"
)

func main() {
	sign := make(chan byte)
	now := time.Now()
	t := time.NewTimer(2 * time.Second)
	fmt.Printf("Now time: %v.\n", now)
	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)

	t.Reset(41 * time.Second)

	go func() {
		u := time.NewTicker(5 * time.Second)

	loop:
		for {
			select {
			case <-u.C:
				fmt.Printf("runing!Now time: %v.\n", time.Now())
			case expire = <-t.C:
				u.Stop()
				break loop
			}
		}

		fmt.Printf("Expiration time2: %v.\n", expire)

		sign <- 1
	}()

	<-sign
}
