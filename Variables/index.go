/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 15:06:09
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 15:09:32
 * @FilePath     : /GO/Variables/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	hello := "hello"
	fmt.Println(hello)

	var world string = "world"
	fmt.Println(world)

	var num1, num2 int = 123, 456
	fmt.Println(num1 + num2)

	flag1, flag2 := false, true
	fmt.Println(flag1, flag2)
	if flag1 {
		fmt.Println("yes flag1!")
	}
	if flag2 {
		fmt.Println("yes flag2!")
	}

}
