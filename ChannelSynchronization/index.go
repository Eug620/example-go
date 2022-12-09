/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-09 11:03:47
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-09 11:04:02
 * @FilePath     : /GO/ChannelSynchronization/index.go
 * @Description  : 通道同步
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	/**
		我们可以使用通道来同步协程之间的执行状态。
		这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。
		如果需要等待多个协程，WaitGroup 是一个更好的选择。
	**/
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
