package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	//
	Writer http.ResponseWriter
	Req    *http.Request
	// request 内容
	Path   string
	Method string
	// response 内容
	StatusCode int
}

// context初始化函数
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// FormValue for post
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// query for get
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 通过可变参数接收非固定数量的参数，可以避免使用切片传入参数
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer) // json需要编解码的过程
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500) // 返回error msg,msg should be in text
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

// 定义了构造四种content-type响应的方法：string,data,html,json
// 代码最开头，给map[string]interface{}起了一个别名gee.H，构建JSON数据时，显得更简洁。
// Context目前只包含了http.ResponseWriter和*http.Request，另外提供了对 Method 和 Path 这两个常用属性的直接访问。
// 提供了访问Query和PostForm参数的方法
