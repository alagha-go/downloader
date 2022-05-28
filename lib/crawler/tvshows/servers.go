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


func (Episode *Episode)SetID() {
	for index, server := range Episode.Servers {
		url := "https://tinyzonetv.to/ajax/get_link/"+ server.WatchID
		data, _, err := GetRequest(url, false)
		HanleError(err)
		res, err := UnmarshalLinkResponse(data)
		HanleError(err)
        if server.Name == "Streamlare" {
			Episode.Servers[index].Id = strings.ReplaceAll(res.Link, "https://streamlare.com/e/", "")
			Episode.Servers[index].Url = "https://streamlare.com/v/" + Episode.Servers[index].Id
		}else if server.Name == "Vidcloud"{
			Episode.Servers[index].Id = strings.ReplaceAll(res.Link, "https://rabbitstream.net/embed-4/", "")
			Episode.Servers[index].Id = strings.ReplaceAll(Episode.Servers[index].Id, "?z=", "")
			Episode.Servers[index].Url = "https://rabbitstream.net/embed/m-download/" + Episode.Servers[index].Id
		}else if server.Name == "UpCloud" {
			Episode.Servers[index].Id = strings.ReplaceAll(res.Link, "https://mzzcloud.life/embed-4/", "")
			Episode.Servers[index].Id = strings.ReplaceAll(Episode.Servers[index].Id, "?z=", "")
			Episode.Servers[index].Url = "https://mzzcloud.life/embed/m-download/" + Episode.Servers[index].Id
		}else {
			Episode.Servers[index].Url = res.Link
		}
	}
}