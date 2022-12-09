/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 16:53:21
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 16:55:58
 * @FilePath     : /GO/Closures/index.go
 * @Description  : 类似于闭包
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

func intSeq() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
func main() {
	addSeq := intSeq()
	println(addSeq())
	println(addSeq())
	println(addSeq())
	println(addSeq())
	println(addSeq())
}
