package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("polling start.")

	// ゴルーチンしてループ処理を外だし
	go doing()

	// mainが終わらないようにする
	for i := 1; ; i++ {
		// 10秒おきに
		time.Sleep(10 * time.Second)
		doing2()
	}
}

func doing() {
	for j := 1; ; j++ {
		// 15秒おきに
		time.Sleep(15 * time.Second)
		fmt.Print("15sec polling")
		fmt.Println(j)
	}
}

func doing2() {
	fmt.Println("10sec polling")
}
