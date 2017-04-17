package action

import (
	"net/http"

	"fmt"
	"net/url"

	"log"

	"io/ioutil"

	"encoding/json"

	"gopkg.in/gin-gonic/gin.v1"
)

type ResTocInfo struct {
	Id       string `json:"_id"`
	Link     string
	Name     string
	Chapters []struct {
		Title     string
		Link      string
		Unreadble bool
	}
}

func ActionGetBookToc(c *gin.Context) {
	sourceId := c.Param("source-id")

	ret := getBookToc(sourceId)

	c.JSON(http.StatusOK, ret)
}

func getBookToc(sourceId string) gin.H {
	sourceId = url.QueryEscape(sourceId)
	url := fmt.Sprintf("http://api.zhuishushenqi.com/toc/%s?view=chapters", sourceId)

	client := http.DefaultClient
	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res ResTocInfo
	json.Unmarshal(body, &res)

	return gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   res,
	}
}
