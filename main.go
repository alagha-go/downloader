package main

import (
	"downloader/lib/crawler/movies"
	"downloader/lib/crawler/tvshows"
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
