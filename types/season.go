package types

type SeasonListResponse struct {
	Year   int     `json:"season_year`
	Season string  `json:"season_name"`
	Animes []Anime `json:"anime"`
}
