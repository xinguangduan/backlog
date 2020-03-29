package main

import (
	"github.com/gin-gonic/gin"
	"microservice/backlog/rest"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1/backlog")
	{
		v1.POST("/", rest.Add)      // 添加新条目
		v1.GET("/", rest.All)       // 查询所有条目
		v1.GET("/:id", rest.Take)   // 获取单个条目
		v1.PUT("/:id", rest.Update) // 更新单个条目
		v1.DELETE("/:id", rest.Del) // 删除单个条目
	}
	r.Run(":10000")
}
