package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type response struct {
	Ok       bool
	Keywords []string
}

func main() {
	client := http.DefaultClient

	query := "鹰"
	query = url.QueryEscape(query)

	resp, e := client.Get("http://api.zhuishushenqi.com/book/auto-complete?query=" + query)
	if e != nil {
		fmt.Println(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var ret response
	json.Unmarshal(body, &ret)

	if ret.Ok {
		for _, name := range ret.Keywords {
			fmt.Println(name)
		}
	}
}
