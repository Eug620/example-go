/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-09 14:58:26
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-09 15:19:35
 * @FilePath     : /GO/Tickers/index.go
 * @Description  : Tickers
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	/**
		定时器 是当你想要在未来某一刻执行一次时使用的
		打点器 则是为你想要以固定的时间间隔重复执行而准备的。
		这里是一个打点器的例子，它将定时的执行，直到我们将它停止。
	**/

	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	// 打点器和定时器的机制有点相似：
	// 使用一个通道来发送数据。
	// 这里我们使用通道内建的 select，等待每 1s 到达一次的值。
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 打点器可以和定时器一样被停止。
	// 打点器一旦停止，将不能再从它的通道中接收到值。
	// 我们将在运行 4s 后停止这个打点器。
	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
