package main

import "fmt"

func main() {
	root := makeList([]int{1, 2, 3, 4, 5})
	fmt.Println(printList(removeNthFromEnd(root, 2)))
}
