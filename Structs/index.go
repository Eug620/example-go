/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 18:30:41
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-12-08 14:04:35
 * @FilePath     : /GO/Structs/index.go
 * @Description  : 类型配合指针
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{age: 22})

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	options := newPerson("张三")
	fmt.Println(options)
	options.age = 22
	fmt.Println(options)
	options.name = "李四"
	fmt.Println(options)

}
