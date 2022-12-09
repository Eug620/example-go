/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 15:21:03
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 15:23:20
 * @FilePath     : /GO/If/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	flag := 1

	if flag%3 == 0 {
		fmt.Println("%3")
	} else if flag%2 == 0 {
		fmt.Println("%2")
	} else {
		fmt.Println("default")
	}
}
