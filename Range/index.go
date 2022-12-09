/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 16:15:32
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 16:23:47
 * @FilePath     : /GO/Range/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	println("sum:", sum)

	// idx,value:=range ary,循环数组
	for i, num := range nums {
		if num == 3 {
			println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// key,value:=range map
	for k, v := range kvs {
		println("%s -> %s\t", k, v)
	}

	// 取key
	for k := range kvs {
		println("key:", k)
	}

	// 字符串: idx,keycode
	for i, c := range "go" {
		println(i, c)
	}
}
