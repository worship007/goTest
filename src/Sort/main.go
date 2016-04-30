package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{120, 33, 44, 1, 23, 90, 87, 13, 57, 43, 42}
	//正序qsort
	sort.Ints(nums)
	fmt.Println(nums)
	//反序qsort
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
}
