package movies

import "go.mongodb.org/mongo-driver/bson/primitive"


func (Movie *Movie) SetMovieID() {
    ID := primitive.NewObjectID()
    for _, movie := range Movies {
        if ID == movie.ID {
            Movie.SetMovieID()
        }
    }
    Movie.ID = ID
}