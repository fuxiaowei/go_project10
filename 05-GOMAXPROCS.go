package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOROOT())

	//n := runtime.GOMAXPROCS(1) //将cpu设置为 单核
	//fmt.Println("n = ", n)

	n := runtime.GOMAXPROCS(2) //将cpu设置为 双核
	fmt.Println("n = ", n)

	for {
		go fmt.Println(0) // 子go 程

		fmt.Println(1) // 主 go 程
	}
}
