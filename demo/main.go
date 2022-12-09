/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2022-11-08 18:56:36
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2022-11-08 19:06:08
 * @FilePath     : /demo/main.go
 * @Description  : filename
 * 
 * Copyright (c) 2022 by eug yyh3531@163.com, All Rights Reserved. 
 */
 package main

 import (
	 "fmt"
	 "net/http"
 )
 
 func hello(w http.ResponseWriter, req *http.Request) {
	
	 fmt.Fprintf(w, "hello\n")
 }
 
 func headers(w http.ResponseWriter, req *http.Request) {
 
	 for name, headers := range req.Header {
		 for _, h := range headers {
			 fmt.Fprintf(w, "%v: %v\n", name, h)
		 }
	 }
 }

 func world(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "world\n")
 }
 
 func main() {
 
	 http.HandleFunc("/hello", hello)
	 http.HandleFunc("/headers", headers)
	 http.HandleFunc("/world", world)
 
	 http.ListenAndServe(":8090", nil)
 }