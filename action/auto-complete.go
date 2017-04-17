package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type ResAutoComplete struct {
	Ok       bool
	Keywords []string
}

func ActionPostAutoComplete(c *gin.Context) {
	name := c.PostForm("name")
	ret := autoComplete(name)
	c.JSON(http.StatusOK, ret)
}

func autoComplete(query string) gin.H {
	client := http.DefaultClient

	query = url.QueryEscape(query)

	resp, e := client.Get("http://api.zhuishushenqi.com/book/auto-complete?query=" + query)
	if e != nil {
		fmt.Println(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var ret ResAutoComplete
	json.Unmarshal(body, &ret)

	return gin.H{
		"status": 0,
		"msg":    "ok",
		"data": gin.H{
			"book_names": ret.Keywords,
		},
	}
}
