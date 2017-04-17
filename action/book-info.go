package action

import (
	"net/http"

	"fmt"
	"net/url"

	"io/ioutil"
	"log"

	"encoding/json"

	"gopkg.in/gin-gonic/gin.v1"
)

type ResBookInfo struct {
	Id            string `json:"_id"`
	Title         string
	Author        string
	LongIntro     string
	Cover         string
	Cat           string
	MajorCate     string
	MinorCate     string
	WordCount     int
	ChaptersCount int
	LastChapter   string
}

func ActionGetBookInfo(c *gin.Context) {
	bookId := c.Param("book-id")

	ret := getBookInfo(bookId)

	c.JSON(http.StatusOK, ret)
}

func getBookInfo(bookId string) gin.H {
	client := http.DefaultClient

	bookId = url.QueryEscape(bookId)
	url := fmt.Sprintf("http://api.zhuishushenqi.com/book/%s", bookId)

	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res ResBookInfo
	json.Unmarshal(body, &res)

	return gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   res,
	}
}
