/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 15:15:40
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 15:20:02
 * @FilePath     : /GO/For/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	j := 11

	//
	for j < 20 {
		println(j)
		j++
	}

	// 跳出
	for {
		println("hello world")
		break
	}

	// 跳出本轮
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		} else {
			println(i)
		}
	}
}
