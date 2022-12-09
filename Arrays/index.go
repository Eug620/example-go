/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 15:43:17
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 15:53:14
 * @FilePath     : /GO/Arrays/index.go
 * @Description  : filenames
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	var intAry [5]int
	fmt.Println(intAry)

	intAry[2] = 100
	fmt.Println(intAry)

	fmt.Println(len(intAry))
	fmt.Println(len("123456"))

	var testAry = [5]int{1, 2, 3, 4, 5}
	var testAry1 = [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(testAry, testAry1)

	var towD [3][5]int
	fmt.Println(towD)
	for i := 0; i < len(towD); i++ {
		for j := 0; j < len(towD[i]); j++ {
			towD[i][j] = i + j
		}
	}

	fmt.Println(towD)
}
