package movies

import (
	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func (Movie *Movie)SetElements(element *colly.HTMLElement) {
    functions := []func(element *colly.HTMLElement){}
    functions = append(functions,  Movie.SetReleased)
    functions = append(functions,  Movie.SetGenre)
    functions = append(functions,  Movie.SetCasts)
    functions = append(functions,  Movie.SetDuration)
    functions = append(functions,  Movie.SetCountries)
    functions = append(functions,  Movie.SetProducers)
    element.ForEach(".row-line", func(index int, element *colly.HTMLElement){
        functions[index](element)
    })
}


func (Movie *Movie) SetMovieID() {
    ID := primitive.NewObjectID()
    for _, movie := range Movies {
        if ID == movie.ID {
            Movie.SetMovieID()
        }
    }
    Movie.ID = ID
}