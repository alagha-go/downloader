package movies

import (
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (Movie *Movie)SetServers() {
	collector := colly.NewCollector()

    url := "https://tinyzonetv.to/ajax/movie/episodes/"+ Movie.Code

    collector.OnHTML(".nav", func(element *colly.HTMLElement) {
        element.ForEach(".nav-item", func(index int, element *colly.HTMLElement) {
            var server Server
			server.ID = primitive.NewObjectID()
            server.WatchID = element.ChildAttr("a", "data-linkid")
            server.Name = element.ChildAttr("a", "title")
            server.Name = strings.ReplaceAll(server.Name, "Server ", "")
           	Movie.Servers = append(Movie.Servers, server)
        })
    })
    collector.Visit(url)
}