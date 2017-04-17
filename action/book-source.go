package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"gopkg.in/gin-gonic/gin.v1"
)

type ResBookSource struct {
	Id            string `json:"_id"`
	Source        string
	Name          string
	Link          string
	LastChapter   string
	ChaptersCount int
	Updated       string
	Host          string
}

func ActionGetBookSource(c *gin.Context) {
	bookId := c.Param("book-id")

	ret := getBookSource(bookId)
	c.JSON(http.StatusOK, ret)
}

func getBookSource(bookId string) gin.H {
	client := http.DefaultClient

	bookId = url.QueryEscape(bookId)

	url := fmt.Sprintf("http://api.zhuishushenqi.com/toc?view=summary&book=%s", bookId)

	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res []ResBookSource
	json.Unmarshal(body, &res)

	return gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   res,
	}
}
