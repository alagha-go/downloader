package movies

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)


func CollectPages(PagesLength int) {
	for index:=1; index<PagesLength+1; index++ {
		CollectPageMovies(index)
		if index % 50 == 0 {
			PrintYellow(fmt.Sprintf("Movies    :%d",index))
		}
	}
	PrintBlue(len(Movies))
	SavePagesData()
	PrintGreen("done collecting all the pages data")
}


func CollectPageMovies(page int) {
	url := "https://tinyzonetv.to/movie?page=" + strconv.Itoa(page)
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", CollectMovies)

	collector.Visit(url)
}


func CollectMovies(element *colly.HTMLElement) {
	element.ForEach(".flw-item", func(_ int, element *colly.HTMLElement) {
		var Movie Movie
		Movie.SetMovieID()
        Movie.Title = element.ChildAttr("a", "title")
        Movie.ImageUrl = element.ChildAttr("img", "data-src")
        Movie.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
		index := strings.Index(Movie.PageUrl, "free-")
    	Movie.Code = Movie.PageUrl[index+5:]
		Movie.CollectMovieContent()
		Movies = append(Movies, Movie)
	})
}


func SavePagesData() {
	data := JsonMarshal(Movies)
	ioutil.WriteFile("./DB/Movies/movies.json", data, 0755)
}