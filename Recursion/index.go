/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 17:00:31
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 17:03:33
 * @FilePath     : /GO/Recursion/index.go
 * @Description  : 递归?
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	println(fact(4))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	println(fib(7))
}
