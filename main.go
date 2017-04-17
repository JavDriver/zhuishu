package main

import (
	"github.com/cabuda/zhuishu/action"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	router.POST("/book/auto-complete", action.ActionPostAutoComplete)
	router.POST("/book/fuzzy-search", action.ActionPostFuzzySearch)
	router.GET("/book/:book-id", action.ActionGetBookInfo)
	router.GET("/book-source/:book-id", action.ActionGetBookSource)
	router.GET("/book-toc/:source-id", action.ActionGetBookToc)
	router.GET("/book-mix-toc/:book-id", action.ActionGetBookMixToc)
	router.POST("/book-chapter", action.ActionGetBookChapter)

	router.Run(":8080")
}
