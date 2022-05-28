package tvshows

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var (
	TvShows		[]TvShow
	EpisodesLength int
	TotalNumberOfPages int
	CurrentPageNumber int
	CurrentPageCollectedMovies int
	TotalNumberOfMovies int
	CurrentMovie	int
)


func Main() {
	CollectPages(GetNumberOfPages())
}


func GetNumberOfPages() int {
	var err error
	var numberofPages int
	collector := colly.NewCollector()

	collector.OnHTML(".pagination.pagination-lg.justify-content-center", func(element *colly.HTMLElement) {
		element.ForEach(".page-item", func(_ int, element *colly.HTMLElement) {
			title := element.ChildAttr("a", "title")
			href := element.ChildAttr("a", "href")
			if title == "Last" {
				href = strings.ReplaceAll(href, "/tv-show?page=", "")
				numberofPages, err = strconv.Atoi(href)
				HanleError(err)
			}
		})
	})

	collector.Visit("https://tinyzonetv.to/tv-show")

	return numberofPages
}



func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return buff.Bytes()
}


func HanleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}