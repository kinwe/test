package main

import (
	"fmt"
	"testmath/utils"
	"time"
)

func main() {

	var list  = utils.InitStack()

	cur := time.Now().Unix()
	for i := 0; i < 10000000 ; i++ {
		list.Push(i)
	}
	tail := time.Now().Unix()
	for i := 0; i < 10000000 ; i++ {
		fmt.Println(list.Pop())
	}
	fmt.Println(tail - cur)

	//list.ShowList()

	//fmt.Print(list.Length())
}
