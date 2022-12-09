/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-09 10:40:09
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-09 10:48:47
 * @FilePath     : /GO/Goroutines/index.go
 * @Description  : 协程
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 同步输出
	f("direct")

	// 异步执行

	// 另起协程执行
	go f("goroutine")

	// 匿名函数启动一个协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
