package action

import (
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"gopkg.in/gin-gonic/gin.v1"
)

type ResMixToxInfo struct {
	Ok     bool
	MixToc struct {
		Id       string `json:"_id"`
		Book     string
		Chapters []struct {
			Title     string
			Link      string
			Unreadble bool
		}
	}
}

func ActionGetBookMixToc(c *gin.Context) {
	bookId := c.Param("book-id")

	ret := getBookMixToc(bookId)

	c.JSON(http.StatusOK, ret)
}

func getBookMixToc(bookId string) gin.H {
	bookId = url.QueryEscape(bookId)
	url := fmt.Sprintf("http://api.zhuishushenqi.com/mix-toc/%s", bookId)

	client := http.DefaultClient
	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res ResMixToxInfo
	json.Unmarshal(body, &res)

	return gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   res.MixToc,
	}
}
