/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-09 10:50:17
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-09 10:55:52
 * @FilePath     : /GO/Channels/index.go
 * @Description  : 通道
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

func main() {

	messages := make(chan string)

	// 协程中给通道传递值
	go func() { messages <- "ping" }()
	// 拿出通道内的值
	msg := <-messages
	fmt.Println(msg)
}
