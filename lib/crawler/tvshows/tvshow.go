package tvshows

import (
	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func (TvShow *TvShow)SetElements(element *colly.HTMLElement) {
    functions := []func(element *colly.HTMLElement){}
    functions = append(functions,  TvShow.SetReleased)
    functions = append(functions,  TvShow.SetGenre)
    functions = append(functions,  TvShow.SetCasts)
    functions = append(functions,  TvShow.SetDuration)
    functions = append(functions,  TvShow.SetCountries)
    functions = append(functions,  TvShow.SetProducers)
    element.ForEach(".row-line", func(index int, element *colly.HTMLElement){
        functions[index](element)
    })
}



func (TvShow *TvShow) SetTvShowID() {
    ID := primitive.NewObjectID()
    for _, movie := range TvShows {
        if ID == movie.ID {
            TvShow.SetTvShowID()
        }
    }
    TvShow.ID = ID
}