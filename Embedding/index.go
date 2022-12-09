/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-12-08 14:35:17
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-08 18:32:43
 * @FilePath     : /GO/Embedding/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func (b base) test() string {
	return fmt.Sprintf("helloworld %v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	c := base{
		num: 2,
	}
	fmt.Println(c)
	fmt.Println(c.describe())
	fmt.Println(c.test())
}
