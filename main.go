package main

import (
	"github.com/cabuda/zhuishu/action"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/book/auto-complete", action.ActionPostAutoComplete)
	router.POST("/book/fuzzy-search", action.ActionPostFuzzySearch)

	router.Run(":8080")
}
