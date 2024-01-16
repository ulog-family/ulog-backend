// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "ulog-backend/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	articleGroup := r.Group("/article")
	articleGroup.GET(".", handler.GetArticleList)
	articleGroup.POST(".")
	articleGroup.GET("/:id", handler.GetArticleById)
	articleGroup.GET("/:id/info", handler.GetArticleInfoById)
	articleGroup.GET("/category", handler.GetArticleCategoryList)
	articleGroup.POST("/:id/like")
	articleGroup.DELETE("/:id")

	tagGroup := r.Group("/tag")
	tagGroup.GET(".")
	tagGroup.POST(".")

	userGroup := r.Group("/user")
	userGroup.GET("/:name")
	userGroup.POST("/login/:method")
	userGroup.POST("/register/:method")
	userGroup.POST("/forget")
	userGroup.POST("/logout")
	userGroup.DELETE("/:id")
}
