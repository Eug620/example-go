/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 15:10:45
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 15:15:09
 * @FilePath     : /GO/Constants/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"math"
)

const str string = "Constants"

func main() {
	fmt.Println(str)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	fmt.Println(math.Pi)
}
