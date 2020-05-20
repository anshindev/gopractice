package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Caller(0))
	fmt.Println(runtime.Caller(1))
	fmt.Println(runtime.Caller(2))
	fmt.Println(runtime.Caller(3))
	fmt.Println(runtime.Caller(4))
}
