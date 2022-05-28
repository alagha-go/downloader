package tvshows


import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CollectPages(PagesLength int) {
	for index:=1; index < PagesLength+1; index++ {
		CollectPageTvShows(index)
	}
	SavePagesData()
	PrintGreen("done collecting all the pages data")
}


func CollectPageTvShows(page int) {
	url := "https://tinyzonetv.to/tv-show?page=" + strconv.Itoa(page)
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", CollectTvShows)

	collector.Visit(url)
}


func CollectTvShows(element *colly.HTMLElement) {
	element.ForEach(".flw-item", func(_ int, element *colly.HTMLElement) {
		var TvShow TvShow
		TvShow.ID = primitive.NewObjectID()
        TvShow.Title = element.ChildAttr("a", "title")
        TvShow.ImageUrl = element.ChildAttr("img", "data-src")
        TvShow.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
		index := strings.Index(TvShow.PageUrl, "free-")
    	TvShow.Code = TvShow.PageUrl[index+5:]
		TvShow.Type = "Tv-Show"
		TvShows = append(TvShows, TvShow)
	})
}