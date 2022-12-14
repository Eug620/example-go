/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-09 10:56:35
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-09 11:01:20
 * @FilePath     : /GO/ChannelBuffering/index.go
 * @Description  : 通道缓冲
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	/**
		默认情况下，通道是 无缓冲 的，这意味着只有对应的接收（<- chan）
		通道准备好接收时，才允许进行发送（chan <-）。
		有缓冲通道 允许在没有对应接收者的情况下，缓存一定数量的值。
	**/

	// 这里我们 make 了一个字符串通道，最多允许缓存 2 个值
	messages := make(chan string, 2)

	// 由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
	messages <- "buffered"
	messages <- "channel"

	// 然后我们可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages <- "buffered1"
	messages <- "channel1"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
