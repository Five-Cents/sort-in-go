package main

import (
	"fmt"
	"sort-in-go/bubble"
	"sort-in-go/insertion"
)

func main() {
	arr := []int{3, 2, 6, 7, 1}
	fmt.Println(bubble.Sort(arr))
	fmt.Println(insertion.Sort(arr))
}
