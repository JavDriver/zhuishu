package zhuishu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// 自动补全
func AutoComplete(query string) []string {
	client := http.DefaultClient

	query = url.QueryEscape(query)

	resp, e := client.Get("http://api.zhuishushenqi.com/book/auto-complete?query=" + query)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var ret struct {
		Ok       bool
		Keywords []string
	}
	json.Unmarshal(body, &ret)

	var keywords []string
	if ret.Ok == true {
		keywords = ret.Keywords
	}

	return keywords
}

type BookSummary struct {
	BookId      string `json:"_id"`
	Title       string
	Cat         string
	Author      string
	Cover       string
	ShortIntro  string
	LastChapter string
}

// 模糊搜索书名－获取简介
func FuzzySearchBook(query string, start, limit int) []BookSummary {
	client := http.DefaultClient

	query = url.QueryEscape(query)
	url := fmt.Sprintf("http://api.zhuishushenqi.com/book/fuzzy-search?query=%s&start=%d&limit=%d", query, start, limit)

	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var ret struct {
		Ok    bool
		Books []BookSummary
	}
	json.Unmarshal(body, &ret)

	var books []BookSummary
	if ret.Ok == true {
		books = ret.Books
		for index := range books {
			books[index].Cover = "http://statics.zhuishushenqi.com" + books[index].Cover
		}
	}

	return books
}

type BookDetail struct {
	BookId        string `json:"_id"`
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

// 获取书籍详情
func GetBookDetail(bookId string) BookDetail {
	client := http.DefaultClient

	bookId = url.QueryEscape(bookId)
	url := fmt.Sprintf("http://api.zhuishushenqi.com/book/%s", bookId)

	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var bookDetail BookDetail
	json.Unmarshal(body, &bookDetail)

	bookDetail.Cover = "http://statics.zhuishushenqi.com" + bookDetail.Cover

	return bookDetail
}

type BookSource struct {
	SourceId      string `json:"_id"`
	Source        string
	Name          string
	Link          string
	LastChapter   string
	ChaptersCount int
	Updated       string
	Host          string
}

// 获取书源信息
func GetBookSource(bookId string) []BookSource {
	client := http.DefaultClient

	bookId = url.QueryEscape(bookId)

	url := fmt.Sprintf("http://api.zhuishushenqi.com/toc?view=summary&book=%s", bookId)

	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res []BookSource
	json.Unmarshal(body, &res)

	return res
}

type TocInfo struct {
	Updated  time.Time
	Chapters []struct {
		Title string
		Link  string
	}
}

// 获取章节列表
func GetBookToc(sourceId string) TocInfo {
	sourceId = url.QueryEscape(sourceId)
	url := fmt.Sprintf("http://api.zhuishushenqi.com/toc/%s?view=chapters", sourceId)

	client := http.DefaultClient
	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res TocInfo
	json.Unmarshal(body, &res)

	return res
}

// 获取混合章节列表
func GetBookMixToc(bookId string) TocInfo {
	bookId = url.QueryEscape(bookId)
	url := fmt.Sprintf("http://api.zhuishushenqi.com/mix-toc/%s", bookId)

	client := http.DefaultClient
	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res struct {
		Ok     bool
		MixToc TocInfo
	}
	json.Unmarshal(body, &res)

	var tocInfo TocInfo
	if res.Ok {
		tocInfo = res.MixToc
	}

	return tocInfo
}

type ChapterInfo struct {
	Title string
	Body  string
}

func GetChapterInfo(chapterLink string) ChapterInfo {
	client := http.DefaultClient
	chapterLink = url.QueryEscape(chapterLink)
	url := fmt.Sprintf("http://chapter2.zhuishushenqi.com/chapter/%s?k=2124b73d7e2e1945&t=1468223717", chapterLink)

	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res struct {
		Ok      bool
		Chapter ChapterInfo
	}
	json.Unmarshal(body, &res)

	var chapterInfo ChapterInfo
	if res.Ok {
		chapterInfo = res.Chapter
	}

	return chapterInfo
}
