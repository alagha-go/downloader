package main

import (
	"downloader/lib/crawler/movies"
	"downloader/lib/crawler/tvshows"
)


func main() {
	movies.Main()
	tvshows.Main()
}