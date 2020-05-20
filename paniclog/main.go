package main

import (
	"fmt"
	"time"
)

func main() {
	// channel生成
	ispanic := make(chan bool)

	// goroutineその１
	go func() {
		// panic検知
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("run time panic: %v", err)
				ispanic <- true // panic発生をmainに伝える
				for {
				}
			}
		}()

		time.Sleep(10 * time.Second)

		// panicる
		panic("poooooo")
	}()

	// goroutineその２
	go func() {
		// panic検知がなければ延々実行される
		for {
			time.Sleep(3 * time.Second)
			fmt.Println("notpanic")
		}
	}()

	// panic検知するまで処理はここでSTOP
	<-ispanic
	return // 終了
}
