package types

type MangaType int

const (
	MANGA MangaType = iota
	NOVELS
	ONESHOTS
	DOUJIN
	MANHWA
	MANHUA
)

func (t MangaType) String() string {
	return [...]string{"manga", "novels", "oneshots", "doujin", "manhwa", "manhua"}[t]
}

type Manga struct {
	ID             int32        `json:"mal_id"`
	URL            string       `json:"url"`
	ImageURL       string       `json:"image_url"`
	Title          string       `json:"title"`
	TitleEnglish   string       `json:"title_english"`
	TitleSynonyms  []string     `json:"title_synonyms"`
	TitleJapanese  string       `json:"title_japanese"`
	Type           string       `json:"type"`
	Status         string       `json:"status"`
	Volumes        int16        `json:"volumes"`
	Chapters       int16        `json:"chapters"`
	Publishing     bool         `json:"publishing"`
	Score          float32      `json:"score"`
	ScoredBy       int32        `json:"scored_by"`
	Rank           int32        `json:"rank"`
	Popularity     int32        `json:"popularity"`
	Members        int32        `json:"members"`
	Favorites      int32        `json:"favorites"`
	Synopsis       string       `json:"synopsis"`
	Background     string       `json:"background"`
	Related        MangaRelated `json:"related"`
	Authors        []Reference  `json:"authors"`
	Serializations []Reference  `json:"serializations"`
	Genres         []Reference  `json:"genres"`
}

type MangaRelated struct {
	Adaptation []Reference `json:"Adaptation"`
	Other      []Reference `json:"Other"`
}

type MangaGenreResponse struct {
	MalURL    Reference `json:"mal_url"`
	ItemCount int32     `json:"item_count"`
	Mangas    []Manga   `json:"manga"`
}

type MangaMagazineResponse struct {
	Meta   Reference `json:"meta"`
	Mangas []Manga   `json:"manga"`
}

type TopMangaResponse struct {
	Top []Manga `json:"top"`
}

type MangaSearchResultResponse struct {
	Results  []Manga `json:"results"`
	LastPage int32   `json:"last_page"`
}
