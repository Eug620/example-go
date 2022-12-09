/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 15:57:54
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 16:05:05
 * @FilePath     : /GO/Slices/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import (
	"fmt"
)

func main() {
	testStr := make([]string, 4)
	fmt.Println(testStr)
	testStr[0] = "a"
	testStr[1] = "b"
	testStr[2] = "c"
	testStr[3] = "d"
	fmt.Println(testStr)
	testStr = append(testStr, "e")
	fmt.Println(testStr)
	testStr = append(testStr, "f", "g")
	fmt.Println(testStr)

	copyStr := make([]string, len(testStr))
	fmt.Println(copyStr)
	// copy(to,from)
	copy(copyStr, testStr)
	fmt.Println(copyStr, testStr)

	// 截取
	fmt.Println(copyStr[2:5]) // 截取 start2 end5 不包含5
	fmt.Println(copyStr[2:])  // start 2
	fmt.Println(copyStr[:3])  // end 3

}
