/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-08 14:23:15
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-08 14:23:24
 * @FilePath     : /GO/Methods/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
