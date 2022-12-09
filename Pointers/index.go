/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 17:13:51
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 18:24:20
 * @FilePath     : /GO/Pointers/index.go
 * @Description  : 指针的使用,可以修改传入数据
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}

func test(str *string) {
	*str = "world"
}
func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	wtf := "hello"
	fmt.Println(wtf)
	test(&wtf)
	fmt.Println(wtf)
}
