/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 16:24:34
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 16:28:18
 * @FilePath     : /GO/Functions/index.go
 * @Description  : 函数
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
func main() {
	println(add(1, 2))
	println(subtract(10, 4))
	println(subtract(10, 40))
}
