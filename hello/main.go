/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-10 14:54:30
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-10 15:43:59
 * @FilePath     : /GO/hello/main.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	fmt.Print("222")
	fmt.Printf("4444")
	fmt.Println("222")

	fmt.Println(2 + 3)
	fmt.Println(false)
	var test float32 = 1.0 / 3
	fmt.Println(1 / 3)
	fmt.Println(test)
	// var testInt int = par test * 3
	var testFloat float32 = test * 3
	fmt.Println(testFloat)
}
