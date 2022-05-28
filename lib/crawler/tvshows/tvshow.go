package tvshows

import "go.mongodb.org/mongo-driver/bson/primitive"

func (TvShow *TvShow) SetTvShowID() {
    ID := primitive.NewObjectID()
    for _, movie := range TvShows {
        if ID == movie.ID {
            TvShow.SetTvShowID()
        }
    }
    TvShow.ID = ID
}