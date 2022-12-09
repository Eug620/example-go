/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-09 16:52:43
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-09 17:23:35
 * @FilePath     : /GO/Sorting/index.go
 * @Description  : 排序
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"sort"
)

// Go 的 sort 包实现了内建及用户自定义数据类型的排序功能。
// 我们先来看看内建数据类型的排序。
func main() {

	// 排序方法是针对内置数据类型的；
	// 这是一个字符串排序的例子。
	// 注意，它是原地排序的，所以他会直接改变给定的切片，而不是返回一个新切片。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 一个 int 排序的例子。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 我们也可以使用 sort 来检查一个切片是否为有序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
	// 运行程序，打印排序好的字符串和整型切片， 以及数组是否有序的检查结果：true。
	test := []int{1, 2, 3}
	t := sort.IntsAreSorted(test)
	fmt.Println(t)
}
