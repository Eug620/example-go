/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 16:38:14
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 16:52:58
 * @FilePath     : /GO/VariadicFunctions/index.go
 * @Description  : 多参数
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

func test(nums ...int) int {
	count := 0
	for _, value := range nums {
		count += value
	}
	println("params:", nums)
	return count
}
func main() {
	println(test(1, 2))
	println(test(1, 2, 3, 4, 5, 6))
	println(test(1, 2, 3, 4, 5, 6, 7, 8, 9))

	nums := []int{22, 33, 44, 55}
	// nums... 类似于js的解构...ary
	println(test(nums...))
}
