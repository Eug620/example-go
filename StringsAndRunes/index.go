/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-17 18:25:41
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-17 18:30:27
 * @FilePath     : /GO/StringsAndRunes/index.go
 * @Description  : 特殊字符的处理 看不太懂
 *
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved.
 */
package main

import (
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	println("len:", len(s))
	for i := 0; i < len(s); i++ {
		println("%x ", s[i])
	}

	println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		println("%#U starts at %d\n", runeValue, idx)
	}

	println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		println("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		println("found tee")
	} else if r == 'ส' {
		println("found so sua")
	}
}
