/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-09 11:11:30
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-09 11:15:46
 * @FilePath     : /GO/Select/index.go
 * @Description  : 通道选择器
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
		Go 的 选择器（select） 让你可以同时等待多个通道操作。
		将协程、通道和选择器结合，是 Go 的一个强大特性。
	**/

	// 在这个例子中，我们将从两个通道中选择。
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在一定时间后接收一个值，
	// 通过这种方式来模拟并行的协程执行（例如，RPC 操作）时造成的阻塞（耗时）。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 我们使用 select 关键字来同时等待这两个值， 并打印各自接收到的值。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	// 跟预期的一样，我们首先接收到值 "one"，然后是 "two"。
}
