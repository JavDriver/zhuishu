package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type N struct {
	Name string
}

func main() {
	url := "http://cq01-rdqa-dev083.cq01.baidu.com:8030/mshop/app/getfirstpageuserinfo?ucid=20142525"
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var m []N
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", m)
}
