package types

type SeasonName int

const (
	SUMMER SeasonName = iota
	SPRING
	FALL
	WINTER
)

func (t SeasonName) String() string {
	return [...]string{"summer", "spring", "fall", "winter"}[t]
}

type SeasonListResponse struct {
	Year   int     `json:"season_year`
	Season string  `json:"season_name"`
	Animes []Anime `json:"anime"`
}
