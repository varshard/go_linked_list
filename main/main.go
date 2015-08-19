package main

import (
	"fmt"
	"gitlab.com/varshard/linked_list"
)

func main() {
	n := linkedList.Node{Value: "Node"}
	fmt.Println(n.Value)

	list := linkedList.List{}
	list.Append(3)
	list.Append(4)
	list.Append(5)
	fmt.Println(list.Get().Value)
	//list.Prev()
	fmt.Println(list.Get().Value)
	//list.Add(6)
	fmt.Println(list.Get().Value)

	fmt.Println(list.GetFirst().Value)
	fmt.Println(list.GetLast().Value)
	fmt.Println(list.Next().Value)
	PrintList(list)
	//fmt.Println(list.Remove())
}

func PrintList(list linkedList.List) {
	fmt.Print(list.Get().Value)
	for i := 1; i < list.Len(); i++ {
		n := list.Next()
		fmt.Print(n.Value)
	}
}