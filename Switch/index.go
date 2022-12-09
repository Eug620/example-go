/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 15:24:43
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 15:42:18
 * @FilePath     : /GO/Switch/index.go
 * @Description  : filename
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	i := 1
	fmt.Println("i:", i)
	for i < 10 {
		switch i {
		case 1:
			fmt.Println("i is 1")
		case 2:
			fmt.Println("i is 2")
		case 3:
			fmt.Println("i is 3")
		default:
			fmt.Println("i is", i)
		}
		i++
	}

	fmt.Println("end:", i)

	fmt.Println(time.Now())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // 多条件
		fmt.Println("It's the weekend")
	case time.Thursday: // 单条件
		fmt.Println("It's the Thursday")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12: // 条件判断
		fmt.Println("前半天", t.Hour())
	default:
		fmt.Println("后半天", t.Hour())
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("布尔")
		case int:
			fmt.Println("数字")
		default:
			fmt.Printf("未知类型 %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(whatAmI)
	whatAmI(2.3444)
	whatAmI(math.Pi)
	whatAmI(math.Abs)

}
