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

type ResBookChapter struct {
	Ok      bool
	Chapter struct {
		Title string
		Body  string
	}
}

func ActionGetBookChapter(c *gin.Context) {
	chapterLink := c.PostForm("chapter_link")

	ret := getBookChapter(chapterLink)

	c.JSON(http.StatusOK, ret)
}

func getBookChapter(chapterLink string) gin.H {
	client := http.DefaultClient
	chapterLink = url.QueryEscape(chapterLink)
	url := fmt.Sprintf("http://chapter2.zhuishushenqi.com/chapter/%s?k=2124b73d7e2e1945&t=1468223717", chapterLink)

	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res ResBookChapter
	json.Unmarshal(body, &res)

	return gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   res.Chapter,
	}
}
