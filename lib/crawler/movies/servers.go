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


func (Movie *Movie)SetID() {
	for index, server := range Movie.Servers {
		url := "https://tinyzonetv.to/ajax/get_link/"+ server.WatchID
		data, _, err := GetRequest(url, false)
		HanleError(err)
		res, err := UnmarshalLinkResponse(data)
		HanleError(err)
        if server.Name == "Streamlare" {
			Movie.Servers[index].Id = strings.ReplaceAll(res.Link, "https://streamlare.com/e/", "")
			Movie.Servers[index].Url = "https://streamlare.com/v/" + Movie.Servers[index].Id
		}else if server.Name == "Vidcloud"{
			Movie.Servers[index].Id = strings.ReplaceAll(res.Link, "https://rabbitstream.net/embed-4/", "")
			Movie.Servers[index].Id = strings.ReplaceAll(Movie.Servers[index].Id, "?z=", "")
			Movie.Servers[index].Url = "https://rabbitstream.net/embed/m-download/" + Movie.Servers[index].Id
		}else if server.Name == "UpCloud" {
			Movie.Servers[index].Id = strings.ReplaceAll(res.Link, "https://mzzcloud.life/embed-4/", "")
			Movie.Servers[index].Id = strings.ReplaceAll(Movie.Servers[index].Id, "?z=", "")
			Movie.Servers[index].Url = "https://mzzcloud.life/embed/m-download/" + Movie.Servers[index].Id
		}else {
			Movie.Servers[index].Url = res.Link
		}
	}
}