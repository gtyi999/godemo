package main

import (
	"time"
	"fmt"
)

func deal(donetype chan int) {
	fmt.Println(time.Now(), " waitting")
	time.Sleep(time.Second * 13)
	donetype <- 2
	fmt.Println(time.Now(), " done")
}

func main() {
	iFlag := -1
	donetype := make(chan int,1)
	go deal(donetype)
	timer := time.NewTimer(time.Second * 10)
	select {
	case temp := <-timer.C:
		fmt.Println("超时时间:", temp)
		iFlag = 1 //超时结束
	case temp := <-donetype:
		fmt.Println("正常结束", temp)
		iFlag = 2 //正常结束
	}
	if iFlag==1 {
		fmt.Println("超时结束")
	} else if iFlag==2 {
		fmt.Println("正常结束")
	} else {
		fmt.Println("未知情况")
	}
	//time.Sleep(time.Second * 11)
	fmt.Println("结束时间:",time.Now())
}
