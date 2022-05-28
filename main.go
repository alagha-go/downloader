package main

import (
	"downloader/lib/crawler/movies"
	"downloader/lib/crawler/tvshows"
	"encoding/json"
	"net/http"
)

type Statistics struct {
	Type									string						`json:"type,omitempty"`
	PagesLength								int							`json:"pages-length,omitempty"`
	CurrentPageNumber						int							`json:"current-page,omitempty"`
	CurrentPageCollectedMovies				int							`json:"page-collected-movies,omitempty"`
	TotalNumberOfMovies						int							`json:"total-movies,omitempty"`
	CurrentMovie							int							`json:"current-movie,omitempty"`
}


func main() {
	movies.Main()
	tvshows.Main()
}


func GetStatistics(res http.ResponseWriter, req http.Request) {
	Statistics := []Statistics{
		{
			Type: "Movie",
			PagesLength: movies.TotalNumberOfPages,
			CurrentPageNumber: movies.CurrentPageNumber,
			CurrentPageCollectedMovies: movies.CurrentPageCollectedMovies,
			TotalNumberOfMovies: len(movies.Movies),
			CurrentMovie: movies.CurrentMovie,
		},
		{
			Type: "Tv-Show",
			PagesLength: tvshows.TotalNumberOfPages,
			CurrentPageNumber: tvshows.CurrentPageNumber,
			CurrentPageCollectedMovies: tvshows.CurrentPageCollectedMovies,
			TotalNumberOfMovies: len(tvshows.TvShows),
			CurrentMovie: tvshows.CurrentMovie,
		},
	}

	json.NewEncoder(res).Encode(Statistics)
}