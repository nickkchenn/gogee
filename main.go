package main

import (
	"gee"
	"net/http"
)

//Engine 是用来处理所有请求的handler，实现ServerHTTP接口就可以传入ListenAndServer函数作为参数

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello there</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("paasword"),
		})
	})

	r.Run(":9999")
}
