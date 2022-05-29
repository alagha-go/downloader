package main

import (
	"downloader/lib/crawler/movies"
	"downloader/lib/crawler/tvshows"
	"encoding/json"
	"net/http"
)

var (
	PORT = ":6000"
)

type Statistics struct {
	Type									string						`json:"type,omitempty"`
	LoopNumber								int							`json:"loop-number,omitempty"`
	PagesLength								int							`json:"pages-length,omitempty"`
	CurrentPageNumber						int							`json:"current-page,omitempty"`
	CurrentPageCollectedMovies				int							`json:"page-collected-movies,omitempty"`
	TotalNumberOfMovies						int							`json:"total-movies,omitempty"`
	CurrentMovie							int							`json:"current-movie,omitempty"`
}


func main() {
	go movies.Main()
	go tvshows.Main()

	http.HandleFunc("/", GetStatistics)
	http.ListenAndServe(PORT, nil)
}


func GetStatistics(res http.ResponseWriter, req *http.Request) {
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
			LoopNumber: tvshows.LoopNumber,
			PagesLength: tvshows.TotalNumberOfPages,
			CurrentPageNumber: tvshows.CurrentPageNumber,
			CurrentPageCollectedMovies: tvshows.CurrentPageCollectedMovies,
			TotalNumberOfMovies: len(tvshows.TvShows),
			CurrentMovie: tvshows.CurrentMovie,
		},
	}

	json.NewEncoder(res).Encode(Statistics)
}