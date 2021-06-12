package main

import (
	"fmt"
	"gee"
	"net/http"
)

//Engine 是用来处理所有请求的handler，实现ServerHTTP接口就可以传入ListenAndServer函数作为参数

func main() {
	r := gee.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.POST("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
