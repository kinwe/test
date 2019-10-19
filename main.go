package main

import (
	"fmt"
	"runtime"
	"testmath/utils"
	"time"
)

func main() {

	var list = new(utils.List)

	//cur := time.Now().Unix()
	for i := 0; i < 100000000; i++ {
		list.Append(i)
	}
	//tail := time.Now().Unix()
	//for i := 0; i < 100000000 ; i++ {
	//	fmt.Println(list.Pop())
	//}
	//fmt.Println(tail - cur)
	cur := time.Now().Unix()
	list.ReverseNode()
	tail := time.Now().Unix()
	fmt.Println(tail - cur)
	fmt.Println(runtime.NumCPU())

	//list.ShowList()
	//for i := 0; i < 1000; i++  {
	//	go func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}

	fmt.Print(list.Length())
}
