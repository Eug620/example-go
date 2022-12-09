/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 16:28:35
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 16:36:38
 * @FilePath     : /GO/MultipleReturnValues/index.go
 * @Description  : 多个返回值
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

func val(i int) int {
	return i
}
func vals() (int, string, map[string]int, func(int) int, []int) {
	str := "hello world"
	maps := map[string]int{"age": 22}
	lists := []int{11, 22, 33}
	return 666, str, maps, val, lists
}

func main() {
	var i, s, m, f, l = vals()
	println(i)
	println(s)
	println(m, m["age"])
	println(f(6))
	println(l, len(l))

	// If you only want a subset of the returned values, use the blank identifier _.
	var _, strss, _, _, _ = vals()
	println(strss)
}
