package main

import (
	"fmt"
	"sort-in-go/quick"
)

func main() {
	arr := []int{6, 9, 3, 5, 3, 5}
	//fmt.Println(bubble.Sort(arr))
	//fmt.Println(merge.Sort(arr))
	fmt.Println(quick.Sort(arr))
}
