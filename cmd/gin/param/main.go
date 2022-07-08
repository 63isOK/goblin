package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Param struct {
	A  int       `json:"a"`
	B  int       `form:"b"`
	C  int       `uri:"c"`
	D  []int     `form:"d[]"`
	T  time.Time `form:"t" time_format:"2006-01-02 15:04:05" time_utc:"8"`
	TJ time.Time `json:"tj" time_format:"RFC3339"`
}

func main() {
	g := gin.Default()
	g.Use(SplitArrayParams())
	g.POST("/post/:c", post)
	g.GET("/get/:c", get)
	g.Run()
}

func post(c *gin.Context) {
	var p Param
	c.ShouldBind(&p)
	fmt.Printf("%+v\n", p)
	c.ShouldBindQuery(&p)
	fmt.Printf("%+v\n", p)
	c.ShouldBindJSON(&p)
	fmt.Printf("%+v\n", p)
	c.ShouldBindUri(&p)
	fmt.Printf("%+v\n", p)
}

func get(c *gin.Context) {
	fmt.Println("拆分之前的uri", c.Request.URL.RawQuery, "路径", c.Request.URL.Path)
	var p Param
	c.ShouldBind(&p)
	fmt.Printf("%+v\n", p)
	c.ShouldBindQuery(&p)
	fmt.Printf("%+v\n", p)
	c.ShouldBindJSON(&p)
	fmt.Printf("%+v\n", p)
	c.ShouldBindUri(&p)
	fmt.Printf("%+v\n", p)
}

// c.Query 针对get请求,获取?k1=v1&k2=v2
// c.Param 针对url的占位符

// SplitArrayParams 拆分数组参数
//
// 目前http涉及4种参数
// 	1. body里的json格式参数
//  2. url占位符rest参数
//  3. url的query参数
//  4. form-data参数
// 复杂数组存在于query参数和form-data参数,此中间件只针对这两种参数进行处理
func SplitArrayParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := splitQueryParam(c); err != nil {
			c.Abort()
			return
		}
		if err := splitFormDataParma(c); err != nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

// splitQueryParam 拆分数组类型的query参数
func splitQueryParam(c *gin.Context) error {
	// for k,v :=range c.Query(){
	// 	c.Request.URL
	// }
	fmt.Println("拆分之前的uri", c.Request.URL.RawQuery, "路径", c.Request.URL.Path)
	c.Request.URL.RawQuery += "12345"

	return nil
}

// splitQueryParam 拆分数组类型的form-data参数
func splitFormDataParma(c *gin.Context) error {
	return nil
}
