package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	// 产品型号接口
	productModel := g.Group("model")
	{
		// TODO: 在gin的路由规范中，下面两个路由会产生冲突，需要修改路由的调用方式
		// productModel.GET("/:id/property", productApi.CreateProductModel)
		productModel.POST("/", f1)
		productModel.GET("/:model_id", f2)
		productModel.GET("/:model_id/property", f6)
		productModel.GET("/", f3)
		productModel.DELETE("/:model_id", f4)
		productModel.PUT("/:model_id", f5)
	}

	g.Run()
}

func f1(*gin.Context) {
	println("f1")
}

func f2(*gin.Context) {
	println("f2")
}

func f3(*gin.Context) {
	println("f3")
}

func f4(*gin.Context) {
	println("f4")
}

func f5(*gin.Context) {
	println("f5")
}

func f6(*gin.Context) {
	println("f6")
}
