package tvshows



func (TvShow *TvShow) FindNewSeasons(Tvshow *TvShow) []Season {
	var Seasons []Season
	for index := range TvShow.Seasons {
		exists := false
		for index1 := range Tvshow.Seasons {
			if TvShow.Seasons[index].Code == Tvshow.Seasons[index1].Code {
				exists = true
				break
			}
		}
		if !exists {
			Seasons = append(Seasons, TvShow.Seasons[index])
		}
	}
	return Seasons
}