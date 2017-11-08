package zhuishu

import (
	"testing"
)

func Test_AutoComplete(t *testing.T) {
	keywords := AutoComplete("领主")

	t.Log(keywords)
}

func Test_FuzzySearchBook(t *testing.T) {
	bookSummarys := FuzzySearchBook("领主", 0, 1)

	t.Log(bookSummarys)
}

func Test_GetBookDetail(t *testing.T) {
	bookDetail := GetBookDetail("557e0b2bdcfc794e1a1cd8b2")

	t.Log(bookDetail)
}

func Test_GetBookSource(t *testing.T) {
	bookSources := GetBookSource("557e0b2bdcfc794e1a1cd8b2")

	t.Log(bookSources)
}

func Test_GetBookToc(t *testing.T) {
	tocInfo := GetBookToc("56ea51cc7393690a39a1ea77")

	t.Log(tocInfo)
}

func Test_GetBookMixToc(t *testing.T) {
	tocInfo := GetBookMixToc("557e0b2bdcfc794e1a1cd8b2")

	t.Log(tocInfo)
}

func Test_GetChapterInfo(t *testing.T) {
	chapterInfo := GetChapterInfo("http://book.my716.com/getBooks.aspx?method=content&bookId=634203&chapterFile=U_634203_201707291114109988_0915_1384.txt")

	t.Log(chapterInfo)
}
