/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 16:09:22
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 16:14:42
 * @FilePath     : /GO/Maps/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

func main() {
	testMap := make(map[string]int)
	testMap["age"] = 22
	println("testMap:", testMap)
	println("testMap[\"age\"]:", testMap["age"])
	println("len:", len(testMap))
	delete(testMap, "age")

	println("delete....")

	println("testMap:", testMap)
	println("testMap[\"age\"]:", testMap["age"])
	println("len:", len(testMap))

	n := map[string]int{"foo": 1, "bar": 2}
	println("map:", n)
}
