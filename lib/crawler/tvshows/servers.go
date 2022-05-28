package tvshows

import (
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Episode *Episode)SetServers() {
	collector := colly.NewCollector()

    url := "https://tinyzonetv.to/ajax/v2/episode/servers/"+ Episode.Code

    collector.OnHTML(".nav", func(element *colly.HTMLElement) {
        element.ForEach(".nav-item", func(index int, element *colly.HTMLElement) {
            var server Server
			server.ID = primitive.NewObjectID()
            server.WatchID = element.ChildAttr("a", "data-id")
            server.Name = element.ChildAttr("a", "title")
            server.Name = strings.ReplaceAll(server.Name, "Server ", "")
           	Episode.Servers = append(Episode.Servers, server)
        })
    })
    collector.Visit(url)
}